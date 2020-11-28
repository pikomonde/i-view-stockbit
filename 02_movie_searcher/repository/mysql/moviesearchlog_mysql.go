package mysql

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pikomonde/i-view-stockbit/02_movie_searcher/entity"
	"github.com/pikomonde/i-view-stockbit/02_movie_searcher/helper/log"
)

// RepositoryMovieSearchLog contains clients and Movie Search Log repositories
type RepositoryMovieSearchLog struct {
	Cli *sqlx.DB
}

// CreateMovieLogQuery is used to add a Movie Log Query
func (r RepositoryMovieSearchLog) CreateMovieLogQuery(searchRes entity.SearchResult) error {
	query := `insert into movie_log_query (search_key, page_key, title, year, imdb_id, type, poster_url, create_time)
	values (?, ?, ?, ?, ?, ?, ?, ?)`

	for _, v := range searchRes.MovieInfos {
		_, err := r.Cli.Exec(query, searchRes.Search, searchRes.Page,
			v.Title, v.Year, v.IMDBID, v.Type, v.PosterURL, time.Now(),
		)
		if err != nil {
			log.Error(log.Fields{
				"searchRes": fmt.Sprintf("%+v", searchRes),
			}, err)
			return err
		}
	}

	return nil
}
