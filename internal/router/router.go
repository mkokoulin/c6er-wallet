package router

import (
	"github.com/go-chi/chi/v5"

	"github.com/mkokoulin/c6er-wallet.git/internal/config"
	"github.com/mkokoulin/c6er-wallet.git/internal/handlers"
	"github.com/mkokoulin/c6er-wallet.git/internal/middlewares"
)

func New(h *handlers.Handlers, cfg *config.Config) *chi.Mux {
	router := chi.NewRouter()


	router.Route("/api/v1", func(r chi.Router) {
		r.Use(middlewares.JWTMiddleware(cfg))
		r.Post("/user/signup", h.Registration)
		r.Post("/user/login", h.Login)
		r.Get("/user/auth", h.CheckAuth)
	})

	return router
}