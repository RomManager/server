package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vallezw/RomManager/backend/gridapi"
)

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Welcome to the RomManager API"})
}

func RunMethod(c *gin.Context) {
	gridapi.SearchForGame("Celeste")
	c.JSON(http.StatusOK, gin.H{"data": "I ran the method"})
}
