package controllers

import (
	"net/http"
	"strconv"

	"github.com/RomManager/server/backend/models"
	"github.com/RomManager/server/backend/utils"
	"github.com/gin-gonic/gin"
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

/*
 * Serve rom file
 * /rom/:id/file
 */
func GetRomFile(c *gin.Context) {
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

	c.File(rom.Filepath)
}

/*
 * Serve grid files
 * /rom/:id/grid
 */
func GetGridFile(c *gin.Context) {
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

	c.File(rom.GridFilepath)
}
