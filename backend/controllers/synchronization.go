package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vallezw/RomManager/backend/roms"
)

func Sync(c *gin.Context) {
	roms.SyncRomFiles()

	c.JSON(http.StatusOK, gin.H{"message": "Started syncing..."})
}
