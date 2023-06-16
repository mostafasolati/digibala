package db

import (
	"currency/models"
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	database, err := gorm.Open(sqlite.Open("currency.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&models.Currency{})
	if err != nil {
		return
	}

	DB = database
}
