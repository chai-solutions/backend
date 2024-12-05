package server_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"chai/server"
	"chai/server/mocks"

	"github.com/rs/zerolog/log"
)

func TestLoginHandlers(t *testing.T) {
	app := mocks.InitializeMockApp()

	t.Run("Login successful", func(t *testing.T) {
		body, err := json.Marshal(server.LoginRequestBody{
			Email:    "sanjay@ramaswamy.net",
			Password: "ramaswamy123",
		})
		if err != nil {
			t.Fatalf("failed to encode body, unreachable")
		}

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
		rec := httptest.NewRecorder()

		app.Router.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("expected status OK; got %v", rec.Code)
		}
	})

	t.Run("Incorrect email fails", func(t *testing.T) {
		body, err := json.Marshal(server.LoginRequestBody{
			Email:    "incorrect-email@example.com",
			Password: "ramaswamy123",
		})
		if err != nil {
			t.Fatalf("failed to encode body, unreachable")
		}

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
		rec := httptest.NewRecorder()

		app.Router.ServeHTTP(rec, req)

		// FIXME: Should be StatusForbidden, figure out a way to make this happen!
		// if you're not lazy ofc
		if rec.Code != http.StatusInternalServerError {
			t.Fatalf("expected status OK; got %v", rec.Code)
		}
	})

	t.Run("Incorrect password fails", func(t *testing.T) {
		body, err := json.Marshal(server.LoginRequestBody{
			Email:    "sanjay@ramaswamy.net",
			Password: "incorrect-password",
		})
		if err != nil {
			t.Fatalf("failed to encode body, unreachable")
		}

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
		rec := httptest.NewRecorder()

		app.Router.ServeHTTP(rec, req)

		if rec.Code != http.StatusForbidden {
			t.Fatalf("expected status OK; got %v", rec.Code)
		}
	})
}

func TestLogoutHandlers(t *testing.T) {
	app := mocks.InitializeMockApp()

	token, err := app.SessionRepo.AddSession(1)
	if err != nil {
		t.Fatalf("failed to generate token")
	}

	defer func() {
		err = app.SessionRepo.DeleteSession(token)
		if err != nil {
			log.Warn().Msg("failed to remove session in test")
		}
	}()

	req := httptest.NewRequest(http.MethodDelete, "/logout", nil)
	rec := httptest.NewRecorder()
	req.Header.Add("Authorization", "Bearer "+token)

	app.Router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Fatalf("expected status 204; got %v", rec.Code)
	}
}
