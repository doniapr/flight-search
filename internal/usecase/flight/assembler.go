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
