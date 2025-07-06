package locWeather

import (
	"PurpleSchool/weather/geo"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const weatherLocationURL = "https://wttr.in/"

func GetLocationWeather(geo geo.Data, format int) string {
	baseUrl, err := url.Parse(weatherLocationURL + geo.City)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()

	response, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Println(response.Status)
		return ""
	}

	defer response.Body.Close()

	respBody, err := io.ReadAll(response.Body)

	return string(respBody)
}
