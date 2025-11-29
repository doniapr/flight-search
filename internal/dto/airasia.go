package dto

import "time"

type AirAsiaResponse struct {
	Status  string        `json:"status"`
	Flights []AirAsiaData `json:"flights"`
}

type AirAsiaData struct {
	FlightCode    string  `json:"flight_code"`
	Airline       string  `json:"airline"`
	FromAirport   string  `json:"from_airport"`
	ToAirport     string  `json:"to_airport"`
	DepartTime    string  `json:"depart_time"`
	ArriveTime    string  `json:"arrive_time"`
	DurationHours float64 `json:"duration_hours"`
	DirectFlight  bool    `json:"direct_flight"`
	PriceIdr      int     `json:"price_idr"`
	Seats         int     `json:"seats"`
	CabinClass    string  `json:"cabin_class"`
	BaggageNote   string  `json:"baggage_note"`
	Stops         []struct {
		Airport         string `json:"airport"`
		WaitTimeMinutes int    `json:"wait_time_minutes"`
	} `json:"stops,omitempty"`
}

type AirAsiaRequest struct {
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
