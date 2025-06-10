package models

import (
	"gorm.io/gorm"
)

type Weather struct {
	gorm.Model
	City        string  `json:"city"`
	Temperature float64 `json:"temperature"`
	Conditions  string  `json:"conditions"`
	Humidity    int     `json:"humidity"`
}

func InitializeDB(db *gorm.DB) error {
	// AutoMigrate creates the table
	err := db.AutoMigrate(&Weather{})
	if err != nil {
		return err
	}

	// Initial data
	initialWeather := []Weather{
		{City: "Warsaw", Temperature: 20.5, Conditions: "Sunny", Humidity: 45},
		{City: "Berlin", Temperature: 18.2, Conditions: "Cloudy", Humidity: 60},
		{City: "Paris", Temperature: 22.1, Conditions: "Partly Cloudy", Humidity: 55},
	}

	// Create initial records
	for _, w := range initialWeather {
		result := db.FirstOrCreate(&w, Weather{City: w.City})
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}