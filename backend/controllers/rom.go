package controllers

import (
	"net/http"
	"strconv"

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

	c.JSON(http.StatusOK, gin.H{"data": roms})
}

// Get a specific rom by ID --> /api/rom/:id
func GetRom(c *gin.Context) {
	romID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.DoError(c, http.StatusBadRequest, err)
		return
	}

	rom, err := models.GetRomByID(uint(romID))
	if err != nil {
		utils.DoError(c, http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rom})
}
