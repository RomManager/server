package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vallezw/RomManager/backend/roms"
)

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Welcome to the RomManager API"})
}

func RunMethod(c *gin.Context) {
	roms.SetupDirectories()
	c.JSON(http.StatusOK, gin.H{"data": "I ran the method"})
}
