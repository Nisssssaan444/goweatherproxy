package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type APIResponse struct {
	Location struct {
		Name    string `json:"name"`
		Region  string `json:"region"`
		Country string `json:"country"`
		Local   string `json:"localtime"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Humidity  int     `json:"humidity"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
		WindKph float64 `json:"wind_kph"`
	} `json:"current"`
}

type WeatherResult struct {
	City        string  `json:"city"`
	Region      string  `json:"region"`
	Country     string  `json:"country"`
	LocalTime   string  `json:"local_time"`
	Temperature float64 `json:"temperature_c"`
	Humidity    int     `json:"humidity"`
	Condition   string  `json:"condition"`
	WindKph     float64 `json:"wind_kph"`
}

func FetchWeather(city string) (*WeatherResult, error) {
	apiKey := os.Getenv("WEATHERAPI_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("WEATHERAPI_KEY environment variable not set")
	}

	url := fmt.Sprintf(
		"https://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no",
		apiKey, city,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API returned %d: %s", resp.StatusCode, string(body))
	}

	var apiResp APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("JSON decode failed: %w", err)
	}

	result := &WeatherResult{
		City:        apiResp.Location.Name,
		Region:      apiResp.Location.Region,
		Country:     apiResp.Location.Country,
		LocalTime:   apiResp.Location.Local,
		Temperature: apiResp.Current.TempC,
		Humidity:    apiResp.Current.Humidity,
		Condition:   apiResp.Current.Condition.Text,
		WindKph:     apiResp.Current.WindKph,
	}

	return result, nil
}
