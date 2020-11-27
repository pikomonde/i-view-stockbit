package entity

// MovieInfo contains movie info from omdb
type MovieInfo struct {
	Title     string `json:"Title"`
	Year      string `json:"Year"`
	IMDBID    string `json:"imdbID"`
	Type      string `json:"Type"`
	PosterURL string `json:"Poster"`
}

// SearchResult contains search terms and movie info result
type SearchResult struct {
	Search     string
	Page       int64
	MovieInfos []MovieInfo
	Error      error
}
