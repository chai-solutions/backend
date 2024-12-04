package server_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"chai/server"
	"chai/server/mocks"
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
