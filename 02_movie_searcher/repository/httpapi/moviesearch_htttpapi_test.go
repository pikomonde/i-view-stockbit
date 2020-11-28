package httpapi_test

import (
	"testing"

	"github.com/pikomonde/i-view-stockbit/02_movie_searcher/config"
	"github.com/pikomonde/i-view-stockbit/02_movie_searcher/repository"
)

func TestGetMoviesBySearchWord(t *testing.T) {
	cfg := config.New()
	rMovs := repository.NewHTTPAPIMovieSearch(cfg)

	// Positive
	rMovs.GetMoviesBySearchWord("Batman", 2)

	// Negative
	rMovs.GetMoviesBySearchWord("-", 0)
}
