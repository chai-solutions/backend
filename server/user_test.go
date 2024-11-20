package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUserHandler(t *testing.T) {
	app := NewTestApp()
	app.Router.Post("/", app.CreateUserHandler)

	tests := []struct {
		name           string
		requestBody    map[string]string
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "valid request",
			requestBody: map[string]string{
				"name":     "Sanjay Ramaswamy",
				"email":    "sanjay@ramaswamy.net",
				"password": "ramaswamy123",
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"name":"Sanjay Ramaswamy","email":"sanjay@ramaswamy.net".*}`,
		},
		{
			name: "missing name",
			requestBody: map[string]string{
				"email":    "sanjay@ramaswamy.net",
				"password": "ramaswamy123",
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "missing name\n",
		},
		{
			name: "missing email",
			requestBody: map[string]string{
				"name":     "Sanjay Ramaswamy",
				"password": "ramaswamy123",
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "missing email\n",
		},
		{
			name: "short password",
			requestBody: map[string]string{
				"name":     "Sanjay Ramaswamy",
				"email":    "sanjay@ramaswamy.net",
				"password": "short",
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "password length must be at least 8 characters\n",
		},
		{
			name:           "malformed JSON",
			requestBody:    nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "malformed JSON\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var body []byte
			if tt.requestBody != nil {
				body, _ = json.Marshal(tt.requestBody)
			}
			req, err := http.NewRequest("POST", "/", bytes.NewReader(body))
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}
			rr := httptest.NewRecorder()
			app.Router.ServeHTTP(rr, req)
			// Validate response status
			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, status)
			}

			// Validate response body using regex matching
			if err := validateResponseBody(rr.Body.String(), tt.expectedBody); err != nil {
				t.Errorf("response body mismatch: %v", err)
			}

			// Cleanup: Delete user after test
			cleanupUser(t, app, tt.requestBody["email"])
		})
	}
}
