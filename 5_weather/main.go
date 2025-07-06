package main

import (
	"PurpleSchool/weather/geo"
	"PurpleSchool/weather/locWeather"
	"flag"
	"fmt"
)

func main() {
	city := flag.String("city", "", "User city")
	format := flag.Int64("format", 1, "Output format")
	flag.Parse()

	geoData, err := geo.GetMyLocation(*city)
	if err != nil && geoData == nil {
		fmt.Println(err)
		panic("ERROR_WHILE_GETTING_LOCATION")
	}

	weatherData := locWeather.GetLocationWeather(*geoData, *format)
	fmt.Printf("Weather in your location: %s", weatherData)
}
