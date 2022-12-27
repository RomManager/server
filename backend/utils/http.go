package utils

import (
	"log"

	"github.com/gin-gonic/gin"
)

func DoError(c *gin.Context, status int, err error) {
	log.Printf("%s\n", err.Error())
	c.JSON(status, gin.H{"error": err.Error()})
}
