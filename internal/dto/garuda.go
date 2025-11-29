package dto

import "time"

type GarudaResponse struct {
	Status  string       `json:"status"`
	Flights []GarudaData `json:"flights"`
}

type GarudaData struct {
	FlightId    string `json:"flight_id"`
	Airline     string `json:"airline"`
	AirlineCode string `json:"airline_code"`
	Departure   struct {
		Airport  string `json:"airport"`
		City     string `json:"city"`
		Time     string `json:"time"`
		Terminal string `json:"terminal"`
	} `json:"departure"`
	Arrival struct {
		Airport  string `json:"airport"`
		City     string `json:"city"`
		Time     string `json:"time"`
		Terminal string `json:"terminal"`
	} `json:"arrival"`
	DurationMinutes int    `json:"duration_minutes"`
	Stops           int    `json:"stops"`
	Aircraft        string `json:"aircraft"`
	Price           struct {
		Amount   int    `json:"amount"`
		Currency string `json:"currency"`
	} `json:"price"`
	AvailableSeats int    `json:"available_seats"`
	FareClass      string `json:"fare_class"`
	Baggage        struct {
		CarryOn int `json:"carry_on"`
		Checked int `json:"checked"`
	} `json:"baggage"`
	Amenities []string `json:"amenities,omitempty"`
	Segments  []struct {
		FlightNumber string `json:"flight_number"`
		Departure    struct {
			Airport string `json:"airport"`
			Time    string `json:"time"`
		} `json:"departure"`
		Arrival struct {
			Airport string `json:"airport"`
			Time    string `json:"time"`
		} `json:"arrival"`
		DurationMinutes int `json:"duration_minutes"`
		LayoverMinutes  int `json:"layover_minutes,omitempty"`
	} `json:"segments,omitempty"`
}

type GarudaRequest struct {
	Origin        string    `json:"origin"`
	Destination   string    `json:"destination"`
	DepartureDate time.Time `json:"departure_date"`
	Passengers    int       `json:"passengers"`
	CabinClass    string    `json:"cabin_class"`

	Airlines   string `json:"airline,omitempty"`
	PriceStart int    `json:"priceStart,omitempty"`
	PriceEnd   int    `json:"priceEnd,omitempty"`
	Stops      int    `json:"stops"`
}
