package server

import (
	"chai/middleware"

	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func (a *App) RegisterRoutes() {
	r := a.Router

	r.Use(chiMiddleware.Recoverer)
	r.Use(middleware.RequestLogger)
	r.Use(middleware.JSONContentType)

	r.Get("/hello", a.HelloHandler)
	r.Post("/users", a.CreateUserHandler)

	r.Post("/login", a.LoginHandler)
	// TODO: add user auth middleware
}
