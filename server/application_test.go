package server

import (
	"chai/config"
	"chai/database/sqlc"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func TestNewApp(t *testing.T) {
	cfg := config.AppConfig{Port: 8080}
	db := &pgxpool.Pool{}
	queries := &sqlc.Queries{}

	app := NewApp(cfg, db, queries)

	if app == nil {
		t.Errorf("app is nil 😳")
	}
	if cfg != app.Config {
		t.Errorf("😿")
	}
	if db != app.DB {
		t.Errorf("😭")
	}
	if queries != app.Queries {
		t.Errorf("😭")
	}
	if app.Router == nil {
		t.Errorf("😳")
	}
	if app.Server == nil {
		t.Errorf("😭")
	}
	if app.Server.Addr != ":8080" {
		t.Errorf("😢")
	}
}

func TestStart(t *testing.T) {
	cfg := config.AppConfig{Port: 8080}
	db := &pgxpool.Pool{}
	queries := &sqlc.Queries{}

	app := NewApp(cfg, db, queries)

	go func() {
		time.Sleep(1 * time.Second)
		if err := syscall.Kill(syscall.Getpid(), syscall.SIGINT); err != nil {
			t.Errorf("failed to send SIGINT: %v", err)
		}
	}()

	app.Start()

	if app.Server.Addr != ":8080" {
		t.Errorf("server port bad")
	}
}

func TestWaitForShutdown(t *testing.T) {
	cfg := config.AppConfig{Port: 8080}
	db := &pgxpool.Pool{}
	queries := &sqlc.Queries{}

	app := NewApp(cfg, db, queries)

	go func() {
		time.Sleep(1 * time.Second)
		if err := syscall.Kill(syscall.Getpid(), syscall.SIGINT); err != nil {
			t.Errorf("failed to send SIGINT: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		sig := <-quit
		if os.Interrupt != sig {
			t.Errorf("Wrong interupt")
		}
	}()

	app.WaitForShutdown()
}
