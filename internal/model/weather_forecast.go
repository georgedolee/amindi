package model

type WeatherForecast struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC     float64   `json:"temp_c"`
		TempF     float64   `json:"temp_f"`
		Condition condition `json:"condition"`
	} `json:"current"`
	Forecast struct {
		ForecastDay []ForecastDay `json:"forecastday"`
	} `json:"forecast"`
}

type condition struct {
	Text string `json:"text"`
	Code int16  `json:"code"`
}

type ForecastDay struct {
	DateEpoch int `json:"date_epoch"`
	Day       struct {
		AvgTempC  float64   `json:"avgtemp_c"`
		AvgTempF  float64   `json:"avgtemp_f"`
		Condition condition `json:"condition"`
	}
}
