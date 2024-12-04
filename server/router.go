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
	r.Post("/users", a.CreateUserHandler)

	// r.Post("/login", a.LoginHandler)

	r.Group(func(r chi.Router) {
		// r.Use(middleware.APIAuthorization(a.Queries))

		// r.Get("/users/@me", a.UserInfoHandler)
		// r.Delete("/logout", a.LogoutHandler)

		// r.Get("/airports", a.AirportsHandler)
		//
		// r.Get("/flights", a.FlightsHandler)
		// r.Get("/flights/{flightNumber}", a.FlightHandler)

		// r.Post("/flight_plans", a.CreateFlightPlanHandler)
		// r.Post("/flight_plans/{id}", a.PatchFlightPlanHandler)
		// r.Get("/flight_plans", a.GetFlightPlansHandler)
		// r.Get("/flight_plans/{id}", a.GetFlightPlanHandler)
		// r.Delete("/flight_plans/{id}/{stepID}", a.DeleteFlightPlanStep)
		// r.Delete("/flight_plans/{id}", a.DeleteFlightPlan)
		//
		// r.Post("/webhook", a.WebhookHandler)
	})
}
