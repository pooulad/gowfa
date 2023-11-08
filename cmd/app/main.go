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
	log.SetPrefix("gowfa⛈  => error happend: ")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	apiKey := os.Getenv("WEATHER_API_KEY")
	q := os.Args[1]

	res, err := http.Get(fmt.Sprintf("https://api.weatherapi.com/v1/forecast.json?key=%v&q=%v", apiKey, q))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.Body)
}
