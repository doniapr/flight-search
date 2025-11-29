package garuda

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"doniapr.github.io/flight-search/internal/dto"
	"doniapr.github.io/flight-search/internal/shared/helper"
)

type wrapper struct{}

func NewWrapper() Wrapper {
	return &wrapper{}
}

func (w wrapper) Find(ctx *context.Context, req dto.GarudaRequest) (resp []dto.FlightResponse, err error) {
	ts := time.Now()
	var garudaResp dto.GarudaResponse
	jsonFile, err := os.Open("./resources/mocks/garuda_indonesia_search_response.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer jsonFile.Close()

	byteResp, err := io.ReadAll(jsonFile)
	err = json.Unmarshal(byteResp, &garudaResp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, v := range garudaResp.Flights {
		if v.AvailableSeats < req.Passengers || req.Origin != v.Departure.Airport || req.Destination != v.Arrival.Airport {
			continue
		}

		// filter price range
		if (req.PriceStart > 0 && v.Price.Amount < req.PriceStart) || (req.PriceEnd > 0 && v.Price.Amount > req.PriceEnd) {
			continue
		}

		// filter stops
		if req.Stops > 0 && v.Stops > req.Stops {
			continue
		}

		departTime, _ := helper.ParseTime(v.Departure.Time, "2006-01-02T15:04:05Z07:00", "Asia/Jakarta")
		arriveTime, _ := helper.ParseTime(v.Arrival.Time, "2006-01-02T15:04:05Z07:00", "Asia/Jakarta")

		if !(req.DepartureDate.Year() == departTime.Year() && req.DepartureDate.Month() == departTime.Month() && req.DepartureDate.Day() == departTime.Day()) {
			continue
		}

		response := dto.FlightResponse{
			Id:       fmt.Sprintf("%s_%s", v.FlightId, v.Airline),
			Provider: v.Airline,
			Airline: struct {
				Name string `json:"name"`
				Code string `json:"code"`
			}{
				Name: v.Airline,
				Code: v.AirlineCode,
			},
			FlightNumber: v.FlightId,
			Departure: struct {
				Airport   string    `json:"airport"`
				City      string    `json:"city"`
				Datetime  time.Time `json:"datetime"`
				Timestamp int       `json:"timestamp"`
			}{
				Airport:   v.Departure.Airport,
				City:      v.Departure.City,
				Datetime:  departTime,
				Timestamp: int(departTime.Unix()),
			},
			Arrival: struct {
				Airport   string    `json:"airport"`
				City      string    `json:"city"`
				Datetime  time.Time `json:"datetime"`
				Timestamp int       `json:"timestamp"`
			}{
				Airport:   v.Arrival.Airport,
				City:      v.Arrival.City,
				Datetime:  arriveTime,
				Timestamp: int(arriveTime.Unix()),
			},
			Duration: struct {
				TotalMinutes int    `json:"total_minutes"`
				Formatted    string `json:"formatted"`
			}{
				TotalMinutes: v.DurationMinutes,
				Formatted:    helper.MinutesToFormatted(v.DurationMinutes),
			},
			Stops: v.Stops,
			Price: struct {
				Amount   int    `json:"amount"`
				Currency string `json:"currency"`
			}{
				Amount:   v.Price.Amount,
				Currency: v.Price.Currency,
			},
			AvailableSeats: v.AvailableSeats,
			Aircraft:       v.Aircraft,
			Amenities:      nil,
			Baggage: struct {
				CarryOn string `json:"carry_on"`
				Checked string `json:"checked"`
			}{
				CarryOn: strconv.Itoa(v.Baggage.CarryOn),
				Checked: strconv.Itoa(v.Baggage.Checked),
			},
		}
		resp = append(resp, response)
	}

	te := time.Now()
	time.Sleep(helper.CalculateDelay(50, 100, int(te.Sub(ts).Milliseconds())))
	return

}
