package utils

import (
	"errors"
	"log"
	"os"
)

// Create directory if doesn't exist with proper error handling
func CreateDirIfNotExists(path string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}
