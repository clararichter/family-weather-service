package app

import (
	"net/http"
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
		HandleFunc("/weather-summary", a.handlerWeatherSummary).
		Methods("GET").
		Name("GetWeatherSummary")

	a.logger.WithFields(logrus.Fields{
		"port": a.port,
	}).Info("Server starting")

	return http.ListenAndServe(a.port, router)
}

func (a *APIServer) handlerWeatherSummary(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
