package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vallezw/RomManager/backend/controllers"
	"github.com/vallezw/RomManager/backend/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)

	r.Run()
}
