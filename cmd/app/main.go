package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/pooulad/gowfa/cmd/api"
	"github.com/pooulad/gowfa/cmd/cli"
	"github.com/pooulad/gowfa/pkg/readFlag"
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

	flags, err := readFlag.ReadFlag()
	if err != nil {
		util.Colorize(util.ColorRed, err.Error())
		return
	}

	res, err := http.Get(fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%v&q=%s", apiKey, flags.City))
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

	if flags.Api {
		api.Init()
	} else {
		cli.Init(body)
	}
}
