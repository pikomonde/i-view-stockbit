package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/pikomonde/i-view-stockbit/02_movie_searcher/config"
	"github.com/pikomonde/i-view-stockbit/02_movie_searcher/delivery"
	"github.com/pikomonde/i-view-stockbit/02_movie_searcher/repository"
	servMovs "github.com/pikomonde/i-view-stockbit/02_movie_searcher/service/moviesearch"
)

func main() {
	cfg := config.New()

	// Repository
	rMovs := repository.NewHTTPAPIMovieSearch(cfg)
	rMovl := repository.NewMySQLMovieSearchLog(cfg)

	// Service
	sMovs := servMovs.New(servMovs.Opt{
		Config:                   cfg,
		RepositoryMovieSearch:    rMovs,
		RepositoryMovieSearchLog: rMovl,
	})
	err := sMovs.GetAndSaveMoviesBySearchWord("Batman", 2)
	if err != nil {
		// fmt.Println(err)
		return
	}

	// Delivery
	delv := delivery.New(delivery.Opt{
		ServiceMovieSearch: sMovs,
	})
	delv.Start()

	term := make(chan os.Signal)
	signal.Notify(term, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-term:
		fmt.Println("Exiting gracefully...")
	}
}
