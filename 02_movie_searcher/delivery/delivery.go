package delivery

import (
	"github.com/pikomonde/i-view-stockbit/02_movie_searcher/delivery/grpchandler"
	"github.com/pikomonde/i-view-stockbit/02_movie_searcher/delivery/httphandler"
	servMovs "github.com/pikomonde/i-view-stockbit/02_movie_searcher/service/moviesearch"
)

// Delivery contains services and endpoints
type Delivery struct {
	HTTPHandler *httphandler.Handler
	GRPCHandler *grpchandler.Handler
}

// Opt is delivery options
type Opt struct {
	ServiceMovieSearch *servMovs.ServiceMovieSearch
}

// New returns the delivery
func New(opt Opt) *Delivery {
	return &Delivery{
		HTTPHandler: httphandler.New(httphandler.Opt{
			ServiceMovieSearch: opt.ServiceMovieSearch,
		}),
		GRPCHandler: grpchandler.New(grpchandler.Opt{
			ServiceMovieSearch: opt.ServiceMovieSearch,
		}),
	}
}

// Start starts the delivery server
func (d *Delivery) Start() {
	d.HTTPHandler.Start()
	d.GRPCHandler.Start()
}
