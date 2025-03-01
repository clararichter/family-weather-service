package main

import (
	"log"
	"net/http"
	"os"
	"weather-aggregation-service/internal/app"
	httpclients "weather-aggregation-service/internal/http"
	"weather-aggregation-service/internal/services"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	// load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	port := os.Getenv("APP_PORT")
	apiKey := os.Getenv("API_KEY_WEATHERAPI")

	// create a global logger
	logger := logrus.New()

	httpClient := &http.Client{}
	openMeteoCllient := httpclients.NewOpenMeteoClient(logger, httpClient, "https://api.open-meteo.com/v1/forecast")
	weatherApiClient := httpclients.NewWeatherApiClient(logger, httpClient, "https://api.weatherapi.com/v1/forecast.json", apiKey)
	weatherSummaryService := services.NewWeatherSummaryService(logger, openMeteoCllient, weatherApiClient)
	api := app.NewAPIServer(logger, port, weatherSummaryService)

	if err := api.Run(); err != nil {
		logger.Fatal("Error starting the server:", err)
	}
}
