package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"chai/config"
	"chai/repos"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type App struct {
	Config config.AppConfig
	Router *chi.Mux
	Server *http.Server

	UserRepo     repos.UserRepository
	SessionRepo  repos.SessionRepository
	AirportsRepo repos.AirportsRepository
}

type Repositories struct {
	UserRepo     repos.UserRepository
	SessionRepo  repos.SessionRepository
	AirportsRepo repos.AirportsRepository
}

func NewApp(cfg config.AppConfig, repos Repositories) *App {
	mux := chi.NewMux()

	addr := fmt.Sprintf(":%d", cfg.Port)

	s := http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &App{
		Config: cfg,
		Router: mux,
		Server: &s,

		UserRepo:     repos.UserRepo,
		SessionRepo:  repos.SessionRepo,
		AirportsRepo: repos.AirportsRepo,
	}
}

func (a *App) Start() {
	log.Info().Msg("starting...")

	go func() {
		if err := a.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Send()
		}
	}()

	log.Info().Msgf("server is listening on %s", a.Server.Addr)
	a.WaitForShutdown()
}

func (a *App) WaitForShutdown() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	sig := <-quit
	log.Info().Msgf("server shutdown signal received: %s", sig.String())

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	a.Server.SetKeepAlivesEnabled(false)
	if err := a.Server.Shutdown(ctx); err != nil {
		log.Fatal().Msgf("unable to perform graceful shutdown of chi mux: %s", err.Error())
	}
	log.Info().Msg("closing database connections")
	a.Server.Close()
	log.Info().Msg("server has been shut down")
}
