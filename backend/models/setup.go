package models

import (
	"github.com/fatih/color"
	"github.com/vallezw/RomManager/backend/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open(config.Config().DataPath+"database.db"), &gorm.Config{
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

	color.Cyan("Made sure data/database.db file exists and connected to database\n")
}
