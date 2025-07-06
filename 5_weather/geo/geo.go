package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Data struct {
	City string `json:"city"`
}

type CityValidateResponse struct {
	Error bool `json:"error"`
}

const cityValidateUrl = "https://countriesnow.space/api/v0.1/countries/population/cities"
const getUserLocationUrl = "https://ipapi.co/json/"

func GetMyLocation(city string) (*Data, error) {
	if city != "" {
		successValidateCity := validateCity(city)
		if !successValidateCity {
			panic("CITY_NOT_FOUND")
		}
		return &Data{
			City: city,
		}, nil
	}

	response, err := http.Get(getUserLocationUrl)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	var geo Data
	err = json.Unmarshal(body, &geo)
	if err != nil {
		return nil, errors.New("ERROR_WHILE_UNPACKING_JSON")
	}

	return &geo, nil
}

func validateCity(city string) bool {
	postBody, err := json.Marshal(map[string]string{
		"city": city,
	})
	if err != nil {
		fmt.Println(err)
		return false
	}

	response, err := http.Post(cityValidateUrl, "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println(err)
		return false
	}

	if response.StatusCode != 200 {
		fmt.Println(response.Status)
		return false
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(response.Status)
		return false
	}

	var cityValidate CityValidateResponse
	err = json.Unmarshal(body, &cityValidate)
	if err != nil {
		fmt.Println(response.Status)
		return false
	}

	return !cityValidate.Error
}
