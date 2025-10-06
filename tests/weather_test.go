package tests

import (
	"fmt"
	"integration-tests/tests/common"
	"testing"

	"github.com/dailymotion/allure-go"
)

const cityName = "London"

func TestGetWeather(t *testing.T) {
	c, g := common.Before(t)

	allure.Test(t, allure.Name(fmt.Sprintf("Get weather info for %s", cityName)),
		allure.Description(fmt.Sprintf("Get weather info for %s and verify that response contains actual info about temperature and humidity", cityName)),
		allure.Action(func() {

			allure.Step(allure.Name(("Check weather info ")), allure.Action(func() {
				c.WeatherSteps.WeatherInfoIsReturned(g, cityName, "United Kingdom")
			}))
		}))

}
