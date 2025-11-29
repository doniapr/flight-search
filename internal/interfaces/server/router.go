package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetUpRouter(server *echo.Echo, handler *Handler) {
	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Service Is Running")
	})

	server.POST("/flights/search", handler.flightHandler.Search)
}
