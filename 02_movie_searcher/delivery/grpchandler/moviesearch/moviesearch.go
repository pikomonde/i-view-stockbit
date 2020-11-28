package moviesearch

import (
	context "context"

	servMovs "github.com/pikomonde/i-view-stockbit/02_movie_searcher/service/moviesearch"
)

// Server is used as a grpc server for movisearch
type Server struct {
	ServiceMovieSearch *servMovs.ServiceMovieSearch
}

// Opt contains moviesearch server option
type Opt struct {
	ServiceMovieSearch *servMovs.ServiceMovieSearch
}

// New returns moviesearch grpc server
func New(opt Opt) *Server {
	return &Server{
		ServiceMovieSearch: opt.ServiceMovieSearch,
	}
}

// GetAndSaveMoviesBySearch getting search data from OMDB and then save it to DB
func (s *Server) GetAndSaveMoviesBySearch(ctx context.Context, in *MessageRequest) (*MessageResponse, error) {
	err := s.ServiceMovieSearch.GetAndSaveMoviesBySearchWord(in.Searchword, in.Pagination)
	if err != nil {
		return &MessageResponse{
			Response: err.Error(),
		}, nil
	}

	return &MessageResponse{
		Response: "success",
	}, nil
}
