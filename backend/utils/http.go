package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/mail"
	"os"

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

func DownloadFileToPath(downloadURL string, filePath string) {
	// Create blank file
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, _ []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	// Put content on file
	resp, err := client.Get(downloadURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	size, _ := io.Copy(file, resp.Body)

	defer file.Close()

	fmt.Printf("Downloaded a file %s with size %d", filePath, size)
}
