package locWeather

import (
	"PurpleSchool/weather/geo"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const weatherLocationURL = "https://wttr.in/"

var ErrIncorrectFormat = errors.New("INCORRECT_FORMAT_NUMBER_ERROR")

func GetLocationWeather(geo geo.Data, format int) (string, error) {
	baseUrl, err := url.Parse(weatherLocationURL + geo.City)
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("UNABLE_PARSE_URL_ERROR")
	}

	if format > 4 || format < 1 {
		return "", ErrIncorrectFormat
	}

	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()

	response, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Println(response.Status)
		return "", errors.New("RESPONSE_ERROR")
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", errors.New("READ_ALL_ERROR")
	}

	return string(body), nil
}
