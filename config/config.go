package config

import (
	"os"
	"strconv"

	"github.com/rs/zerolog/log"
)

type AppConfig struct {
	Env         string
	Port        int
	DatabaseURL string
}

func GetConfig() AppConfig {
	env, found := os.LookupEnv("APP_ENV")
	if !found {
		log.Fatal().Msg("APP_ENV is not defined")
	}
	if env != "prod" && env != "dev" {
		log.Fatal().Msg("APP_ENV must be either 'dev' or 'prod'")
	}

	portVar, found := os.LookupEnv("APP_PORT")
	if !found {
		log.Fatal().Msg("APP_PORT is not defined")
	}
	port, err := strconv.Atoi(portVar)
	if err != nil {
		log.Fatal().Msg("APP_PORT must be an integer")
	}

	dbURL, found := os.LookupEnv("DATABASE_URL")
	if !found {
		log.Fatal().Msg("DATABASE_URL is not defined")
	}

	return AppConfig{
		Env:         env,
		Port:        port,
		DatabaseURL: dbURL,
	}
}
