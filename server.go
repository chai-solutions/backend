package main

import (
	"context"
	"encoding/json"
	"math/rand/v2"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"chai/middleware"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

type App struct {
	db *pgxpool.Pool
	s  *http.Server
}

type helloRes struct {
	Message string `json:"message"`
	Number  int    `json:"number"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	number := rand.IntN(69_420_000)
	_ = number

	_ = json.NewEncoder(w).Encode(helloRes{
		Message: "Hello, world! The server is alive.",
		Number:  number,
	})
}

func newAPI() *chi.Mux {
	r := chi.NewRouter()

	r.Use(chiMiddleware.Recoverer)
	r.Use(middleware.RequestLogger)
	r.Use(middleware.JSONContentType)

	r.Get("/hello", helloHandler)

	return r
}

func newDBPool() *pgxpool.Pool {
	db_url, found := os.LookupEnv("DATABASE_URL")
	if !found {
		log.Fatal().Msg("DATABASE_URL is not defined")
	}

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

	log.Info().Msgf("server is listening on %s", a.s.Addr)
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
