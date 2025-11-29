package dto

import "time"

type LionResponse struct {
	Success bool `json:"success"`
	Data    struct {
		AvailableFlights []LionData `json:"available_flights"`
	} `json:"data"`
}

type LionData struct {
	Id      string `json:"id"`
	Carrier struct {
		Name string `json:"name"`
		Iata string `json:"iata"`
	} `json:"carrier"`
	Route struct {
		From struct {
			Code string `json:"code"`
			Name string `json:"name"`
			City string `json:"city"`
		} `json:"from"`
		To struct {
			Code string `json:"code"`
			Name string `json:"name"`
			City string `json:"city"`
		} `json:"to"`
	} `json:"route"`
	Schedule struct {
		Departure         string `json:"departure"`
		DepartureTimezone string `json:"departure_timezone"`
		Arrival           string `json:"arrival"`
		ArrivalTimezone   string `json:"arrival_timezone"`
	} `json:"schedule"`
	FlightTime int  `json:"flight_time"`
	IsDirect   bool `json:"is_direct"`
	Pricing    struct {
		Total    int    `json:"total"`
		Currency string `json:"currency"`
		FareType string `json:"fare_type"`
	} `json:"pricing"`
	SeatsLeft int    `json:"seats_left"`
	PlaneType string `json:"plane_type"`
	Services  struct {
		WifiAvailable    bool `json:"wifi_available"`
		MealsIncluded    bool `json:"meals_included"`
		BaggageAllowance struct {
			Cabin string `json:"cabin"`
			Hold  string `json:"hold"`
		} `json:"baggage_allowance"`
	} `json:"services"`
	StopCount int `json:"stop_count,omitempty"`
	Layovers  []struct {
		Airport         string `json:"airport"`
		DurationMinutes int    `json:"duration_minutes"`
	} `json:"layovers,omitempty"`
}

type LionRequest struct {
	Origin        string    `json:"origin"`
	Destination   string    `json:"destination"`
	DepartureDate time.Time `json:"departure_date"`
	Passengers    int       `json:"passengers"`
	CabinClass    string    `json:"cabin_class,omitempty"`

	Airlines   string `json:"airline,omitempty"`
	PriceStart int    `json:"priceStart,omitempty"`
	PriceEnd   int    `json:"priceEnd,omitempty"`
	Stops      int    `json:"stops"`
}
