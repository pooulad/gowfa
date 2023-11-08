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

