package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/pooulad/gowfa/model"
)

func Init() {
	http.HandleFunc("/forecast", forecastHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func forecastHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		json.NewEncoder(w).Encode("error")
		return
	}
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	apiKey := os.Getenv("WEATHER_API_KEY")

	q:= r.URL.Query().Get("city")
	if q == "" {
		json.NewEncoder(w).Encode(map[string]string{"error": "please send city name in parameter"})
		return
	}

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

	json.NewEncoder(w).Encode(weather)
}
