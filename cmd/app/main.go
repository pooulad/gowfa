package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("gowfa => error happend: ")

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	apiKey := os.Getenv("WEATHER_API_KEY")

	res, err := http.Get(fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%v&q=Tehran", apiKey))
	if err != nil {
		panic(err)
	}

	fmt.Println(res.Body)
}
