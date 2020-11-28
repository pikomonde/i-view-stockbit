package delivery_test

import (
	"testing"

	"github.com/pikomonde/i-view-stockbit/02_movie_searcher/config"
	"github.com/pikomonde/i-view-stockbit/02_movie_searcher/delivery"
	"github.com/pikomonde/i-view-stockbit/02_movie_searcher/repository"
	servMovs "github.com/pikomonde/i-view-stockbit/02_movie_searcher/service/moviesearch"
)

func TestDelivery(t *testing.T) {
	cfg := config.New()
	rMovs := repository.NewHTTPAPIMovieSearch(cfg)
	rMovl := repository.NewMySQLMovieSearchLog(cfg)
	sMovs := servMovs.New(servMovs.Opt{
		Config:                   cfg,
		RepositoryMovieSearch:    rMovs,
		RepositoryMovieSearchLog: rMovl,
	})

	// New & Start
	delv := delivery.New(delivery.Opt{
		ServiceMovieSearch: sMovs,
	})
	delv.Start()
}
