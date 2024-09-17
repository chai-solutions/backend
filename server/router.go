package server

import (
	"chai/middleware"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func (a *App) RegisterRoutes() {
	r := a.Router

	r.Use(chiMiddleware.Recoverer)
	r.Use(middleware.RequestLogger)
	r.Use(middleware.JSONContentType)

	r.Get("/hello", a.HelloHandler)

	r.Route("/{account}", func(r chi.Router) {
		r.Post("/", a.AccountHandler)
		r.Get("/", a.AccountHandler)
	})
}
