package main

import (
	"log"
	"os"
	"weather-app/controllers"
	"weather-app/models"
	"weather-app/proxy"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Initialize Echo
	e := echo.New()

	// Create data directory if it doesn't exist
	if err := os.MkdirAll("data", 0755); err != nil {
		log.Fatal("Failed to create data directory:", err)
	}

	// Use file-based SQLite with pure Go driver
	db, err := gorm.Open(sqlite.Open("data/weather.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize models
	err = models.InitializeDB(db)
	if err != nil {
		e.Logger.Fatal("Failed to initialize database:", err)
	}

	// Initialize weather proxy
	weatherProxy := proxy.NewWeatherProxy()

	// Initialize controller
	weatherController := controllers.NewWeatherController(db, weatherProxy)

	// Routes
	e.GET("/weather", weatherController.GetWeather)
	e.POST("/weather", weatherController.GetWeather)

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}