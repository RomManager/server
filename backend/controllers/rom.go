package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vallezw/RomManager/backend/models"
	"github.com/vallezw/RomManager/backend/utils"
)

func GetAllRoms(c *gin.Context) {
	roms, err := models.GetAllRoms()
	if err != nil {
		utils.DoError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": roms})
}
