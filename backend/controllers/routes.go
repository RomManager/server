package controllers

import (
	"fmt"

	"github.com/RomManager/server/backend/config"
	"github.com/RomManager/server/backend/middlewares"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func SetupRouter() {
	if !config.Config().DebugEnabled {
		fmt.Println("Running server in production mode...")
		gin.SetMode(gin.ReleaseMode)
	} else {
		fmt.Println("Running server in debug mode...")
	}

	r := gin.Default()

	if config.Config().DebugEnabled {
		// Use CORS when debug is enabled
		r.Use(cors.Default())
	}

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
