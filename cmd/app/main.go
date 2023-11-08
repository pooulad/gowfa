package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/pooulad/gowfa/model"
	util "github.com/pooulad/gowfa/util/colorize"
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

	if len(os.Args) < 2 {
		util.Colorize(util.ColorRed, "Please add city name. example: Paris")
		return
	}
	q := os.Args[1]

	res, err := http.Get(fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%v&q=%s", apiKey, q))
	if err != nil {
		log.Fatal(errors.New("error happend from weather api.please try again later"))
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

		message := fmt.Sprintf("time:%s - temp:%.0fC - chance of rain:%.0f%% - text:%s\n",
			date.Format("15:04"),
			hour.TempC,
			hour.ChanceOfRain,
			hour.Condition.Text,
		)

		if hour.ChanceOfRain < 40 {
			util.Colorize(util.ColorRed, message)
		} else {
			util.Colorize(util.ColorReset, message)
		}
	}
}
