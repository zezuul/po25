package controllers

import (
	"fmt"
	"net/http"
	"weather-app/models"
	"weather-app/proxy"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type WeatherController struct {
	DB    *gorm.DB
	Proxy *proxy.WeatherProxy
}

func NewWeatherController(db *gorm.DB, proxy *proxy.WeatherProxy) *WeatherController {
	return &WeatherController{DB: db, Proxy: proxy}
}

type WeatherRequest struct {
	City string  `json:"city" query:"city"`
	Lat  float64 `json:"lat" query:"lat"`
	Long float64 `json:"long" query:"long"`
}

func (wc *WeatherController) GetWeather(c echo.Context) error {
	var req WeatherRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	// Try coordinates first if provided
	if req.Lat != 0 && req.Long != 0 {
		temp, conditions, _, err := wc.Proxy.GetWeatherByCoords(req.Lat, req.Long)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to fetch weather data"})
		}

		// Save to database for future requests
		location := fmt.Sprintf("(%.4f,%.4f)", req.Lat, req.Long)
		weather := models.Weather{
			City:        location,
			Temperature: temp,
			Conditions:  conditions,
		}
		wc.DB.Create(&weather)

		return c.JSON(http.StatusOK, weather)
	}

	// Fall back to city name if provided
	if req.City != "" {
		// First try database
		var weather models.Weather
		result := wc.DB.Where("city = ?", req.City).First(&weather)

		if result.Error == nil {
			return c.JSON(http.StatusOK, weather)
		}

		// If not in DB, use proxy
		temp, conditions, _, err := wc.Proxy.GetWeather(req.City)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to fetch weather data"})
		}

		// Save to database
		newWeather := models.Weather{
			City:        req.City,
			Temperature: temp,
			Conditions:  conditions,
		}
		wc.DB.Create(&newWeather)

		return c.JSON(http.StatusOK, newWeather)
	}

	return c.JSON(http.StatusBadRequest, map[string]string{"error": "city or coordinates required"})
}