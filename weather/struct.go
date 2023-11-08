package weather

type Wheather struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
	Forecast Forecast `json:"forecast"`
}

type Location struct {
	Name      string `json:"name"`
	Region    string `json:"region"`
	Coutry    string `json:"country"`
	Lat       string `json:"lat"`
	Lon       string `json:"lon"`
	LocalTime string `json:"localtime"`
}

type Current struct {
	TempC     float64   `json:"temp_c"`
	Condition Condition `json:"condition"`
}

type Condition struct {
	Text string `json:"text"`
}

type Forecast struct {
	ForecastDay []ForecastDay `json:"forecastday"`
}

type ForecastDay struct {
	Hour []Hour `json:"hour"`
}

type Hour struct {
	TimeEpoch   int64     `josn:"time_epoch"`
	TimeC       int64     `josn:"time_c"`
	Condition   Condition `josn:"condition"`
	ChainOfRain float64   `josn:"chain_of_rain"`
}
