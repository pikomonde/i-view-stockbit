package moviesearch

import (
	"errors"

	"github.com/pikomonde/i-view-stockbit/02_movie_searcher/entity"
)

// GetMoviesBySearchWord is used to get omdb data by search word
func (s *ServiceMovieSearch) GetMoviesBySearchWord(search string, page int64) (entity.SearchResult, error) {
	res, err := s.RepositoryMovieSearch.GetMoviesBySearchWord(search, page)
	if err != nil {
		return entity.SearchResult{}, errors.New(ErrServerError)
	}

	// Error message from client
	if res.Error != nil {
		return entity.SearchResult{}, res.Error
	}

	return res, nil
}

// SaveMoviesBySearchWord is used to save the movie data based on search word
func (s *ServiceMovieSearch) SaveMoviesBySearchWord(searchRes entity.SearchResult) error {
	err := s.RepositoryMovieSearchLog.CreateMovieLogQuery(searchRes)
	if err != nil {
		return errors.New(ErrResultNotSaved)
	}

	return nil
}

// GetAndSaveMoviesBySearchWord is used to get movie data and save it
func (s *ServiceMovieSearch) GetAndSaveMoviesBySearchWord(search string, page int64) error {
	res, errMsg := s.GetMoviesBySearchWord(search, page)
	if errMsg != nil {
		return errMsg
	}

	errMsg = s.SaveMoviesBySearchWord(res)
	if errMsg != nil {
		return errMsg
	}

	return nil
}
