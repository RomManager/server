package main

import (
	"fmt"

	"github.com/vallezw/RomManager/backend/controllers"
	"github.com/vallezw/RomManager/backend/models"
)

func main() {
	fmt.Println("Welcome to the RomManager server")

	// Connect to DB and setup router
	models.ConnectDatabase()
	controllers.SetupRouter()
}
