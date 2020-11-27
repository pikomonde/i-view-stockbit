package httpapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pikomonde/i-view-stockbit/02_movie_searcher/entity"
	"github.com/pikomonde/i-view-stockbit/02_movie_searcher/helper/log"
)

// RepositoryMovieSearch contains clients and Movie Search repositories
type RepositoryMovieSearch struct {
	APIKeyOMDB string
	Cli        *http.Client
}

// GetMoviesBySearchWord getting movie info from external source
func (r *RepositoryMovieSearch) GetMoviesBySearchWord(search string, page int64) (entity.SearchResult, error) {
	url := fmt.Sprintf(urlDataBySearchWord,
		r.APIKeyOMDB,
		search,
		page,
	)
	searchResultResponse := SearchResultResponse{}
	movieInfos := make([]entity.MovieInfo, 0)

	// Getting Data
	resp, err := r.Cli.Get(url)
	if err != nil {
		return entity.SearchResult{}, log.Error(log.Fields{
			"search": search,
			"page":   page,
			"url":    url,
		}, err)
	}
	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return entity.SearchResult{}, log.Error(log.Fields{
			"search": search,
			"page":   page,
			"url":    url,
		}, err)
	}

	if json.Unmarshal(respData, &searchResultResponse) != nil {
		return entity.SearchResult{}, log.Error(log.Fields{
			"search": search,
			"page":   page,
			"url":    url,
		}, err)
	}

	var errMsg error
	if searchResultResponse.Response == "False" {
		errMsg = errors.New(searchResultResponse.Error)
	}

	for _, v := range searchResultResponse.Search {
		movieInfos = append(movieInfos, v)
	}

	return entity.SearchResult{
		Search:     search,
		Page:       page,
		MovieInfos: movieInfos,
		Error:      errMsg,
	}, nil
}
