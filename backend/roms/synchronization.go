package roms

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/vallezw/RomManager/backend/config"
	"github.com/vallezw/RomManager/backend/models"
	"github.com/vallezw/RomManager/backend/utils"
)

func SetupDirectories() {
	// Create data directory
	utils.CreateDirIfNotExists(config.Config().DataPath)

	// Run through the emulator list and create a directory for each emulator
	for _, emulator := range EmulatorList {
		utils.CreateDirIfNotExists(config.Config().DataPath + emulator.FolderName)
		color.Cyan("Made sure directory exists: %v \n", emulator.Name)
	}
	color.Blue("All emulator folders are present in data/roms")
}

// TODO: Write method for missing files
func SyncRomFiles() {

	dataPath := config.Config().DataPath

	if config.Config().DataPath == "data/" {
		currentDirectory, _ := os.Getwd()

		dataPath = currentDirectory + "/" + "data/"
	}

	// Iterate through the each emulator folder
	for _, emulator := range EmulatorList {
		emulatorPath := dataPath + emulator.FolderName
		walkThroughDir(emulatorPath, emulator)
	}
}

// TODO: Make log process a bit cleaner
func walkThroughDir(path string, emulator Emulator) {
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		c := color.New(color.FgCyan)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		// Ignore the directories
		if info.IsDir() {
			return nil
		}

		c.Printf("Found '%s' --- Checking for DB entry --- ", info.Name())

		// Check if a rom with the filepath already exists
		if !models.CheckRomExistsByFilepath(path) {
			rom := models.Rom{
				Name:     info.Name(), // TODO: Make api call or something to get proper name
				Filepath: path,
				Emulator: emulator.FolderName,
			}
			rom.SaveRom()
			c.Printf("Didn't find in DB --- Created a new entry --- Continuing...\n")
			return nil
		}

		c.Printf("Found in DB --- Continuing...\n")
		return nil
	})
}
