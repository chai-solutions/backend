package server_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"chai/database/sqlc"
	"chai/server/mocks"

	"github.com/rs/zerolog/log"
)

func TestAirportsHandler(t *testing.T) {
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

	req := httptest.NewRequest(http.MethodGet, "/airports", nil)
	rec := httptest.NewRecorder()
	req.Header.Add("Authorization", "Bearer "+token)

	app.Router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200; got %v", rec.Code)
	}

	body := rec.Result().Body
	defer body.Close()

	var airports []sqlc.Airport
	err = json.NewDecoder(body).Decode(&airports)
	if err != nil {
		t.Fatalf("failed to deserialize aiports body")
	}

	// A basic sanity check for checking if the number of airports
	// is what exists
	if len(airports) != 3 {
		t.Fatalf("expected 3 airports, got %v", len(airports))
	}
}
