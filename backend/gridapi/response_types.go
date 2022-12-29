package gridapi

type GameResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ReleaseDate int64  `json:"release_date"` // Given as UNIX timestamp
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
