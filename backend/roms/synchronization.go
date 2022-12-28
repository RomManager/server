package roms

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/vallezw/RomManager/backend/config"
	"github.com/vallezw/RomManager/backend/utils"
)

func SetupDirectories() {
	// Create data directory
	utils.CreateDirIfNotExists(config.Config().DataPath)

	// Run through the emulator list and create a directory for each emulator
	for _, emulator := range EmulatorList {
		utils.CreateDirIfNotExists(config.Config().DataPath + emulator.FolderName)
		color.Cyan("Checked directory for: %v \n", emulator.Name)
	}
	color.Blue("All emulator folders are present in data/roms")

	SyncRomFiles()
}

func SyncRomFiles() {

	dataPath := config.Config().DataPath

	if config.Config().DataPath == "data/" {
		currentDirectory, _ := os.Getwd()

		dataPath = currentDirectory + "data/"
	}

	// Iterate through the each emulator folder
	for _, emulator := range EmulatorList {
		emulatorPath := dataPath + emulator.FolderName
		walkThroughDir(emulatorPath)
	}
}

func walkThroughDir(path string) {
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		fmt.Printf("File Name: %s\n", info.Name())
		return nil
	})
}
