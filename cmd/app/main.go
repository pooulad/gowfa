package main

import (
	"log"
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

	if flags.Api {
		api.Init(flags,apiKey)
	} else {
		cli.Init(flags,apiKey)
	}
}
