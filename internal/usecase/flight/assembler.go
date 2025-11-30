package flight

import (
	"time"

	"doniapr.github.io/flight-search/internal/dto"
)

func assembleAirAsiaReq(in dto.FlightRequest) dto.AirAsiaRequest {
	d, _ := time.Parse("2006-01-02", in.DepartureDate)
	return dto.AirAsiaRequest{
		Origin:        in.Origin,
		Destination:   in.Destination,
		DepartureDate: d,
		Passengers:    in.Passengers,
		CabinClass:    in.CabinClass,

		Airlines:   in.Airlines,
		PriceStart: in.PriceStart,
		PriceEnd:   in.PriceEnd,
		Stops:      in.Stops,
	}
}

func assembleBatikReq(in dto.FlightRequest) dto.BatikRequest {
	d, _ := time.Parse("2006-01-02", in.DepartureDate)
	return dto.BatikRequest{
		Origin:        in.Origin,
		Destination:   in.Destination,
		DepartureDate: d,
		Passengers:    in.Passengers,
		CabinClass:    in.CabinClass,

		Airlines:   in.Airlines,
		PriceStart: in.PriceStart,
		PriceEnd:   in.PriceEnd,
		Stops:      in.Stops,
	}
}

func assembleGarudaReq(in dto.FlightRequest) dto.GarudaRequest {
	d, _ := time.Parse("2006-01-02", in.DepartureDate)
	return dto.GarudaRequest{
		Origin:        in.Origin,
		Destination:   in.Destination,
		DepartureDate: d,
		Passengers:    in.Passengers,
		CabinClass:    in.CabinClass,

		Airlines:   in.Airlines,
		PriceStart: in.PriceStart,
		PriceEnd:   in.PriceEnd,
		Stops:      in.Stops,
	}
}

func assembleLionReq(in dto.FlightRequest) dto.LionRequest {
	d, _ := time.Parse("2006-01-02", in.DepartureDate)
	return dto.LionRequest{
		Origin:        in.Origin,
		Destination:   in.Destination,
		DepartureDate: d,
		Passengers:    in.Passengers,
		CabinClass:    in.CabinClass,

		Airlines:   in.Airlines,
		PriceStart: in.PriceStart,
		PriceEnd:   in.PriceEnd,
		Stops:      in.Stops,
	}
}

func assembleFlightResp(in []dto.FlightResponse, req dto.FlightRequest, successQuery, failedQuery, ts int) dto.BaseResponse {
	return dto.BaseResponse{
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
			TotalResults:       len(in),
			ProvidersQueried:   successQuery + failedQuery,
			ProvidersSucceeded: successQuery,
			ProvidersFailed:    failedQuery,
			SearchTimeMs:       int(ts),
			CacheHit:           false,
		},
		Flights: in,
	}
}
