package services

import (
	"weather-aggregation-service/internal/http"
	"weather-aggregation-service/internal/models"

	"github.com/sirupsen/logrus"
)

type WeatherSummaryService struct {
	logger           *logrus.Logger
	openMeteoClient  *http.OpenMeteoClient
	weatherApiClient *http.WeatherApiClient
}

func NewWeatherSummaryService(
	logger *logrus.Logger,
	openMeteoClient *http.OpenMeteoClient,
	weatherApiClient *http.WeatherApiClient) *WeatherSummaryService {
	return &WeatherSummaryService{
		logger: logger,
	}
}

func (engine *WeatherSummaryService) GenerateWeatherSummary(latitude float32, longitude float32) (*models.WeatherSummary, error) {
	return &models.WeatherSummary{}, nil
}
