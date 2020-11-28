package httphandler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	servMovs "github.com/pikomonde/i-view-stockbit/02_movie_searcher/service/moviesearch"
)

// MovieSearchHandler contains handler for Movie Search
type MovieSearchHandler struct {
	ServiceMovieSearch *servMovs.ServiceMovieSearch
	Mux                *http.ServeMux
}

// RegisterMovieSearch is used to register endpoints for movie search
func (h *Handler) RegisterMovieSearch() {
	hh := MovieSearchHandler{
		ServiceMovieSearch: h.ServiceMovieSearch,
		Mux:                h.Mux,
	}

	hh.Mux.HandleFunc("/moviesearch", hh.GetAndSaveMoviesBySearch)
}

// GetAndSaveMoviesBySearch get movies by search word and then save it to db
func (hh *MovieSearchHandler) GetAndSaveMoviesBySearch(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Search string `json:"searchword"`
		Page   string `json:"pagination"`
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErrorJSON(w, r, http.StatusBadRequest, errBadRequest)
		return
	}

	err = json.Unmarshal(body, &input)
	if err != nil {
		respErrorJSON(w, r, http.StatusBadRequest, errBadRequest)
		return
	}

	pageInt64, err := strconv.ParseInt(input.Page, 10, 64)
	if err != nil {
		respErrorJSON(w, r, http.StatusBadRequest, errBadRequest)
		return
	}

	err = hh.ServiceMovieSearch.GetAndSaveMoviesBySearchWord(input.Search, pageInt64)
	if err != nil {
		respSuccessJSON(w, r, err.Error())
		return
	}
	respSuccessJSON(w, r, "success")
}
