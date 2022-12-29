package main

import (
	"fmt"

	"github.com/vallezw/RomManager/backend/config"
	"github.com/vallezw/RomManager/backend/controllers"
	"github.com/vallezw/RomManager/backend/gridapi"
	"github.com/vallezw/RomManager/backend/models"
	"github.com/vallezw/RomManager/backend/roms"
)

func RunServer() {
	fmt.Println("Welcome to the RomManager server")
	fmt.Printf("Loaded config with SteamGridDBEnabled: %v\n", config.Config().GridAPIEnabled)

	if config.Config().GridAPIEnabled {
		gridapi.SetupGridAPI()
	}

	roms.SetupDirectories()

	// Connect to DB and setup router
	models.ConnectDatabase()
	controllers.SetupRouter()
}
