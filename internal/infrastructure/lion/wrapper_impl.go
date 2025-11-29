package lion

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"doniapr.github.io/flight-search/internal/dto"
	"doniapr.github.io/flight-search/internal/shared/helper"
)

type wrapper struct{}

func NewWrapper() Wrapper {
	return &wrapper{}
}

func (w wrapper) Find(ctx *context.Context, req dto.LionRequest) (resp []dto.FlightResponse, err error) {
	ts := time.Now()
	var lionResponse dto.LionResponse
	jsonFile, err := os.Open("./resources/mocks/lion_air_search_response.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer jsonFile.Close()

	byteResp, err := io.ReadAll(jsonFile)
	err = json.Unmarshal(byteResp, &lionResponse)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, v := range lionResponse.Data.AvailableFlights {
		if v.SeatsLeft < req.Passengers || req.Origin != v.Route.From.Code || req.Destination != v.Route.To.Code {
			continue
		}

		// filter price range
		if (req.PriceStart > 0 && v.Pricing.Total < req.PriceStart) || (req.PriceEnd > 0 && v.Pricing.Total > req.PriceEnd) {
			continue
		}

		// filter stops
		if req.Stops > 0 && v.StopCount > req.Stops {
			continue
		}

		departTime, _ := helper.ParseTime(v.Schedule.Departure, "2006-01-02T15:04:05", v.Schedule.DepartureTimezone)
		arriveTime, _ := helper.ParseTime(v.Schedule.Arrival, "2006-01-02T15:04:05", v.Schedule.ArrivalTimezone)

		if !(req.DepartureDate.Year() == departTime.Year() && req.DepartureDate.Month() == departTime.Month() && req.DepartureDate.Day() == departTime.Day()) {
			continue
		}

		totalTimeMinutes := v.FlightTime
		if v.StopCount > 0 {
			// add stop wait time to arrive time
			for _, stop := range v.Layovers {
				totalTimeMinutes += stop.DurationMinutes
			}
		}

		response := dto.FlightResponse{
			Id:       fmt.Sprintf("%s_%s", v.Id, v.Carrier.Name),
			Provider: v.Carrier.Name,
			Airline: struct {
				Name string `json:"name"`
				Code string `json:"code"`
			}{
				Name: v.Carrier.Name,
				Code: v.Carrier.Iata,
			},
			FlightNumber: v.Id,
			Departure: struct {
				Airport   string    `json:"airport"`
				City      string    `json:"city"`
				Datetime  time.Time `json:"datetime"`
				Timestamp int       `json:"timestamp"`
			}{
				Airport:   v.Route.From.Code,
				City:      v.Route.From.City,
				Datetime:  departTime,
				Timestamp: int(departTime.Unix()),
			},
			Arrival: struct {
				Airport   string    `json:"airport"`
				City      string    `json:"city"`
				Datetime  time.Time `json:"datetime"`
				Timestamp int       `json:"timestamp"`
			}{
				Airport:   v.Route.To.Code,
				City:      v.Route.To.City,
				Datetime:  arriveTime,
				Timestamp: int(arriveTime.Unix()),
			},
			Duration: struct {
				TotalMinutes int    `json:"total_minutes"`
				Formatted    string `json:"formatted"`
			}{
				TotalMinutes: totalTimeMinutes,
				Formatted:    helper.MinutesToFormatted(totalTimeMinutes),
			},
			Stops: v.StopCount,
			Price: struct {
				Amount   int    `json:"amount"`
				Currency string `json:"currency"`
			}{
				Amount:   v.Pricing.Total,
				Currency: v.Pricing.Currency,
			},
			AvailableSeats: v.SeatsLeft,
			Aircraft:       v.PlaneType,
			Amenities:      nil,
			Baggage: struct {
				CarryOn string `json:"carry_on"`
				Checked string `json:"checked"`
			}{
				CarryOn: v.Services.BaggageAllowance.Cabin,
				Checked: v.Services.BaggageAllowance.Hold,
			},
		}
		resp = append(resp, response)
	}

	te := time.Now()
	time.Sleep(helper.CalculateDelay(100, 200, int(te.Sub(ts).Milliseconds())))
	return

}
