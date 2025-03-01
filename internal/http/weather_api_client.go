package http

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type WeatherApiClient struct {
	logger     *logrus.Logger
	httpClient *http.Client
	url        string
	apiKey     string
}

func NewWeatherApiClient(logger *logrus.Logger, httpClient *http.Client, url string, apiKey string) *WeatherApiClient {
	return &WeatherApiClient{
		logger:     logger,
		httpClient: httpClient,
		url:        url,
		apiKey:     apiKey,
	}
}
