package moviesearch_test

import (
	"context"
	"testing"

	"github.com/pikomonde/i-view-stockbit/02_movie_searcher/config"
	"github.com/pikomonde/i-view-stockbit/02_movie_searcher/delivery/grpchandler/moviesearch"
	"github.com/pikomonde/i-view-stockbit/02_movie_searcher/repository"
	servMovs "github.com/pikomonde/i-view-stockbit/02_movie_searcher/service/moviesearch"
)

func TestGetAndSaveMoviesBySearch(t *testing.T) {
	cfg := config.New()
	rMovs := repository.NewHTTPAPIMovieSearch(cfg)
	rMovl := repository.NewMySQLMovieSearchLog(cfg)
	sMovs := servMovs.New(servMovs.Opt{
		Config:                   cfg,
		RepositoryMovieSearch:    rMovs,
		RepositoryMovieSearchLog: rMovl,
	})
	dMovs := moviesearch.New(moviesearch.Opt{
		ServiceMovieSearch: sMovs,
	})

	// Positive test
	dMovs.GetAndSaveMoviesBySearch(context.Background(), &moviesearch.MessageRequest{
		Searchword: "Batman",
		Pagination: 2,
	})

	// Negative test
	dMovs.GetAndSaveMoviesBySearch(context.Background(), &moviesearch.MessageRequest{
		Searchword: "-",
		Pagination: 0,
	})
}

func TestProto(t *testing.T) {
	cfg := config.New()
	rMovs := repository.NewHTTPAPIMovieSearch(cfg)
	rMovl := repository.NewMySQLMovieSearchLog(cfg)
	sMovs := servMovs.New(servMovs.Opt{
		Config:                   cfg,
		RepositoryMovieSearch:    rMovs,
		RepositoryMovieSearchLog: rMovl,
	})
	dMovs := moviesearch.New(moviesearch.Opt{
		ServiceMovieSearch: sMovs,
	})

	req := &moviesearch.MessageRequest{
		Searchword: "Batman",
		Pagination: 2,
	}
	resp, _ := dMovs.GetAndSaveMoviesBySearch(context.Background(), req)

	req.Descriptor()
	req.ProtoMessage()
	req.ProtoReflect()
	req.GetPagination()
	req.GetSearchword()
	req.String()
	req.Reset()

	resp.Descriptor()
	resp.ProtoMessage()
	resp.ProtoReflect()
	resp.GetResponse()
	resp.String()
	resp.Reset()

}
