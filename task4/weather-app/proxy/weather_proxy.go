package proxy

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type WeatherProxy struct{}

type OpenMeteoResponse struct {
	CurrentWeather struct {
		Temperature float64 `json:"temperature"`
		WeatherCode int     `json:"weathercode"`
	} `json:"current_weather"`
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
}

func NewWeatherProxy() *WeatherProxy {
	return &WeatherProxy{}
}

// Get weather by city name (maintains backward compatibility)
func (wp *WeatherProxy) GetWeather(city string) (float64, string, int, error) {
	coordinates := map[string]struct {
		Lat  float64
		Long float64
	}{
		"tokyo":    {35.6762, 139.6503},
		"warsaw":   {52.2297, 21.0122},
		"berlin":   {52.5200, 13.4050},
		"new york": {40.7128, -74.0060},
	}

	loc, ok := coordinates[city]
	if !ok {
		return 0, "", 0, errors.New("city not in predefined list")
	}

	return wp.GetWeatherByCoords(loc.Lat, loc.Long)
}

// New method for coordinate-based lookup
func (wp *WeatherProxy) GetWeatherByCoords(lat, long float64) (float64, string, int, error) {
	url := fmt.Sprintf(
		"https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current_weather=true",
		lat,
		long,
	)

	resp, err := http.Get(url)
	if err != nil {
		return 0, "", 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, "", 0, errors.New("failed to fetch weather data")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, "", 0, err
	}

	var weatherResp OpenMeteoResponse
	err = json.Unmarshal(body, &weatherResp)
	if err != nil {
		return 0, "", 0, err
	}

	// Map weather code to conditions
	conditions := map[int]string{
		0:  "Clear",
		1:  "Partly Cloudy",
		2:  "Cloudy",
		3:  "Overcast",
		45: "Fog",
		61: "Rain",
		80: "Showers",
	}

	condition := "Unknown"
	if val, ok := conditions[weatherResp.CurrentWeather.WeatherCode]; ok {
		condition = val
	}

	// Use location name from API or "Unknown"
	cityName := weatherResp.Location.Name
	if cityName == "" {
		cityName = "Unknown Location"
	}

	return weatherResp.CurrentWeather.Temperature, condition, 0, nil
}