package controllers

import (
	"fmt"

	"github.com/RomManager/server/config"
	"github.com/RomManager/server/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() {
	if !config.Config().DebugEnabled {
		fmt.Println("Running server in production mode...")
		gin.SetMode(gin.ReleaseMode)
	} else {
		fmt.Println("Running server in debug mode...")
	}

	r := gin.Default()

	// Use CORS when debug is enabled
	if config.Config().DebugEnabled {
		r.Use(middlewares.CORSMiddleware())
		fmt.Println("CORS are enabled")
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
	public.GET("/rom/:id/file", GetRomFile)
	public.GET("/rom/:id/grid", GetGridFile)

	// Synchronization
	public.GET("/sync", Sync)

	// Static files
	// public.Static("/files", config.Config().DataPath)

	runServer(r)
}

func runServer(router *gin.Engine) {
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
