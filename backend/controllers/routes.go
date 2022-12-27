package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/vallezw/RomManager/backend/middlewares"
)

func SetupRouter() {
	r := gin.Default()

	public := r.Group("/api")
	protected := r.Group("/api")
	protected.Use(middlewares.JwtAuthMiddleware())

	public.GET("/", Home)
	public.GET("/books", FindBooks)
	public.POST("/books", CreateBook)
	public.POST("/register", Register)
	public.POST("/login", Login)

	protected.GET("/user", CurrentUser)

	runServer(r)
}

func runServer(router *gin.Engine) {
	router.Run(":8080")
}
