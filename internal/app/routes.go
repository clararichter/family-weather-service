package app

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"strconv"
)

// Example Matching strings
// 23.42,23.235
// 23,55
// TODO: should also match 22.,23.
var latLonRegex = regexp.MustCompile(`^(-?\d+(?:\.\d+)?),(-?\d+(?:\.\d+)?)$`)

// TODO create unit tests for function
func extractLatLon(location string) (float32, float32, error) {
	matches := latLonRegex.FindStringSubmatch(location)

	if len(matches) != 3 {
		return 0.0, 0.0, errors.New("invalid location parameter, expected lat,lon")
	}

	latitude, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return 0.0, 0.0, errors.New("failed to parse latitude")
	}

	longitude, err := strconv.ParseFloat(matches[2], 64)
	if err != nil {
		return 0.0, 0.0, errors.New("failed to parse longitude")
	}

	// Validate range
	if latitude < -90.0 || latitude > 90.0 || longitude < -180.0 || longitude > 180.0 {
		return 0.0, 0.0, errors.New("latitude or longitude out of range")
	}

	return float32(latitude), float32(longitude), nil
}

func (a *APIServer) handlerWeatherSummary(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("location")
	latitude, longitude, err := extractLatLon(location)

	if err != nil {
		// TODO also send a message explaining that locations parameter is invalid
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	summary, err := a.weatherSummaryEngine.GenerateWeatherSummary(latitude, longitude)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(summary)
}
