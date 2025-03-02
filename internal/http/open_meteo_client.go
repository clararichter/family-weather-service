package http

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

type OpenMeteoClient struct {
	logger     *logrus.Logger
	httpClient *resty.Client
	url        string
}

type OpenMeteoForecast struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	GenerationtimeMs     float64 `json:"generationtime_ms"`
	UtcOffsetSeconds     int     `json:"utc_offset_seconds"`
	Timezone             string  `json:"timezone"`
	TimezoneAbbreviation string  `json:"timezone_abbreviation"`
	Elevation            float64 `json:"elevation"`
	HourlyUnits          struct {
		Time                     string `json:"time"`
		Temperature2M            string `json:"temperature_2m"`
		PrecipitationProbability string `json:"precipitation_probability"`
		Precipitation            string `json:"precipitation"`
		WindSpeed10M             string `json:"wind_speed_10m"`
		WindDirection10M         string `json:"wind_direction_10m"`
	} `json:"hourly_units"`
	Hourly struct {
		Time                     []string  `json:"time"`
		Temperature2M            []float64 `json:"temperature_2m"`
		PrecipitationProbability []int     `json:"precipitation_probability"`
		Precipitation            []float64 `json:"precipitation"`
		WindSpeed10M             []float64 `json:"wind_speed_10m"`
		WindDirection10M         []int     `json:"wind_direction_10m"`
	} `json:"hourly"`
	DailyUnits struct {
		Time             string `json:"time"`
		Temperature2MMax string `json:"temperature_2m_max"`
		Temperature2MMin string `json:"temperature_2m_min"`
		Sunrise          string `json:"sunrise"`
		Sunset           string `json:"sunset"`
		UvIndexMax       string `json:"uv_index_max"`
		PrecipitationSum string `json:"precipitation_sum"`
		WindSpeed10MMax  string `json:"wind_speed_10m_max"`
		WindGusts10MMax  string `json:"wind_gusts_10m_max"`
	} `json:"daily_units"`
	Daily struct {
		Time             []string  `json:"time"`
		Temperature2MMax []float64 `json:"temperature_2m_max"`
		Temperature2MMin []float64 `json:"temperature_2m_min"`
		Sunrise          []string  `json:"sunrise"`
		Sunset           []string  `json:"sunset"`
		UvIndexMax       []float64 `json:"uv_index_max"`
		PrecipitationSum []float64 `json:"precipitation_sum"`
		WindSpeed10MMax  []float64 `json:"wind_speed_10m_max"`
		WindGusts10MMax  []float64 `json:"wind_gusts_10m_max"`
	} `json:"daily"`
}

func NewOpenMeteoClient(logger *logrus.Logger, httpClient *resty.Client, url string) *OpenMeteoClient {
	return &OpenMeteoClient{
		logger:     logger,
		httpClient: httpClient,
		url:        url,
	}
}

func (client *OpenMeteoClient) RetrieveForecast(latitude float32, longitude float32) (*OpenMeteoForecast, error) {
	var result OpenMeteoForecast

	resp, err := client.httpClient.R().
		SetQueryParams(map[string]string{
			"latitude":      fmt.Sprint(latitude),
			"longitude":     fmt.Sprint(longitude),
			"hourly":        "temperature_2m,precipitation_probability,precipitation,wind_speed_10m,wind_direction_10m",
			"daily":         "temperature_2m_max,temperature_2m_min,sunrise,sunset,uv_index_max,precipitation_sum,wind_speed_10m_max,wind_gusts_10m_max",
			"timezone":      "Europe/Berlin",
			"forecast_days": "2",
		}).
		SetHeader("Accept", "application/json").
		SetResult(&result).
		Get(client.url)

	if err != nil {
		client.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error sending request in OpenMeteoClient.RetrieveForecast")
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	if resp.IsError() {
		client.logger.WithFields(logrus.Fields{
			"status code": resp.StatusCode(),
		}).Error("Unexpected HTTP response in OpenMeteoClient.RetrieveForecast")

		return nil, fmt.Errorf("unexpected HTTP response: %s", resp.Status())
	}

	return &result, nil
}
