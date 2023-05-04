package controllers

import (
	"net/http"

	"github.com/RomManager/server/backend/roms"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Welcome to the RomManager API"})
}

func RunMethod(c *gin.Context) {
	roms.DownloadGridImg("https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/3eed6d27b0b2dd008c1be88cce8245fc.png", "nds", "")
	c.JSON(http.StatusOK, gin.H{"data": "te"})
}
