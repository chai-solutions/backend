package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type Server struct {
	*http.Server
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Goodbye, cruel world!"))
}

func newAPI() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/hello", helloHandler)

	return r
}

func NewServer(addr string) (*Server, error) {
	mux := newAPI()

	s := http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{&s}, nil
}

func (s *Server) Start() {
	log.Info().Msg("starting...")

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Send()
		}
	}()

	log.Info().Msg("server is ready")
	s.WaitForShutdown()
}

func (s *Server) WaitForShutdown() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	sig := <-quit
	log.Info().Msgf("server shutdown signal received: %s", sig.String())

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	s.SetKeepAlivesEnabled(false)
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal().Msgf("unable to perform graceful shutdown: %s", err.Error())
	}
	log.Info().Msg("server has been shut down")
}
