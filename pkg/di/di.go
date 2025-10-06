package di

import (
	"integration-tests/pkg/config"
	"integration-tests/pkg/log"
	"integration-tests/pkg/weather"
)

type Components struct {
	Logger           *log.TestLogger
	WeatherAPIClient *weather.Client
	WeatherSteps     *weather.Steps
}

func InitComponents() *Components {
	conf := config.LoadConfig()
	comps := new(Components)

	comps.Logger = log.NewLogger()
	comps.WeatherAPIClient = weather.NewClient(comps.Logger, conf)
	comps.WeatherSteps = weather.NewSteps(comps.WeatherAPIClient)
	return comps
}
