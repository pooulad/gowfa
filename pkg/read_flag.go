package pkg

import (
	"errors"
	"flag"
)

type FlagReturns struct {
	city string
	api  bool
}

func ReadFlag() (*FlagReturns, error) {
	fr := FlagReturns{}
	flag.StringVar(&fr.city, "c", "", "get city name")
	flag.BoolVar(&fr.api, "a", false, "choose app type(api or cli)")
	flag.Parse()

	if fr.city == "" {
		return nil, errors.New("please enter city name for getting data [-c or --c] => example : -c=Paris")
	}

	return &fr, nil
}
