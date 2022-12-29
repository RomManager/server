package gridapi

type GameResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ReleaseDate string `json:"release_date"`
	// types       []string
	// verified    bool
}

type SingleDataResponse struct {
	Success bool         `json:"success"`
	Game    GameResponse `json:"data"`
}

// For an array of data, used by search for example
type ArrayDataResponse struct {
	Success   bool           `json:"success"`
	GameArray []GameResponse `json:"data"`
}
