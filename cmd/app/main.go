package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/pooulad/gowfa/model"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("gowfa⛈  => error happend: ")

	// run api version
	// run terminal version
	// with flag -api

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

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatal("weather api response error! please try again.")
	}

	fmt.Println(res.Body)
}
