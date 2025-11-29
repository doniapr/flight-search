package batik

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"doniapr.github.io/flight-search/internal/dto"
	"doniapr.github.io/flight-search/internal/shared/helper"
)

type wrapper struct {
}

func NewWrapper() Wrapper {
	return &wrapper{}
}

func (w *wrapper) Find(ctx *context.Context, req dto.BatikRequest) (resp []dto.FlightResponse, err error) {
	ts := time.Now()
	var batikResp dto.BatikResponse
	jsonFile, err := os.Open("./resources/mocks/batik_air_search_response.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer jsonFile.Close()

	byteResp, err := io.ReadAll(jsonFile)
	err = json.Unmarshal(byteResp, &batikResp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, v := range batikResp.Results {
		if v.SeatsAvailable < req.Passengers || req.Origin != v.Origin || req.Destination != v.Destination {
			continue
		}

		// filter price range
		if (req.PriceStart > 0 && v.Fare.TotalPrice < req.PriceStart) || (req.PriceEnd > 0 && v.Fare.TotalPrice > req.PriceEnd) {
			continue
		}

		// filter stops
		if req.Stops > 0 && v.NumberOfStops > req.Stops {
			continue
		}

		departTime, _ := helper.ParseTime(v.DepartureDateTime, "2006-01-02T15:04:05Z0700", "Asia/Jakarta")
		arriveTime, _ := helper.ParseTime(v.ArrivalDateTime, "2006-01-02T15:04:05Z0700", "Asia/Jakarta")
		totalMinute, _ := time.ParseDuration(strings.ReplaceAll(v.TravelTime, " ", ""))

		if !(req.DepartureDate.Year() == departTime.Year() && req.DepartureDate.Month() == departTime.Month() && req.DepartureDate.Day() == departTime.Day()) {
			continue
		}

		if v.NumberOfStops > 0 {
			// add stop wait time to arrive time
			for _, stop := range v.Connections {
				stopDurr, _ := time.ParseDuration(stop.StopDuration)
				totalMinute += stopDurr
			}
		}

		response := dto.FlightResponse{
			Id:       fmt.Sprintf("%s_%s", v.FlightNumber, v.AirlineName),
			Provider: v.AirlineName,
			Airline: struct {
				Name string `json:"name"`
				Code string `json:"code"`
			}{
				Name: v.AirlineName,
				Code: v.AirlineIATA,
			},
			FlightNumber: v.FlightNumber,
			Departure: struct {
				Airport   string    `json:"airport"`
				City      string    `json:"city"`
				Datetime  time.Time `json:"datetime"`
				Timestamp int       `json:"timestamp"`
			}{
				Airport:   v.Origin,
				Datetime:  departTime,
				Timestamp: int(departTime.Unix()),
			},
			Arrival: struct {
				Airport   string    `json:"airport"`
				City      string    `json:"city"`
				Datetime  time.Time `json:"datetime"`
				Timestamp int       `json:"timestamp"`
			}{
				Airport:   v.Destination,
				Datetime:  arriveTime,
				Timestamp: int(arriveTime.Unix()),
			},
			Duration: struct {
				TotalMinutes int    `json:"total_minutes"`
				Formatted    string `json:"formatted"`
			}{
				TotalMinutes: int(totalMinute.Minutes()),
				Formatted:    helper.MinutesToFormatted(int(totalMinute)),
			},
			Stops: v.NumberOfStops,
			Price: struct {
				Amount   int    `json:"amount"`
				Currency string `json:"currency"`
			}{
				Amount:   v.Fare.TotalPrice,
				Currency: v.Fare.CurrencyCode,
			},
			AvailableSeats: v.SeatsAvailable,
			Aircraft:       v.AircraftModel,
			Amenities:      nil,
			Baggage: struct {
				CarryOn string `json:"carry_on"`
				Checked string `json:"checked"`
			}{
				CarryOn: v.BaggageInfo,
			},
		}
		resp = append(resp, response)
	}
	te := time.Now()

	time.Sleep(helper.CalculateDelay(200, 400, int(te.Sub(ts).Milliseconds())))
	return

}
