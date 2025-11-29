package flight

import (
	"context"

	"doniapr.github.io/flight-search/internal/dto"
)

type Service interface {
	Search(ctx *context.Context, req dto.FlightRequest) (resp dto.BaseResponse, err error)
}
