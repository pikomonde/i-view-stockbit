package moviesearch

import (
	"github.com/pikomonde/i-view-stockbit/02_movie_searcher/config"
	"github.com/pikomonde/i-view-stockbit/02_movie_searcher/repository"
)

// ServiceOMDB contains services for OMDB
type ServiceOMDB struct {
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

// New returns ServiceOMDB
func New(opt Opt) *ServiceOMDB {
	return &ServiceOMDB{
		Config:                   opt.Config,
		RepositoryMovieSearch:    opt.RepositoryMovieSearch,
		RepositoryMovieSearchLog: opt.RepositoryMovieSearchLog,
	}
}
