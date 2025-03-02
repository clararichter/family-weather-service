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

You will also need to create an .env file in the project root, containing environemnt variables for:

`APP_PORT` - specifies which port the service should be accessible at

`API_KEY_WEATHERAPI` - the API key for the third-party service WeatherAPI.com

### Running locally
You need Go version 1.24.0.

At the root of the directory, build the binary by executing the following:
```
$ go mod download
$ go build -o main ./cmd
```

To run the binary:
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
The container reads `APP_PORT` from the .env file copied onto the image during runtime. 
However, in order to have the service be accessible on your local system, you'll need to 
supply a port mapping, where the container port must match the `APP_PORT` environment variable
set in your .env file.
```
$ APP_PORT=8080           <---- must match APP_PORT in .env
$ LOCALHOST_PORT=8080
```
(`APP_PORT` of container maps to `HOST_PORT` on your local system.)

```
$ docker run -p "$LOCALHOST_PORT":"$APP_PORT" family-weather-service
```