package database

import (
	"chai/database/sqlc"
	"context"
	"embed"
	"net/url"

	"github.com/amacneil/dbmate/v2/pkg/dbmate"
	// Import the postgres driver for dbmate
	_ "github.com/amacneil/dbmate/v2/pkg/driver/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

//go:embed migrations/*.sql
var fs embed.FS

func NewPool(url string) (*pgxpool.Pool, *sqlc.Queries) {
	pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	queries := sqlc.New(pool)

	return pool, queries
}

func RunMigrations(databaseURL string) {
	u, err := url.Parse(databaseURL)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	db := dbmate.New(u)
	db.FS = fs // Set the embedded filesystem for migrations
	db.MigrationsDir = []string{"migrations"}
	db.AutoDumpSchema = false

	err = db.CreateAndMigrate()
	if err != nil {
		log.Fatal().Err(err).Send()
	}
}
