Note that, in order to run and test this service, you will need an API key for the third-part service WeatherAPI.com (https://www.weatherapi.com/).

## Running locally
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


## Using Docker

### Building Docker image
At the root of the directory run,
```
docker build . -t family-weather-service
```

### Starting a container
You need to specify ports and API key for WeatherAPI.com.
```
APP_PORT=8080
HOST_PORT=8080
API_KEY_WEATHERAPI={your key here}
```
(`APP_PORT` of container maps to `HOST_PORT` on your local system.)

```
docker run -p "$HOST_PORT":"$APP_PORT" --env APP_PORT=:"$APP_PORT" --env API_KEY_WEATHERAPI family-weather-service
```