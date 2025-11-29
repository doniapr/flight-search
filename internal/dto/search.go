package dto

import "time"

type BaseResponse struct {
	SearchCriteria struct {
		Origin        string `json:"origin"`
		Destination   string `json:"destination"`
		DepartureDate string `json:"departure_date"`
		Passengers    int    `json:"passengers"`
		CabinClass    string `json:"cabin_class"`
	} `json:"search_criteria"`
	Metadata struct {
		TotalResults       int  `json:"total_results"`
		ProvidersQueried   int  `json:"providers_queried"`
		ProvidersSucceeded int  `json:"providers_succeeded"`
		ProvidersFailed    int  `json:"providers_failed"`
		SearchTimeMs       int  `json:"search_time_ms"`
		CacheHit           bool `json:"cache_hit"`
	} `json:"metadata"`
	Flights []FlightResponse `json:"flights"`
}

type FlightResponse struct {
	Id       string `json:"id"`
	Provider string `json:"provider"`
	Airline  struct {
		Name string `json:"name"`
		Code string `json:"code"`
	} `json:"airline"`
	FlightNumber string `json:"flight_number"`
	Departure    struct {
		Airport   string    `json:"airport"`
		City      string    `json:"city"`
		Datetime  time.Time `json:"datetime"`
		Timestamp int       `json:"timestamp"`
	} `json:"departure"`
	Arrival struct {
		Airport   string    `json:"airport"`
		City      string    `json:"city"`
		Datetime  time.Time `json:"datetime"`
		Timestamp int       `json:"timestamp"`
	} `json:"arrival"`
	Duration struct {
		TotalMinutes int    `json:"total_minutes"`
		Formatted    string `json:"formatted"`
	} `json:"duration"`
	Stops int `json:"stops"`
	Price struct {
		Amount   int    `json:"amount"`
		Currency string `json:"currency"`
	} `json:"price"`
	AvailableSeats int           `json:"available_seats"`
	CabinClass     string        `json:"cabin_class"`
	Aircraft       interface{}   `json:"aircraft"`
	Amenities      []interface{} `json:"amenities"`
	Baggage        struct {
		CarryOn string `json:"carry_on"`
		Checked string `json:"checked"`
	} `json:"baggage"`
}

type FlightRequest struct {
	Origin        string `json:"origin"`
	Destination   string `json:"destination"`
	DepartureDate string `json:"departureDate"`
	ReturnDate    string `json:"returnDate,omitempty"`
	Passengers    int    `json:"passengers"`
	CabinClass    string `json:"cabinClass"`

	Airlines   string `json:"airline,omitempty" header:"airlines"`
	PriceStart int    `json:"priceStart,omitempty" header:"price_start"`
	PriceEnd   int    `json:"priceEnd,omitempty" header:"price_end"`
	Stops      int    `json:"stops" header:"stops"`
	SortBy     string `json:"sortBy,omitempty" header:"sort_by" default:"price asc"`
}
