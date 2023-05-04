package controllers

import (
	"net/http"

	"github.com/RomManager/server/roms"
	"github.com/gin-gonic/gin"
)

func Sync(c *gin.Context) {
	roms.SyncRomFiles()

	c.JSON(http.StatusOK, gin.H{"message": "Started syncing..."})
}
