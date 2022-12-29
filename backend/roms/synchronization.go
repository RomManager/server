package roms

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/vallezw/RomManager/backend/config"
	"github.com/vallezw/RomManager/backend/gridapi"
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

// TODO: Write method for check if files are missing
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
		// Check if a rom with the filepath already exists; if not -> Save to roms db
		if !models.CheckRomExistsByFilepath(path) {
			// searchedForName is either just cropped or a found result by the API
			game, croppedName := makeAPISearchCall(info.Name())
			steamGridDBID := 0
			searchedForName := croppedName
			var releaseDate time.Time
			if croppedName == "" {
				c.Printf("Found game: %v in SteamGridDB API\n", game.Name)

				steamGridDBID = game.ID
				searchedForName = game.Name
				fmt.Printf("This is the releaseDate value %v\n", game.ReleaseDate)
				if game.ReleaseDate != 0 {
					releaseDate = time.Unix(game.ReleaseDate, 0)
				}
			}

			rom := models.Rom{
				Name:          searchedForName,
				Filepath:      path,
				Emulator:      emulator.FolderName,
				SteamGridDBID: steamGridDBID,
				ReleaseDate:   releaseDate,
			}
			rom.SaveRom()
			c.Printf("Didn't find in DB --- Created a new entry --- Continuing...\n")
			return nil
		}

		c.Printf("Found in DB --- Continuing...\n")
		return nil
	})
}

/*
Takes the file name -> crops it by . (e.g. .txt) ->
searches api -> if found returns found game and empty string if not returns empty game and cropped name
*/
func makeAPISearchCall(filename string) (gridapi.GameResponse, string) {
	croppedName := strings.Split(filename, ".")[-0]

	game, _ := gridapi.SearchForGame(croppedName)
	// No game found or err
	if (game == gridapi.GameResponse{}) {
		return gridapi.GameResponse{}, croppedName
	}

	return game, ""
}
