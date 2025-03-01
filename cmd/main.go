package main

import (
	"log"
	"os"
	"weather-aggregation-service/internal/app"

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
	// apiKeyWeatherApi := os.Getenv("API_KEY_WEATHERAPI")

	// create a global logger
	logger := logrus.New()

	api := app.NewAPIServer(logger, port)

	if err := api.Run(); err != nil {
		logger.Fatal("Error starting the server:", err)
	}
}
