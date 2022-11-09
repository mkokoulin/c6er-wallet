package server

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	addr string
	handler *chi.Mux
	s *http.Server
}

func New(addr string, handler *chi.Mux) *Server {
	srv := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	return &Server {
		addr: addr,
		handler: handler,
		s: srv,
	}
}

func (s *Server) Start(ctx context.Context) (func(ctx context.Context) error, error) {
	err := s.s.ListenAndServe()
	if err != nil {
		return nil, err
	}

	return s.s.Shutdown, nil
}