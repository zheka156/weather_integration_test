package weather

import (
	"strconv"
	"time"

	"github.com/dailymotion/allure-go"
	"github.com/onsi/gomega"
)

type Steps struct {
	weatherClient *Client
}

func NewSteps(weatherAPIClient *Client) *Steps {
	return &Steps{weatherClient: weatherAPIClient}
}

func (s *Steps) WeatherInfoIsReturned(g gomega.Gomega, city string, country string) {

	allure.Step(allure.Name("Current Conditions are actual and contain temperature and humidity"), allure.Action(func() {
		response := s.weatherClient.GetWeatherDataForLocation(g, city)
		g.Expect(response.NearestArea[0].Country[0].Value).Should(gomega.Equal(country))

		g.Expect(response.CurrentCondition).ShouldNot(gomega.BeNil())
		currentConditions := response.CurrentCondition[0]
		g.Expect(currentConditions.TempC).ShouldNot(gomega.BeEmpty())
		g.Expect(currentConditions.Humidity).ShouldNot(gomega.BeEmpty())

		g.Expect(response.Weather[0].Date).Should(gomega.Equal(time.Now().Format("2006-01-02")))

		temp, _ := strconv.Atoi(currentConditions.TempC)
		g.Expect(temp).To((gomega.BeNumerically(">", -40)))
		g.Expect(temp).To(gomega.BeNumerically("<", 60))
	}))
}
