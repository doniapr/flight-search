package airasia

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"

	"doniapr.github.io/flight-search/internal/dto"
	"doniapr.github.io/flight-search/internal/shared/helper"
)

type wrapper struct {
}

func NewWrapper() Wrapper {
	return &wrapper{}
}

func (w *wrapper) Find(ctx *context.Context, req dto.AirAsiaRequest) (resp []dto.FlightResponse, err error) {
	randomNumber := rand.Intn(100)
	// If the number is less than 10 (representing 10%), return an error
	if randomNumber < 10 {
		err = errors.New("error from airasia provider")
		return
	}

	ts := time.Now()
	var airAsiaResp dto.AirAsiaResponse
	jsonFile, err := os.Open("./resources/mocks/airasia_search_response.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer jsonFile.Close()

	byteResp, err := io.ReadAll(jsonFile)
	err = json.Unmarshal(byteResp, &airAsiaResp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, v := range airAsiaResp.Flights {
		if v.Seats < req.Passengers || req.Origin != v.FromAirport || req.Destination != v.ToAirport {
			continue
		}

		// filter price range
		if (req.PriceStart > 0 && v.PriceIdr < req.PriceStart) || (req.PriceEnd > 0 && v.PriceIdr > req.PriceEnd) {
			continue
		}

		// filter stops
		if req.Stops > 0 && len(v.Stops) > req.Stops {
			continue
		}

		departTime, _ := helper.ParseTime(v.DepartTime, "2006-01-02T15:04:05Z07:00", "Asia/Jakarta")
		arriveTime, _ := helper.ParseTime(v.ArriveTime, "2006-01-02T15:04:05Z07:00", "Asia/Jakarta")

		if !(req.DepartureDate.Year() == departTime.Year() && req.DepartureDate.Month() == departTime.Month() && req.DepartureDate.Day() == departTime.Day()) {
			continue
		}

		totalTimeMinutes := v.DurationHours * 60
		if len(v.Stops) > 0 {
			// add stop wait time to arrive time
			for _, stop := range v.Stops {
				totalTimeMinutes += float64(stop.WaitTimeMinutes)
			}
		}

		response := dto.FlightResponse{
			Id:       fmt.Sprintf("%s_%s", v.FlightCode, v.Airline),
			Provider: v.Airline,
			Airline: struct {
				Name string `json:"name"`
				Code string `json:"code"`
			}{
				Name: v.Airline,
				Code: v.FlightCode[0:2],
			},
			FlightNumber: v.FlightCode,
			Departure: struct {
				Airport   string    `json:"airport"`
				City      string    `json:"city"`
				Datetime  time.Time `json:"datetime"`
				Timestamp int       `json:"timestamp"`
			}{
				Airport:   v.FromAirport,
				Datetime:  departTime,
				Timestamp: int(departTime.Unix()),
			},
			Arrival: struct {
				Airport   string    `json:"airport"`
				City      string    `json:"city"`
				Datetime  time.Time `json:"datetime"`
				Timestamp int       `json:"timestamp"`
			}{
				Airport:   v.ToAirport,
				Datetime:  arriveTime,
				Timestamp: int(arriveTime.Unix()),
			},
			Duration: struct {
				TotalMinutes int    `json:"total_minutes"`
				Formatted    string `json:"formatted"`
			}{
				TotalMinutes: int(totalTimeMinutes),
				Formatted:    helper.MinutesToFormatted(int(totalTimeMinutes)),
			},
			Stops: len(v.Stops),
			Price: struct {
				Amount   int    `json:"amount"`
				Currency string `json:"currency"`
			}{
				Amount:   v.PriceIdr,
				Currency: "IDR",
			},
			AvailableSeats: v.Seats,
			CabinClass:     v.CabinClass,
			Aircraft:       nil,
			Amenities:      nil,
			Baggage: struct {
				CarryOn string `json:"carry_on"`
				Checked string `json:"checked"`
			}{
				CarryOn: v.BaggageNote,
			},
		}
		resp = append(resp, response)
	}

	te := time.Now()
	time.Sleep(helper.CalculateDelay(50, 150, int(te.Sub(ts).Milliseconds())))
	return

}
