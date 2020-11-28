package httphandler_test

import (
	"testing"
	"time"

	"github.com/pikomonde/i-view-stockbit/02_movie_searcher/config"
	"github.com/pikomonde/i-view-stockbit/02_movie_searcher/delivery/httphandler"
	"github.com/pikomonde/i-view-stockbit/02_movie_searcher/repository"
	servMovs "github.com/pikomonde/i-view-stockbit/02_movie_searcher/service/moviesearch"
)

func TestStart(t *testing.T) {
	cfg := config.New()
	rMovs := repository.NewHTTPAPIMovieSearch(cfg)
	rMovl := repository.NewMySQLMovieSearchLog(cfg)
	sMovs := servMovs.New(servMovs.Opt{
		Config:                   cfg,
		RepositoryMovieSearch:    rMovs,
		RepositoryMovieSearchLog: rMovl,
	})

	// New & Start
	hh := httphandler.New(httphandler.Opt{
		ServiceMovieSearch: sMovs,
	})
	hh.Start()

	time.Sleep(500 * time.Millisecond)

	// New & Start (error)
	hhh := httphandler.New(httphandler.Opt{
		ServiceMovieSearch: sMovs,
	})
	hhh.Start()
}
