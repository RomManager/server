package gridapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/RomManager/server/config"
	"github.com/fatih/color"
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

// Unmarshal the data into the structs/types given in response_types.go
func UnmarshalData(resp *http.Response, target interface{}) error {
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &target); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func SetupGridAPI() {
	client = httpClient{
		c:        http.Client{},
		apiToken: config.Config().SteamGridDBAPIKey,
	}

	err := healthCheck()
	if err != nil {
		// TODO: Make a fix that it sets GridAPIEnabled to false in Config
		return
	}
}

// Check if you can connect to the API
func healthCheck() error {
	// Make a test on Super Mario Sunshine
	resp, err := client.Get("/games/id/34899")
	if err != nil {
		fmt.Println(err)
		return err
	}

	if resp.StatusCode != 200 {
		color.Red("Problem connecting with the SteamGridDB database, please check your given API key")
		return errors.New("problem connecting with the SteamGridDB database, please check your given API key")
	}

	color.Green("Connection with SteamGridDB API has been successful")
	return nil
}
