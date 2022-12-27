package main

import (
	"fmt"

	"github.com/vallezw/RomManager/backend/controllers"
	"github.com/vallezw/RomManager/backend/models"
)

func main() {
	fmt.Printf("Welcome to the RomManager Server")

	// Connect to DB and setup router
	models.ConnectDatabase()
	controllers.SetupRouter()
}
