package main

import (
	"os"
	"time"

	"chai/config"
	"chai/database"
	"chai/repos"
	"chai/server"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Warn().Err(err).Send()
	}

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

	_, queries := database.NewPool(cfg.DatabaseURL)

	userRepo := repos.NewUserRepository(queries)

	database.RunMigrations(cfg.DatabaseURL)
	server := server.NewApp(cfg, userRepo)
	server.RegisterRoutes()

	server.Start()
}
