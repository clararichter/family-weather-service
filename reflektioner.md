### Choice of third-party Weather APIs

I considered the following:
- [Open-Meteo](https://open-Meteo.com) - no API key required for non-commerical use;  comprehensive coverage and API; open-source
- [OpenWeather](https://openweathermap.org) - looked promising, but required credit card details for accessing API, so that was a no-go
- [WeatherAPI](https://www.weatherapi.com/docs/) - requires API key but generous rate-limiting; acceptable but a bit messy documentation; not quite as good coverage as Open-Meteo but gives comprehensive forecast data
- [WeatherStack](https://weatherstack.com/documentation) - very strict rate-limiting for free users, so decided against it

...and ended up using Open-Meteo and WeatherAPI.

### Definition of /weather-summary endpoint
The query parameter `locations` is supplied as a combined `{lat},{lon}` parameter.
Why latitude and longitude, as opposed to passing the name of a location?
Because there are locations with conflicting names. I think the best user interface would be to 
allow search for location based on name, have a separate endpoint that returns the geocoding,
the next step in the flow is to make the call to /weather-summary with latitude and longitude.

The response data for a weather summary is very limited given the problem description.
There are many more data points of interest that could be included, such as air quality,
hourly data, weather alerts, ...

### Technical limitations and improvements
- To test out the service, I relied on manual end-to-end tests, defining a number of test cases
in `rest-client.http` and running them through the VS code integration `REST-client`.

- I tried exercising a number of possible inputs, but I may have missed important cases,
e.g. as related to the number format of {lat,lon}. Example edge case: `12.,123.` (i.e. digit with trailing period). The service recognizes this location input as invalid currently, which it probably shouldn't.

- It would be useful to have an automated suite of end-to-end tests. As far as individual data points go, it's difficult to asses for correctness given we so heavily rely on third-parties for the determination of their values.

- I didn't implement any unit tests for the lack of time. I usually like using more of a TDD approach but what
happened, like it often does when there's a bit of time-crunch involved to finish functionality, I used a code-first approach, and didn't get to the point of writing tests.

- Aggregation of weather data from multiple sources is very rudimentary. The approach taken is to use single
data points from two sources and let these complement each other. Where there are conflicts, we let the "most
extreme" value take precedence, e.g., if there are two Daily max temperatures, we use the largest value.
With more time I'd look into ways of incorpating hourly forecast data. There are cases where I'd like more rigourous handling of unexpected response data from the external weather APIs, such as when latitude/longitude is off or other reasons for questioning accuracy of data points.

Some very basic security considerations...
- The service is, on its own, obviously vulnerable to important attack vectors and faults since there is no rate-limiting. If this were to be spun up in production, we'd a layer in front of the service, such as reverse proxy, to guard against DOS attacks, but that's clearly out of scope for this project.

- The services communicates over http and not TLS/https.



### Dependencies used
- gorilla/mux - web framework (std lib net/http may have been sufficient given limited API, but feeling
was gorilla/mux made it slightly simpler)
- sirpusen/logging - logging library that I'm accustomed to and like from before
- godotenv - a convenience library for loading environment variables
- resty - an http convenience library that makes sending http requests easier