package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, World!")
}

func main() {
	app_env := os.Getenv("APP_ENV")

	if app_env == "prod" {
		// Do nothing, use defaults
	} else if app_env == "dev" {
		// I like pretty output
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
			Level(zerolog.TraceLevel).
			With().
			Timestamp().
			Caller().Logger()
	} else {
		log.Fatal().Msg("APP_ENV is not set")
	}

	log.Info().Msg("starting...")
	log.Info().Str("env", app_env).Send()

	r := chi.NewRouter()
	r.Get("/hello", helloHandler)

	log.Fatal().Err(http.ListenAndServe(":8080", r))
}
