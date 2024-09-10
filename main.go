package main

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

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

	log.Info().Str("env", app_env).Send()

	app_port := os.Getenv("APP_PORT")

	server := NewApp(":" + app_port)
	server.Start()
}
