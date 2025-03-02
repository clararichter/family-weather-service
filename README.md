The family-weather-service provides an API for retrieving weather forecast summaries, with weather data tailored to families planning activities. Internally, it uses two third-party APIs, [Open-Meteo](https://open-meteo.com/) and [WeatherAPI](https://www.weatherapi.com/) to retrieve forecast data which it combines to form a unified forecast summary. The service exposes a single endpoint:

```
GET /weather-summary?location={latitude},{longitude}
```
Where `latitude` and `longitude` represent decimal degree coordinates and must be in the ranges `-90<=latitude<=90` and `-180<=latitude<=180`.

The API is documentet in the Open API spec, as defined in `./api/swagger.yaml`. You may import this file in the online [Swagger editor](https://editor.swagger.io) to check out the API definition.


### Example request:

`GET /weather-summary?location=52.5243,13.4105`

Responds with:

```
{
  "latitude": 52.52,
  "longitude": 13.419998,
  "timezone": "Europe/Berlin",
  "units": {
    "temperature": "C",
    "uv_index_max": "N/A",
    "precipitation": "mm",
    "wind_speed": "km/h"
  },
  "today": {
    "date": "2025-03-02",
    "temperature_low": 0,
    "temperature_high": 8.5,
    "precipitation_total": 0,
    "wind_speed_high": 18,
    "wind_speed_low": 0,
    "uv_index_max": 3.15
  },
  "tomorrow": {
    "date": "2025-03-03",
    "temperature_low": 1.5,
    "temperature_high": 9.7,
    "precipitation_total": 0,
    "wind_speed_high": 16.9,
    "wind_speed_low": 0,
    "uv_index_max": 3.3
  }
}
```

## Run and test the service
Note that, in order to run and test this service, you will need an API key for the third-part service WeatherAPI.com (https://www.weatherapi.com/).

### Running locally
You need Go version 1.24.0.

At the root of the directory, build the binary by executing the following:
```
$ go mod download
$ go build cmd/main.go
```

To run the binary:
Define the following two environment variables, for instance in a .env file at the root:

`APP_PORT` - specifies which port the service should be accessible at

`API_KEY_WEATHERAPI` - the API key for the third-party service WeatherAPI.com

Then start the service:

```
$ go ./main
```


### Using Docker

#### Building Docker image
At the root of the directory run,
```
$ docker build . -t family-weather-service
```

#### Starting a container
You need to specify ports and API key for WeatherAPI.com.
```
$ APP_PORT=8080
$ HOST_PORT=8080
$ API_KEY_WEATHERAPI={your key here}
```
(`APP_PORT` of container maps to `HOST_PORT` on your local system.)

```
$ docker run -p "$HOST_PORT":"$APP_PORT" --env APP_PORT=:"$APP_PORT" --env API_KEY_WEATHERAPI family-weather-service
```