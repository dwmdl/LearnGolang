package geo_test

import (
	"PurpleSchool/weather/geo"
	"errors"
	"testing"
)

func TestGetMyLocation(t *testing.T) {
	//Arrange -> expected result
	city := "Sochi"
	expected := geo.Data{City: "Sochi"}

	//Act -> execute function with mock data
	got, err := geo.GetMyLocation(city)

	//Assert -> check result with expected
	if err != nil {
		t.Error(err)
	}

	if got.City != expected.City {
		t.Errorf("expected %v, got %v", expected.City, got.City)
	}
}

func TestNotCityTest(t *testing.T) {
	//Arrange
	city := "Tomsksksk"

	//Act
	_, err := geo.GetMyLocation(city)

	//Assert
	if !errors.Is(err, geo.ErrNoCity) {
		t.Errorf("expected %v, got %v", geo.ErrNoCity, err)
	}
}
