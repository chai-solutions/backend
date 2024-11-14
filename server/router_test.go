package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func newTestApp() *App {
	app := &App{
		Router: chi.NewRouter(),
	}
	return app
}

func TestRegisterRoutes(t *testing.T) {
	app := newTestApp()
	app.RegisterRoutes()

	tests := []struct {
		method     string
		url        string
		body       interface{} // Request body (if any)
		statusCode int         // Expected HTTP status code
	}{
		{"GET", "/hello", nil, http.StatusOK},
		{"POST", "/users", map[string]string{"username": "test", "password": "password123"}, http.StatusBadRequest},
		{"POST", "/login", map[string]string{"username": "test", "password": "password123"}, http.StatusBadRequest},
		{"GET", "/users/@me", nil, http.StatusForbidden},
		//{"GET", "/flights", nil, http.StatusOK},
		//{"GET", "/flights/123", nil, http.StatusOK},
		//{"POST", "/flight_plans", nil, http.StatusUnauthorized},
	}

	for _, tt := range tests {
		var requestBody []byte
		if tt.body != nil {
			var err error
			requestBody, err = json.Marshal(tt.body)
			if err != nil {
				t.Fatalf("could not marshal request body: %v", err)
			}
		}

		req, err := http.NewRequest(tt.method, tt.url, bytes.NewReader(requestBody))
		if err != nil {
			t.Fatalf("could not create request for %s %s: %v", tt.method, tt.url, err)
		}
		if tt.method == "POST" {
			req.Header.Set("Content-Type", "application/json")
		}

		rr := httptest.NewRecorder()

		app.Router.ServeHTTP(rr, req)

		if status := rr.Code; status != tt.statusCode {
			t.Errorf("method %s %s: expected status %d, got %d", tt.method, tt.url, tt.statusCode, status)
		}
	}
}
