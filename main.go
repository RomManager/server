package main

import (
	"fmt"

	"github.com/RomManager/server/config"
	"github.com/RomManager/server/controllers"
	"github.com/RomManager/server/gridapi"
	"github.com/RomManager/server/models"
	"github.com/RomManager/server/roms"
)

func main() {
	err := RunServer()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func RunServer() error {
	fmt.Println("Welcome to the RomManager server")
	fmt.Printf("Loaded config with SteamGridDBEnabled: %v\n", config.Config().GridAPIEnabled)

	if config.Config().GridAPIEnabled {
		gridapi.SetupGridAPI()
	}

	roms.SetupDirectories()

	// Connect to DB and setup router
	models.ConnectDatabase()
	err := controllers.SetupRouter()
	return err
}
