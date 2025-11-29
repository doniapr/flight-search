package garuda

import (
	"context"

	"doniapr.github.io/flight-search/internal/dto"
)

type Wrapper interface {
	Find(ctx *context.Context, req dto.GarudaRequest) (resp []dto.FlightResponse, err error)
}
