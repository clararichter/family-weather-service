package app

import (
	"net/http"
	"time"
	"weather-aggregation-service/internal/services"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	logger               *logrus.Logger
	port                 string
	weatherSummaryEngine *services.WeatherSummaryService
}

// NewAPI initializes a new APIServer instance.
func NewAPIServer(logger *logrus.Logger, port string) *APIServer {
	return &APIServer{
		logger: logger,
		port:   port,
	}
}

// Run starts the API server.
func (a *APIServer) Run() error {
	router := mux.NewRouter().StrictSlash(true)

	router.
		HandleFunc("/weather-summary", withLogging(a.logger, a.handlerWeatherSummary)).
		Methods("GET").
		Name("GetWeatherSummary")

	a.logger.WithFields(logrus.Fields{
		"port": a.port,
	}).Info("Server starting")

	return http.ListenAndServe(a.port, router)
}

func withLogging(logger *logrus.Logger, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		logger.WithFields(logrus.Fields{
			"method": r.Method,
			"path":   r.URL.Path,
		}).Info("Endpoint hit")

		next(w, r)

		logger.WithFields(logrus.Fields{
			"method":   r.Method,
			"path":     r.URL.Path,
			"duration": time.Since(start),
		}).Info("Completed request")
	}
}
