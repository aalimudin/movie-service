package model

type (
	MovieList struct {
		Search       []MovieListItem `json:"Search"`
		TotalResults string          `json:"totalResults"`
		Response     string          `json:"Response"`
	}

	MovieListItem struct {
		Title  string `json:"Title"`
		Year   string `json:"Year"`
		ImdbID string `json:"imdbID"`
		Type   string `json:"Type"`
		Poster string `json:"Poster"`
	}
)
