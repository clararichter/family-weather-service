package http

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type OpenMeteoClient struct {
	logger     *logrus.Logger
	httpClient *http.Client
	url        string
}

func NewOpenMeteoClient(logger *logrus.Logger, httpClient *http.Client, url string) *OpenMeteoClient {
	return &OpenMeteoClient{
		logger:     logger,
		httpClient: httpClient,
		url:        url,
	}
}
