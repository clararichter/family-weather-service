package models

type WeatherSummary struct {
	Latitude  float32      `json:"latitude"`
	Longitude float32      `json:"longitude"`
	Timezone  string       `json:"timezone,omitempty"`
	Units     *Units       `json:"units"`
	Today     *DayForecast `json:"today"`
	Tomorrow  *DayForecast `json:"tomorrow"`
}

type Units struct {
	Temperature   string `json:"temperature"`
	UvIndexMax    string `json:"uv_index_max,omitempty"`
	Precipitation string `json:"precipitation"`
	WindSpeed     string `json:"wind_speed"`
}

type DayForecast struct {
	Date               string  `json:"date"`
	TemperatureLow     float32 `json:"temperature_low"`
	TemperatureHigh    float32 `json:"temperature_high"`
	PrecipitationTotal float32 `json:"precipitation_total"`
	WindSpeedHigh      float32 `json:"wind_speed_high"`
	WindSpeedLow       float32 `json:"wind_speed_low"`
	UvIndexMax         float32 `json:"uv_index_max"`
}
