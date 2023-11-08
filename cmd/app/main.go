package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/pooulad/gowfa/cmd/cli"
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

	fmt.Println(os.Args)
	
	useApi := flag.Bool("api", false, "choose version of programm(api and cli mode)")
	flag.Parse()

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
	fmt.Println(*useApi)

	if *useApi {
		return
	}
	cli.CliInit(body)
}
