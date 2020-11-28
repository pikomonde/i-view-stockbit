package grpchandler

import (
	"fmt"
	"net"

	"github.com/pikomonde/i-view-stockbit/02_movie_searcher/delivery/grpchandler/moviesearch"
	"github.com/pikomonde/i-view-stockbit/02_movie_searcher/helper/log"
	servMovs "github.com/pikomonde/i-view-stockbit/02_movie_searcher/service/moviesearch"
	"google.golang.org/grpc"
)

// Handler is used to contains handler delivery
type Handler struct {
	ServiceMovieSearch *servMovs.ServiceMovieSearch
}

// Opt contains option for httphandler
type Opt struct {
	ServiceMovieSearch *servMovs.ServiceMovieSearch
}

// New returns new http handler
func New(opt Opt) *Handler {
	return &Handler{
		ServiceMovieSearch: opt.ServiceMovieSearch,
	}
}

// Start http server
func (h *Handler) Start() {
	port := ":9044"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Error(log.Fields{}, err)
		return
	}

	grpcServer := grpc.NewServer()
	sMovs := moviesearch.New(moviesearch.Opt{
		ServiceMovieSearch: h.ServiceMovieSearch,
	})
	moviesearch.RegisterMovieSearchServiceServer(grpcServer, sMovs)

	go func(lis net.Listener) {
		err = grpcServer.Serve(lis)
		if err != nil {
			log.Error(log.Fields{}, err)
		}
	}(lis)
	fmt.Println("GRPC Listen to ", port)
}
