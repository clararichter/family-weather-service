package services

import (
	"weather-aggregation-service/internal/models"
)

type WeatherSummaryService struct {
}

func (engine *WeatherSummaryService) GenerateWeatherSummary(latitude float32, longitude float32) (*models.WeatherSummary, error) {
	return &models.WeatherSummary{}, nil
}
