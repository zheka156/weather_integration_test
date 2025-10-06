package weather

import (
	"encoding/json"
	"fmt"
	"integration-tests/pkg/config"
	"integration-tests/pkg/log"
	"net/http"
	"time"

	"github.com/dailymotion/allure-go"
	"github.com/go-resty/resty/v2"
	"github.com/onsi/gomega"
)

type WeatherAPIClient interface {
	GetWeatherDataForLocation(g gomega.Gomega, location string) (response Response)
}

type Client struct {
	*resty.Client
	log *log.TestLogger
	url string
}

func (c *Client) GetWeatherDataForLocation(g gomega.Gomega, city string) (response Response) {
	allure.Step(allure.Name(fmt.Sprintf("Get weather data for location %s", city)), allure.Action(func() {
		url := fmt.Sprintf("%s/%s?format=j1", c.url, city)
		resp := c.senGetRequest(g, url)
		err := json.Unmarshal(resp.Body(), &response)
		g.Expect(err).ShouldNot(gomega.HaveOccurred(), "err unmarshalling JSON", err)
	}))
	return response
}

func NewClient(l *log.TestLogger, conf *config.Config) *Client {
	r := resty.New()
	r.SetCloseConnection(true)
	r.SetTimeout(15 * time.Second)
	r.SetLogger(l.Zap.Sugar())
	return &Client{Client: r, log: l, url: conf.WeatherClient.URL}
}

func (c *Client) senGetRequest(g gomega.Gomega, url string) (resp *resty.Response) {
	c.log.InfofJSON("Request to ", url)
	var err error
	resp, err = c.R().
		SetHeader("Content-type", "application/json").
		Get(url)
	g.Expect(err).ShouldNot(gomega.HaveOccurred(), "request failed with err")
	c.log.InfofJSON("response is ", string(resp.Body()))
	g.Expect(resp.StatusCode()).Should(gomega.Equal(http.StatusOK))
	return resp
}
