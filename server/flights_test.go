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

func TestGetFlightsHandlers(t *testing.T) {
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

	t.Run("Fetch flights by ID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/flights/DL456", nil)
		rec := httptest.NewRecorder()
		req.Header.Add("Authorization", "Bearer "+token)

		app.Router.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("expected status 200; got %v", rec.Code)
		}

		body := rec.Result().Body
		defer body.Close()

		var flight sqlc.GetFlightRow
		err = json.NewDecoder(body).Decode(&flight)
		if err != nil {
			t.Fatalf("failed to deserialize flight body")
		}

		if flight.FlightNumber != "DL456" {
			t.Fatalf("did not receive correct flight number %s from response, expected %s", "DL456", flight.FlightNumber)
		}
	})

	t.Run("Fetch flights by departure/arrival airports", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/flights?departure_airport=OAK&arrival_airport=SJC", nil)
		rec := httptest.NewRecorder()
		req.Header.Add("Authorization", "Bearer "+token)

		app.Router.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("expected status 200; got %v", rec.Code)
		}

		body := rec.Result().Body
		defer body.Close()

		var flights []sqlc.GetFlightsRow
		err = json.NewDecoder(body).Decode(&flights)
		if err != nil {
			t.Fatalf("failed to deserialize flight body")
		}

		if len(flights) != 1 {
			t.Fatalf("expected 1 result in flights array, got %v", len(flights))
		}

		flight := flights[0]
		if flight.FlightNumber != "DL456" {
			t.Fatalf("did not receive correct flight number %s from response, expected %s", "DL456", flight.FlightNumber)
		}
	})
}
