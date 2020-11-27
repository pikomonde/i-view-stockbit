package httpapi

import "github.com/pikomonde/i-view-stockbit/02_movie_searcher/entity"

const (
	urlBase             = "www.omdbapi.com/"
	urlDataBySearchWord = "http://" + urlBase + "?apikey=%s&s=%s&page=%d"
)

// SearchResultResponse is a response struct from omdb
type SearchResultResponse struct {
	Search   []entity.MovieInfo `json:"Search"`
	Total    string             `json:"totalResults"`
	Response string             `json:"Response"`
	Error    string             `json:"Error"`
}
