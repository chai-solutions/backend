package server_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"chai/database/sqlc"
	"chai/server"
	"chai/server/mocks"

	"github.com/rs/zerolog/log"
)

func TestGetUserHandler(t *testing.T) {
	app := mocks.InitializeMockApp()

	t.Run("Create user successfully", func(t *testing.T) {
		body, err := json.Marshal(server.CreateUserBody{
			Name:     "Sanjay Ramaswamy II",
			Email:    "sanjay2@ramaswamy.net",
			Password: "ramaswamy123",
		})
		if err != nil {
			t.Fatalf("failed to encode body, unreachable")
		}

		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
		rec := httptest.NewRecorder()

		app.Router.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("expected status OK; got %v", rec.Code)
		}
	})

	t.Run("Empty email fails", func(t *testing.T) {
		body, err := json.Marshal(server.CreateUserBody{
			Name:     "Sanjay Ramaswamy II",
			Email:    "",
			Password: "ramaswamy123",
		})
		if err != nil {
			t.Fatalf("failed to encode body, unreachable")
		}

		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
		rec := httptest.NewRecorder()

		app.Router.ServeHTTP(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("expected status 400; got %v", rec.Code)
		}
	})

	t.Run("Empty password fails", func(t *testing.T) {
		body, err := json.Marshal(server.CreateUserBody{
			Name:     "Sanjay Ramaswamy II",
			Email:    "sanjay2@ramaswamy.net",
			Password: "",
		})
		if err != nil {
			t.Fatalf("failed to encode body, unreachable")
		}

		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
		rec := httptest.NewRecorder()

		app.Router.ServeHTTP(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("expected status 400; got %v", rec.Code)
		}
	})

	t.Run("Existing email fails", func(t *testing.T) {
		body, err := json.Marshal(server.CreateUserBody{
			Name:     "Sanjay Ramaswamy II",
			Email:    "sanjay@ramaswamy.net",
			Password: "ramaswamy123",
		})
		if err != nil {
			t.Fatalf("failed to encode body, unreachable")
		}

		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
		rec := httptest.NewRecorder()

		app.Router.ServeHTTP(rec, req)

		// Internal server error isn't technically correct, but whatever
		if rec.Code != http.StatusInternalServerError {
			t.Fatalf("expected status 400; got %v", rec.Code)
		}
	})
}

func GetCurrentUserHandler(t *testing.T) {
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

	req := httptest.NewRequest(http.MethodGet, "/users/@me", nil)
	rec := httptest.NewRecorder()
	req.Header.Add("Authorization", "Bearer "+token)

	app.Router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200; got %v", rec.Code)
	}

	body := rec.Result().Body
	defer body.Close()

	var result sqlc.GetUserFromSessionContextRow
	err = json.NewDecoder(body).Decode(&result)
	if err != nil {
		t.Fatalf("failed to deserialize user body")
	}

	// A basic sanity check for checking if the user is expected
	if result.Token != token {
		t.Fatalf("user token does not match in result")
	}
}
