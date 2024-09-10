package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

type App struct {
	db *pgxpool.Pool
	s  *http.Server
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Goodbye, cruel world!"))
}

func newAPI() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/hello", helloHandler)

	return r
}

func newDBPool() *pgxpool.Pool {
	user := os.Getenv("DB_USER")
	pw := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_NAME")

	db_url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, pw, host, port, db_name)

	pool, err := pgxpool.New(context.Background(), db_url)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	return pool
}

func NewApp(addr string) *App {
	mux := newAPI()
	db := newDBPool()

	s := http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &App{
		db: db,
		s:  &s,
	}
}

func (a *App) Start() {
	log.Info().Msg("starting...")

	go func() {
		if err := a.s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Send()
		}
	}()

	log.Info().Msg("server is ready")
	a.WaitForShutdown()
}

func (a *App) WaitForShutdown() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	sig := <-quit
	log.Info().Msgf("server shutdown signal received: %s", sig.String())

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	a.s.SetKeepAlivesEnabled(false)
	if err := a.s.Shutdown(ctx); err != nil {
		log.Fatal().Msgf("unable to perform graceful shutdown of chi mux: %s", err.Error())
	}
	log.Info().Msg("closing database connections")
	a.db.Close()
	log.Info().Msg("server has been shut down")
}
