package flight

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"doniapr.github.io/flight-search/internal/dto"
	"doniapr.github.io/flight-search/internal/infrastructure/airasia"
	"doniapr.github.io/flight-search/internal/infrastructure/batik"
	"doniapr.github.io/flight-search/internal/infrastructure/garuda"
	"doniapr.github.io/flight-search/internal/infrastructure/lion"
)

type service struct {
	airAsiaWrp airasia.Wrapper
	batikWrp   batik.Wrapper
	garudaWrp  garuda.Wrapper
	lionWrp    lion.Wrapper
}

func NewService(airAsiaWrp airasia.Wrapper, batikWrp batik.Wrapper, garudaWrp garuda.Wrapper, lionWrp lion.Wrapper) Service {
	if airAsiaWrp == nil {
		panic("airAsiaWrp is nil")
	}

	if batikWrp == nil {
		panic("batikWrp is nil")
	}

	if garudaWrp == nil {
		panic("garudaWrp is nil")
	}

	if lionWrp == nil {
		panic("lionWrp is nil")
	}

	return &service{
		airAsiaWrp: airAsiaWrp,
		batikWrp:   batikWrp,
		garudaWrp:  garudaWrp,
		lionWrp:    lionWrp,
	}
}

func (s *service) Search(ctx *context.Context, req dto.FlightRequest) (resp dto.BaseResponse, err error) {
	ts := time.Now()
	c := make(chan []dto.FlightResponse, 4)
	errChan := make(chan error, 4)
	successQuery, failedQuery := 0, 0
	var flights []dto.FlightResponse
	wg := sync.WaitGroup{}

	if req.Airlines != "" {
		airlines := strings.Split(req.Airlines, ",")
		for _, airline := range airlines {
			switch strings.ToLower(strings.TrimSpace(airline)) {
			case "airasia":
				wg.Add(1)
				go s.queryToAirAsia(&wg, ctx, req, c, errChan)
			case "batik":
				wg.Add(1)
				go s.queryToBatik(&wg, ctx, req, c, errChan)
			case "garuda":
				wg.Add(1)
				go s.queryToGaruda(&wg, ctx, req, c, errChan)
			case "lion":
				wg.Add(1)
				go s.queryToLion(&wg, ctx, req, c, errChan)
			}
		}
	} else {
		wg.Add(4)
		go s.queryToAirAsia(&wg, ctx, req, c, errChan)
		go s.queryToBatik(&wg, ctx, req, c, errChan)
		go s.queryToGaruda(&wg, ctx, req, c, errChan)
		go s.queryToLion(&wg, ctx, req, c, errChan)
	}

	wg.Wait()
	close(c)
	close(errChan)

	for flight := range c {
		fmt.Println("receive data from channel")
		fmt.Println(flight)
		successQuery++
		flights = append(flights, flight...)
	}

	for e := range errChan {
		fmt.Println("error received from channel:", e.Error())
		failedQuery++
	}

	te := time.Now()
	resp = dto.BaseResponse{
		SearchCriteria: struct {
			Origin        string `json:"origin"`
			Destination   string `json:"destination"`
			DepartureDate string `json:"departure_date"`
			Passengers    int    `json:"passengers"`
			CabinClass    string `json:"cabin_class"`
		}{
			Origin:        req.Origin,
			Destination:   req.Destination,
			DepartureDate: req.DepartureDate,
			Passengers:    req.Passengers,
			CabinClass:    req.CabinClass,
		},
		Metadata: struct {
			TotalResults       int  `json:"total_results"`
			ProvidersQueried   int  `json:"providers_queried"`
			ProvidersSucceeded int  `json:"providers_succeeded"`
			ProvidersFailed    int  `json:"providers_failed"`
			SearchTimeMs       int  `json:"search_time_ms"`
			CacheHit           bool `json:"cache_hit"`
		}{
			TotalResults:       len(flights),
			ProvidersQueried:   successQuery + failedQuery,
			ProvidersSucceeded: successQuery,
			ProvidersFailed:    failedQuery,
			SearchTimeMs:       int(te.Sub(ts).Milliseconds()),
			CacheHit:           false,
		},
		Flights: flights,
	}

	return

}

func (s *service) queryToAirAsia(wg *sync.WaitGroup, ctx *context.Context, req dto.FlightRequest, c chan<- []dto.FlightResponse, errChan chan<- error) {
	res, err := s.airAsiaWrp.Find(ctx, assembleAirAsiaReq(req))
	if err != nil {
		errChan <- err
		wg.Done()
		return
	}

	fmt.Println("send data to channel")
	c <- res
	wg.Done()
}

func (s *service) queryToBatik(wg *sync.WaitGroup, ctx *context.Context, req dto.FlightRequest, c chan []dto.FlightResponse, errChan chan<- error) {
	res, err := s.batikWrp.Find(ctx, assembleBatikReq(req))
	if err != nil {
		errChan <- err
		wg.Done()
		return
	}
	c <- res
	wg.Done()
}

func (s *service) queryToGaruda(wg *sync.WaitGroup, ctx *context.Context, req dto.FlightRequest, c chan []dto.FlightResponse, errChan chan<- error) {
	res, err := s.garudaWrp.Find(ctx, assembleGarudaReq(req))
	if err != nil {
		errChan <- err
		wg.Done()
		return
	}
	c <- res
	wg.Done()
}

func (s *service) queryToLion(wg *sync.WaitGroup, ctx *context.Context, req dto.FlightRequest, c chan []dto.FlightResponse, errChan chan<- error) {
	res, err := s.lionWrp.Find(ctx, assembleLionReq(req))
	if err != nil {
		errChan <- err
		wg.Done()
		return
	}
	c <- res
	wg.Done()
}
