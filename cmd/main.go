package main

import (
	"os"
	"time"

	"chai/config"
	"chai/database"
	"chai/server"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg := config.GetConfig()

	if cfg.Env == "dev" {
		// I like pretty output
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
			Level(zerolog.TraceLevel).
			With().
			Timestamp().
			Caller().
			Logger()
	}

	log.Info().Str("env", cfg.Env).Send()

	db := database.NewPool(cfg.DatabaseURL)

	server := server.NewApp(cfg, db)
	server.RegisterRoutes()

	server.Start()
}
