package controllers

import (
	"github.com/RomManager/server/backend/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() {
	r := gin.Default()

	public := r.Group("/api")
	protected := r.Group("/api")
	protected.Use(middlewares.JwtAuthMiddleware())

	public.GET("/", Home)

	// For testing (Should be removed afterwards)
	public.GET("/run_method", RunMethod)

	// Authentication
	public.POST("/register", Register)
	public.POST("/login", Login)
	protected.GET("/user", CurrentUser)

	// Roms
	public.GET("/roms", GetAllRoms)
	public.GET("/rom/:id", GetRom)

	// Synchronization
	public.GET("/sync", Sync)

	runServer(r)
}

func runServer(router *gin.Engine) {
	router.Run(":8080")
}
