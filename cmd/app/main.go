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

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var weather model.Weather

	err = json.Unmarshal(body, &weather)
	if err != nil {
		log.Fatal(err)
	}

	location, current, hours := weather.Location, weather.Current, weather.Forecast.ForecastDay[0].Hour

	fmt.Printf("%s, %s, %.0fC, %s\n",
		location.Name,
		location.Coutry,
		current.TempC,
		current.Condition.Text,
	)

	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)

		fmt.Printf("%s - %.0fC, %.0f, %s\n",
			date.Format("14:00"),
			hour.TimeC,
			hour.ChainOfRain,
			hour.Condition.Text,
		)
	}
}
