package service

import (
	"github.com/georgedolee/amindi/internal/apiclient"
	"github.com/georgedolee/amindi/internal/model"
)

type ForecastService struct {
	client *apiclient.Client
}

func NewForecastService(client *apiclient.Client) *ForecastService {
	return &ForecastService{client: client}
}

func (s *ForecastService) Get(location string, days int) (*model.WeatherForecast, error) {
	return s.client.FetchForecast(location, days)
}
