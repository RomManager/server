package main

import (
	"fmt"

	"github.com/vallezw/RomManager/backend/controllers"
	"github.com/vallezw/RomManager/backend/models"
	"github.com/vallezw/RomManager/backend/roms"
)

func main() {
	fmt.Println("Welcome to the RomManager server")

	roms.SetupDirectories()

	// Connect to DB and setup router
	models.ConnectDatabase()
	controllers.SetupRouter()
}
