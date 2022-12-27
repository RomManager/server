package utils

import (
	"log"
	"net/mail"

	"github.com/gin-gonic/gin"
)

func DoError(c *gin.Context, status int, err error) {
	log.Printf("%s\n", err.Error())
	c.JSON(status, gin.H{"error": err.Error()})
}

func IsValidMailAddress(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
