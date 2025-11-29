package server

import (
	"doniapr.github.io/flight-search/internal/interfaces"
	"github.com/labstack/echo/v4"
)

func StartService(container *interfaces.Container) {
	StartHttpService(container)
}

func StartHttpService(container *interfaces.Container) {
	server := echo.New()

	SetUpRouter(server, SetUpHandler(container))

	if err := server.Start(container.Config.AppAddress()); err != nil {
		panic(err)
	}
}
