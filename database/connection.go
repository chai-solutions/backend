package database

import (
	"chai/database/sqlc"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

func NewPool(url string) (*pgxpool.Pool, *sqlc.Queries) {
	pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	queries := sqlc.New(pool)

	return pool, queries
}
