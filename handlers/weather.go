package handlers

import (
	"cep-weather/services"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

func ValidateCEP(cep string) (string, error) {
	regex := regexp.MustCompile(`[^0-9-]`)
	if regex.MatchString(cep) {
		return "", fmt.Errorf("invalid CEP: contains disallowed characters")
	}

	cep = strings.ReplaceAll(cep, "-", "")

	// Verifica se o CEP tem exatamente 8 dígitos
	if len(cep) != 8 {
		return "", fmt.Errorf("invalid CEP: must contain exactly 8 digits")
	}

	return cep, nil
}

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")

	validCEP, err := ValidateCEP(cep)
	if err != nil {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	city, err := services.GetCityByCEP(validCEP)
	if err != nil {
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}

	tempC, err := services.GetTemperature(city)
	if err != nil {
		http.Error(w, "failed to get temperature", http.StatusInternalServerError)
		return
	}

	tempF := tempC*1.8 + 32
	tempK := tempC + 273

	response := map[string]float64{
		"temp_C": tempC,
		"temp_F": tempF,
		"temp_K": tempK,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
