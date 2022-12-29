package gridapi

import "fmt"

// Uses SteamGridDB to search for a game
func SearchForGame(gameName string) (GameResponse, error) {
	// Make a test on Mario Sunshine
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
