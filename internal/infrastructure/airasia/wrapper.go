package airasia

import (
	"context"

	"doniapr.github.io/flight-search/internal/dto"
)

type Wrapper interface {
	Find(ctx *context.Context, req dto.AirAsiaRequest) (resp []dto.FlightResponse, err error)
}
