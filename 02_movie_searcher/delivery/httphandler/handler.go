package httphandler

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/pikomonde/i-view-stockbit/02_movie_searcher/helper/log"
	servMovs "github.com/pikomonde/i-view-stockbit/02_movie_searcher/service/moviesearch"
)

// Handler is used to contains handler delivery
type Handler struct {
	ServiceMovieSearch *servMovs.ServiceMovieSearch
	Mux                *http.ServeMux
}

// Opt contains option for httphandler
type Opt struct {
	ServiceMovieSearch *servMovs.ServiceMovieSearch
}

// New returns new http handler
func New(opt Opt) *Handler {
	return &Handler{
		ServiceMovieSearch: opt.ServiceMovieSearch,
		Mux:                http.NewServeMux(),
	}
}

// Start http server
func (h *Handler) Start() {
	h.RegisterMovieSearch()

	// Starting server
	port := ":5001" // TODO: should be from config
	srv := &http.Server{
		ReadTimeout:  5000 * time.Millisecond, // TODO: should be from config
		WriteTimeout: 5000 * time.Millisecond, // TODO: should be from config
		TLSConfig: &tls.Config{
			PreferServerCipherSuites: true,
			CurvePreferences: []tls.CurveID{
				tls.CurveP256,
				tls.X25519,
			},
		},
		Handler: h.Mux,
		Addr:    port,
	}

	go func(srv *http.Server) {
		err := srv.ListenAndServe()
		if err != nil {
			log.Error(log.Fields{}, err)
		}
	}(srv)
	fmt.Println("Server listen to ", port)
}
