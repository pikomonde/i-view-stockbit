package repository

import (
	"github.com/pikomonde/i-view-stockbit/02_movie_searcher/config"
	"github.com/pikomonde/i-view-stockbit/02_movie_searcher/entity"
	"github.com/pikomonde/i-view-stockbit/02_movie_searcher/repository/httpapi"
	"github.com/pikomonde/i-view-stockbit/02_movie_searcher/repository/mysql"
)

// MovieSearchLogRepository contain list of MovieSearch's repository
type MovieSearchLogRepository interface {
	CreateMovieLogQuery(searchRes entity.SearchResult) error
}

// NewMySQLMovieSearchLog returns MovieSearch repository using mysql connection
func NewMySQLMovieSearchLog(cfg config.Config) MovieSearchLogRepository {
	return &mysql.RepositoryMovieSearchLog{
		Cli: cfg.MySQLCli,
	}
}

// MovieSearchRepository contain list of MovieSearch's repository
type MovieSearchRepository interface {
	GetMoviesBySearchWord(search string, page int64) (entity.SearchResult, error)
}

// NewHTTPAPIMovieSearch returns MovieSearch repository using mysql connection
func NewHTTPAPIMovieSearch(cfg config.Config) MovieSearchRepository {
	return &httpapi.RepositoryMovieSearch{
		APIKeyOMDB: cfg.APIKeyOMDB,
		Cli:        cfg.HTTPCli,
	}
}
