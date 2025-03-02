package services

import (
	"fmt"
	"math"
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
		logger:           logger,
		openMeteoClient:  openMeteoClient,
		weatherApiClient: weatherApiClient,
	}
}

func (service *WeatherSummaryService) GenerateWeatherSummary(latitude float32, longitude float32) (*models.WeatherSummary, error) {
	// TODO send requests concurrently

	// TODO Reasonably, if both endpoints give rather comprehensive forecast data,
	// we could still respond with useful information to the client, as long as at least one succeed
	forecastOpenMeteo, err := service.openMeteoClient.RetrieveForecast(latitude, longitude)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve forecast from Open Meteo")
	}
	forecastWeatherApi, err := service.weatherApiClient.RetrieveForecast(latitude, longitude)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve forecast from WeatherAPI")
	}
	return service.reconcileForecasts(forecastOpenMeteo, forecastWeatherApi)
}

func (service *WeatherSummaryService) reconcileForecasts(omForecast *http.OpenMeteoForecast, waForecast *http.WeatherApiForecast) (*models.WeatherSummary, error) {
	if err := validateForecastData_OpenMeteo(omForecast); err != nil {
		return nil, err
	}
	if err := validateForecastData_WeatherApi(waForecast); err != nil {
		return nil, err
	}
	return &models.WeatherSummary{
		Latitude:  float32(omForecast.Latitude),
		Longitude: float32(omForecast.Longitude),
		Timezone:  omForecast.Timezone,
		Units: &models.Units{
			Temperature:   "C",
			UvIndexMax:    "N/A",
			Precipitation: "mm",
			WindSpeed:     "km/h",
		},
		Today: &models.DayForecast{
			Date:               omForecast.Daily.Time[0],
			TemperatureLow:     float32(math.Min(omForecast.Daily.Temperature2MMin[0], waForecast.Forecast.Forecastday[0].Day.MintempC)),
			TemperatureHigh:    float32(math.Max(omForecast.Daily.Temperature2MMax[0], waForecast.Forecast.Forecastday[0].Day.MaxtempC)),
			PrecipitationTotal: float32(math.Max(omForecast.Daily.PrecipitationSum[0], waForecast.Forecast.Forecastday[0].Day.TotalprecipMm)),
			WindSpeedHigh:      float32(math.Max(omForecast.Daily.WindSpeed10MMax[0], waForecast.Forecast.Forecastday[0].Day.MaxwindKph)),
			WindSpeedLow:       0.0, // TODO remove
			UvIndexMax:         float32(math.Max(omForecast.Daily.UvIndexMax[0], waForecast.Forecast.Forecastday[0].Day.Uv)),
		},
		Tomorrow: &models.DayForecast{
			Date:               omForecast.Daily.Time[1],
			TemperatureLow:     float32(math.Min(omForecast.Daily.Temperature2MMin[1], waForecast.Forecast.Forecastday[1].Day.MintempC)),
			TemperatureHigh:    float32(math.Max(omForecast.Daily.Temperature2MMin[1], waForecast.Forecast.Forecastday[1].Day.MaxtempC)),
			PrecipitationTotal: float32(math.Max(omForecast.Daily.PrecipitationSum[1], waForecast.Forecast.Forecastday[1].Day.TotalprecipMm)),
			WindSpeedHigh:      float32(math.Max(omForecast.Daily.WindSpeed10MMax[1], waForecast.Forecast.Forecastday[1].Day.MaxwindKph)),
			WindSpeedLow:       0.0, // TODO remove
			UvIndexMax:         float32(math.Max(omForecast.Daily.UvIndexMax[1], waForecast.Forecast.Forecastday[1].Day.Uv)),
		},
	}, nil
}

func validateForecastData_OpenMeteo(forecast *http.OpenMeteoForecast) error {
	if forecast == nil {
		return fmt.Errorf("validateForecastData_OpenMeteo: unable to exract weather summary from nil object")
	}

	if len(forecast.Daily.Temperature2MMin) != 2 ||
		len(forecast.Daily.Temperature2MMax) != 2 ||
		len(forecast.Daily.PrecipitationSum) != 2 ||
		len(forecast.Daily.WindSpeed10MMax) != 2 ||
		len(forecast.Daily.UvIndexMax) != 2 {
		return fmt.Errorf("validateForecastData_OpenMeteo: unable to exract weather summary from malformed object")
	}
	return nil
}

func validateForecastData_WeatherApi(forecast *http.WeatherApiForecast) error {
	if forecast == nil {
		return fmt.Errorf("validateForecastData_WeatherApi: unable to exract weather summary from nil object")
	}

	if len(forecast.Forecast.Forecastday) != 2 {
		return fmt.Errorf("validateForecastData_WeatherApi: unable to exract weather summary from malformed object")
	}

	return nil
}
