package main

import (
	"cep-weather/handlers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWeatherHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/weather?cep=13330250", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.WeatherHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Verifique se o corpo da resposta contém as temperaturas
	expectedKeys := []string{"temp_C", "temp_F", "temp_K"}
	for _, key := range expectedKeys {
		if !strings.Contains(rr.Body.String(), key) {
			t.Errorf("handler returned unexpected body: missing key %v", key)
		}
	}
}
