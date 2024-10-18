package server

import (
	"chai/database/sqlc"
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *App) FlightsHandler(w http.ResponseWriter, r *http.Request) {
	depAirport := r.URL.Query().Get("departure_airport")
	arrAirport := r.URL.Query().Get("arrival_airport")
	if depAirport == "" || arrAirport == "" {
		http.Error(w, "Missing Arrival or Departure Airport", http.StatusBadRequest)
	}

	var params sqlc.GetFlightsParams
	params.Arr = arrAirport
	params.Dep = depAirport

	flights, err := a.Queries.GetFlights(context.Background(), params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(
		flights)
	_ = err
}

func (a *App) FlightHandler(w http.ResponseWriter, r *http.Request) {
	flight_number := chi.URLParam(r, "flight_number")
	if flight_number == "" {
		http.Error(w, "Invalid Flight Number", http.StatusBadRequest)
		return
	}

	flight, err := a.Queries.GetFlight(context.Background(), flight_number)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(
		flight)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
