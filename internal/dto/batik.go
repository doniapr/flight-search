package dto

import "time"

type BatikResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Results []BatikData `json:"results"`
}

type BatikData struct {
	FlightNumber      string `json:"flightNumber"`
	AirlineName       string `json:"airlineName"`
	AirlineIATA       string `json:"airlineIATA"`
	Origin            string `json:"origin"`
	Destination       string `json:"destination"`
	DepartureDateTime string `json:"departureDateTime"`
	ArrivalDateTime   string `json:"arrivalDateTime"`
	TravelTime        string `json:"travelTime"`
	NumberOfStops     int    `json:"numberOfStops"`
	Fare              struct {
		BasePrice    int    `json:"basePrice"`
		Taxes        int    `json:"taxes"`
		TotalPrice   int    `json:"totalPrice"`
		CurrencyCode string `json:"currencyCode"`
		Class        string `json:"class"`
	} `json:"fare"`
	SeatsAvailable  int      `json:"seatsAvailable"`
	AircraftModel   string   `json:"aircraftModel"`
	BaggageInfo     string   `json:"baggageInfo"`
	OnboardServices []string `json:"onboardServices"`
	Connections     []struct {
		StopAirport  string `json:"stopAirport"`
		StopDuration string `json:"stopDuration"`
	} `json:"connections,omitempty"`
}

type BatikRequest struct {
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
