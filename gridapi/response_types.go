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

// For the Grids of the Games
type GridResponse struct {
	ID  int    `json:"id"`
	URL string `jsong:"url"`
}

// For an array of grids response
type ArrayGridResponse struct {
	Success   bool           `json:"success"`
	GridArray []GridResponse `json:"data"`
}
