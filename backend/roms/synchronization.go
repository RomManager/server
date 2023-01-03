package roms

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/RomManager/server/backend/config"
	"github.com/RomManager/server/backend/gridapi"
	"github.com/RomManager/server/backend/models"
	"github.com/RomManager/server/backend/utils"
	"github.com/fatih/color"
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
			rom := createRom(c, path, info.Name(), emulator.FolderName)
			rom.SaveRom()
			c.Printf("Didn't find in DB --- Created a new entry --- Continuing...\n")
			return nil
		}

		c.Printf("Found in DB --- Continuing...\n")
		return nil
	})
}

func createRom(c *color.Color, path string, filename string, emulator string) models.Rom {
	// searchedForName is either just cropped or a found result by the API
	game, grid, croppedName := makeAPISearchCall(filename)
	steamGridDBID := 0
	searchedForName := croppedName
	gridURL := ""
	var releaseDate time.Time
	if croppedName == "" {
		c.Printf("Found game: %v in SteamGridDB API\n", game.Name)

		steamGridDBID = game.ID
		searchedForName = game.Name
		gridURL = grid.URL
		if game.ReleaseDate != 0 {
			releaseDate = time.Unix(game.ReleaseDate, 0)
		}
	}

	rom := models.Rom{
		Name:          searchedForName,
		Filepath:      path,
		Emulator:      emulator,
		SteamGridDBID: steamGridDBID,
		ReleaseDate:   releaseDate,
		GridURL:       gridURL,
	}
	return rom
}

/*
Takes the file name -> crops it by . (e.g. .txt) ->
searches api -> if found returns found game and empty string if not returns empty game and cropped name
@return GameResponse, GridResponse, string
*/
func makeAPISearchCall(filename string) (gridapi.GameResponse, gridapi.GridResponse, string) {
	croppedName := strings.Split(filename, ".")[-0]
	fmt.Println(croppedName)
	game, _ := gridapi.SearchForGame(croppedName)
	// No game found or err
	if (game == gridapi.GameResponse{}) {
		return gridapi.GameResponse{}, gridapi.GridResponse{}, croppedName
	}

	// Get grid
	grid, _ := gridapi.GetGameGrid(game.ID)

	return game, grid, ""
}
