package main

import (
	"fmt"

	"github.com/RomManager/server/backend/config"
	"github.com/RomManager/server/backend/controllers"
	"github.com/RomManager/server/backend/gridapi"
	"github.com/RomManager/server/backend/models"
	"github.com/RomManager/server/backend/roms"
)

func main() {
	RunServer()
}

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
