package apiclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/georgedolee/amindi/internal/model"
)

const (
	host    = "http://api.weatherapi.com/"
	version = "v1/"
	format  = "forecast.json"
	baseURL = host + version + format
)

type Client struct {
	APIKey string
	Client *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		APIKey: apiKey,
		Client: &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *Client) FetchForecast(location string, days int) (*model.WeatherForecast, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %w", err)
	}

	q := u.Query()
	q.Set("key", c.APIKey)
	q.Set("q", location)
	q.Set("days", fmt.Sprint(days))
	q.Set("aqi", "no")
	q.Set("alerts", "no")
	u.RawQuery = q.Encode()

	var (
		forecast      model.WeatherForecast
		errorResponse model.ErrorResponse
	)

	if err := c.doGet(u.String(), &forecast, &errorResponse); err != nil {
		return nil, err
	}

	return &forecast, nil
}

func (c *Client) doGet(fullURL string, successObj, errorObj any) error {
	res, err := c.Client.Get(fullURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	dec := json.NewDecoder(res.Body)

	if res.StatusCode != http.StatusOK {
		if err := dec.Decode(errorObj); err != nil {
			return fmt.Errorf("API error (status %d) and decode error: %w", res.StatusCode, err)
		}

		if errorResponse, ok := errorObj.(*model.ErrorResponse); ok {
			return fmt.Errorf("API error: %s", errorResponse.Error.Message)
		}

		return fmt.Errorf("API error: status %d", res.StatusCode)
	}

	if err := dec.Decode(successObj); err != nil {
		return fmt.Errorf("failed to decode API response: %w", err)
	}

	return nil
}
