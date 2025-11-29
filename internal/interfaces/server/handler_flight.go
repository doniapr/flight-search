package server

import (
	"encoding/json"
	"fmt"

	"doniapr.github.io/flight-search/internal/dto"
	"doniapr.github.io/flight-search/internal/interfaces"
	Flight "doniapr.github.io/flight-search/internal/usecase/flight"
	"github.com/labstack/echo/v4"
)

type flightHandler struct {
	flight Flight.Service
}

func SetUpFlightHandler(container *interfaces.Container) *flightHandler {
	if container.Flight == nil {
		panic("Flight Service is not initialized")
	}

	return &flightHandler{
		flight: container.Flight,
	}
}

func (h *flightHandler) Search(e echo.Context) error {
	fmt.Println("Incoming Request")
	ctx := e.Request().Context()
	binder := &echo.DefaultBinder{}

	var req dto.FlightRequest
	if err := e.Bind(&req); err != nil {
		return e.JSON(400, map[string]string{"error": "Invalid request"})
	}
	
	if err := binder.BindHeaders(e, &req); err != nil {
		return e.JSON(400, map[string]string{"error": "Invalid request header"})
	}

	res, err := h.flight.Search(&ctx, req)
	if err != nil {
		return e.JSON(500, map[string]string{"error": err.Error()})
	}

	byteResp, _ := json.Marshal(res)
	fmt.Println(string(byteResp))
	return e.JSON(200, res)
}
