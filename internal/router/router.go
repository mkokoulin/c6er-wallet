package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/mkokoulin/c6er-wallet.git/internal/config"
	"github.com/mkokoulin/c6er-wallet.git/internal/handlers"
)

func New(h *handlers.Handlers, cfg *config.Config) *chi.Mux {
	router := chi.NewRouter()


	router.Route("/", func(r chi.Router) {
		r.Post("/api/v1/user/registration", h.Registration)
		r.Post("/api/v1/user/login", h.Login)
	})

	return router
}