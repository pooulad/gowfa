package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/pooulad/gowfa/model"
	"github.com/pooulad/gowfa/pkg/readFlag"
	util "github.com/pooulad/gowfa/util/colorize"
)

func Init(flags *readFlag.FlagReturns ,apiKey string) {
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

	var weather model.Weather

	err = json.Unmarshal(body, &weather)
	if err != nil {
		log.Fatal(err)
	}

	location, current, hours := weather.Location, weather.Current, weather.Forecast.ForecastDay[0].Hour

	CurrentMessage := fmt.Sprintf("Name:%s - Country:%s - Current temp:%.0fC - Text:%s - Lat & Lon:%.2f-%.2f - Localtime:%s\n",
		location.Name,
		location.Coutry,
		current.TempC,
		current.Condition.Text,
		location.Lat,
		location.Lon,
		location.LocalTime,
	)
	util.Colorize(util.ColorGreen, CurrentMessage)

	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)

		message := fmt.Sprintf("Time:%s - Temp:%.0fC - Chance of rain:%.0f%% - Text:%s\n",
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
