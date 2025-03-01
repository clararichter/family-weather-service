package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	logger *logrus.Logger
	port   string
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
	router := mux.NewRouter()
	// router.HandleFunc("/weather-summary", a.WeatherSummaryEngine.generateSummary)

	a.logger.WithFields(logrus.Fields{
		"port": a.port,
	}).Info("Server starting")

	return http.ListenAndServe(a.port, router)
}
