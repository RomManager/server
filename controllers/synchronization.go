package controllers

import (
	"net/http"

	"github.com/RomManager/server/roms"
	"github.com/gin-gonic/gin"
)

func Sync(c *gin.Context) {
	err := roms.SyncRomFiles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error while syncing"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Started syncing..."})
}
