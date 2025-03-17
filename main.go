package main

import (
	"cep-weather/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/weather", handlers.WeatherHandler)
	log.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error runnin server: ", err)
	}
}
