package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vallezw/RomManager/backend/config"
	"github.com/vallezw/RomManager/backend/controllers"
	"github.com/vallezw/RomManager/backend/middlewares"
	"github.com/vallezw/RomManager/backend/models"
)

func main() {

	fmt.Printf("Welcome to %v", config.Config().Version)

	r := gin.Default()

	models.ConnectDatabase()

	public := r.Group("/api")
	protected := r.Group("/api")
	protected.Use(middlewares.JwtAuthMiddleware())

	public.GET("/", controllers.Home)
	public.GET("/books", controllers.FindBooks)
	public.POST("/books", controllers.CreateBook)
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected.GET("/user", controllers.CurrentUser)

	r.Run(":8080")
}
