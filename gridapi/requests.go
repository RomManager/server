package gridapi

import (
	"fmt"
	"strconv"
)

// Uses SteamGridDB to search for a game
func SearchForGame(gameName string) (GameResponse, error) {
	resp, err := client.Get("/search/autocomplete/" + gameName)
	if err != nil {
		fmt.Println(err)
		return GameResponse{}, err
	}

	dataRes := new(ArrayDataResponse)

	UnmarshalData(resp, dataRes)

	if len(dataRes.GameArray) == 0 {
		return GameResponse{}, nil
	}

	return dataRes.GameArray[0], nil
}

// Search for a thumbnail of the given game via the game ID
func GetGameGrid(gameID int) (GridResponse, error) {
	resp, err := client.Get("/grids/game/" + strconv.Itoa(gameID))
	if err != nil {
		fmt.Println(err)
		return GridResponse{}, err
	}

	dataRes := new(ArrayGridResponse)

	UnmarshalData(resp, dataRes)

	if len(dataRes.GridArray) == 0 {
		return GridResponse{}, nil
	}

	return dataRes.GridArray[0], nil
}
