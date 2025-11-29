package main

import (
	"doniapr.github.io/flight-search/internal/interfaces"
	"doniapr.github.io/flight-search/internal/interfaces/server"
)

func main() {
	server.StartService(interfaces.New())
}
