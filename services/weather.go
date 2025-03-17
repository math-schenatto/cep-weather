package services

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type WeatherAPIResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func GetTemperature(city string) (float64, error) {
	apiKey := "ef75abdde5f840bca86181556251603"
	encodedCity := url.QueryEscape(city)
	url := "http://api.weatherapi.com/v1/current.json?key=" + apiKey + "&q=" + encodedCity

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var data WeatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, err
	}

	return data.Current.TempC, nil
}
