package models

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("data/database.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // Silence the logger
	})

	if err != nil {
		panic("Failed to connect to database!")
	}

	// Migrate the models to the database
	err = database.AutoMigrate(&Rom{}, &User{})
	if err != nil {
		return
	}

	DB = database

	fmt.Println("Connected to database")
}
