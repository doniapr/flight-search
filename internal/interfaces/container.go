package interfaces

import (
	"fmt"

	"doniapr.github.io/flight-search/internal/infrastructure/airasia"
	Batik "doniapr.github.io/flight-search/internal/infrastructure/batik"
	Garuda "doniapr.github.io/flight-search/internal/infrastructure/garuda"
	Lion "doniapr.github.io/flight-search/internal/infrastructure/lion"
	cfg "doniapr.github.io/flight-search/internal/shared/config"
	Flight "doniapr.github.io/flight-search/internal/usecase/flight"
)

type Container struct {
	Config *cfg.DefaultConfig
	Flight Flight.Service
}

func New() *Container {
	fmt.Println("Try NewContainer ... ")

	// ========Construct Config
	config := cfg.New("./resources/config.json")

	// ========Construct Infra
	airAsia := airasia.NewWrapper()
	batik := Batik.NewWrapper()
	garuda := Garuda.NewWrapper()
	lion := Lion.NewWrapper()

	// ========Construct Usecase
	flight := Flight.NewService(airAsia, batik, garuda, lion)

	container := &Container{
		Config: config,
		Flight: flight,
	}

	return container
}
