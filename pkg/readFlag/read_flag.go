package readFlag

import (
	"errors"
	"flag"
)

type FlagReturns struct {
	City string
	Api  bool
}

func ReadFlag() (*FlagReturns, error) {
	fr := FlagReturns{}
	flag.StringVar(&fr.City, "c", "", "get city name")
	flag.BoolVar(&fr.Api, "a", false, "choose app type(api or cli)")
	flag.Parse()

	if fr.City == "" {
		return nil, errors.New("please enter city name for getting data [-c or --c] => example : -c=Paris")
	}

	return &fr, nil
}
