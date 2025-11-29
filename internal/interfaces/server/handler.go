package server

import "doniapr.github.io/flight-search/internal/interfaces"

type Handler struct {
	flightHandler *flightHandler
}

func SetUpHandler(container *interfaces.Container) *Handler {
	flightHandler := SetUpFlightHandler(container)
	return &Handler{
		flightHandler: flightHandler,
	}
}
