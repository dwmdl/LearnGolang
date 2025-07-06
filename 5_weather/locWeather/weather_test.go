package locWeather_test

import (
	"PurpleSchool/weather/geo"
	"PurpleSchool/weather/locWeather"
	"errors"
	"strings"
	"testing"
)

func TestGetLocationWeather(t *testing.T) {
	//Arrange
	expected := "London"
	geoData := geo.Data{City: expected}
	format := 3

	//Act
	result, _ := locWeather.GetLocationWeather(geoData, format)

	//Assert
	if !strings.Contains(result, geoData.City) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

var testFormatCases = []struct {
	name   string
	format int
}{
	{name: "BigFormat", format: 13},
	{name: "ZeroFormat", format: 0},
	{name: "NegativeFormat", format: -1},
}

func TestIncorrectFormatGetLocationWeather(t *testing.T) {
	for _, tfc := range testFormatCases {
		t.Run(tfc.name, func(t *testing.T) {
			//Arrange
			geoData := geo.Data{City: "London"}

			//Act
			_, err := locWeather.GetLocationWeather(geoData, tfc.format)

			//Assert
			if !errors.Is(err, locWeather.ErrIncorrectFormat) {
				t.Errorf("expected %v, got %v", locWeather.ErrIncorrectFormat, err)
			}
		})
	}
}
