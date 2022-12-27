package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vallezw/RomManager/backend/controllers"
	"github.com/vallezw/RomManager/backend/models"
)

func main() {
	// Load env file (to be optimized)
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	r := gin.Default()

	models.ConnectDatabase()

	public := r.Group("/api")

	public.GET("/books", controllers.FindBooks)
	public.POST("/books", controllers.CreateBook)
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	r.Run()
}
