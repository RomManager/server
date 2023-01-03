package controllers

import (
	"net/http"

	"github.com/RomManager/server/backend/gridapi"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Welcome to the RomManager API"})
}

func RunMethod(c *gin.Context) {
	grid, _ := gridapi.GetGameGrid(34899)
	c.JSON(http.StatusOK, gin.H{"data": grid})
}
