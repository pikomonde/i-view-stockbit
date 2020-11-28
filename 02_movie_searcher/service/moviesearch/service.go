package moviesearch

import (
	"github.com/pikomonde/i-view-stockbit/02_movie_searcher/config"
	"github.com/pikomonde/i-view-stockbit/02_movie_searcher/repository"
)

// ServiceMovieSearch contains services for OMDB
type ServiceMovieSearch struct {
	Config                   config.Config
	RepositoryMovieSearch    repository.MovieSearchRepository
	RepositoryMovieSearchLog repository.MovieSearchLogRepository
}

// Opt is options of the services
type Opt struct {
	Config                   config.Config
	RepositoryMovieSearch    repository.MovieSearchRepository
	RepositoryMovieSearchLog repository.MovieSearchLogRepository
}

// New returns ServiceMovieSearch
func New(opt Opt) *ServiceMovieSearch {
	return &ServiceMovieSearch{
		Config:                   opt.Config,
		RepositoryMovieSearch:    opt.RepositoryMovieSearch,
		RepositoryMovieSearchLog: opt.RepositoryMovieSearchLog,
	}
}
