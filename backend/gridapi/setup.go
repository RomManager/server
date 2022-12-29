package gridapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/fatih/color"
	"github.com/vallezw/RomManager/backend/config"
)

var client httpClient

type httpClient struct {
	c        http.Client
	apiToken string
}

// Url param is in this case just the endpoint starting with a / e.g. /games/id/34899
func (c *httpClient) Get(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", "https://www.steamgriddb.com/api/v2"+url, nil)
	if err != nil {
		return nil, err

	}

	// Add header
	req.Header.Add("Authorization", "Bearer "+c.apiToken)

	return c.c.Do(req)
}

/*
Example:

	{
		"id": 34899,
		"name": "Super Mario Sunshine",
		"release_date": null,
		"types": [],
		"verified": true
	}
*/
type GameResponse struct {
	ID          int `json:"id"`
	Name        string
	ReleaseDate string
	// types       []string
	// verified    bool
}

type DataResponse struct {
	Success bool         `json:"success"`
	Data    GameResponse `json:"data"`
}

func SetupGridAPI() {
	client = httpClient{
		c:        http.Client{},
		apiToken: config.Config().SteamGridDBAPIKey,
	}

	// Health check

	resp, err := client.Get("/games/id/34899")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode == 200 {
		color.Green("Connection with SteamGridDB API has been successful")
	}

	var dataresponse DataResponse

	json.NewDecoder(resp.Body).Decode(&dataresponse)

	fmt.Println(dataresponse)

	fmt.Println(string(body))
}
