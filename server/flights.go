package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *App) FlightsHandler(w http.ResponseWriter, r *http.Request) {
	depAirport := r.URL.Query().Get("departure_airport")
	arrAirport := r.URL.Query().Get("arrival_airport")
	if depAirport == "" {
		http.Error(w, "Missing Departure Airport", http.StatusBadRequest)
		return
	}

	if arrAirport == "" {
		http.Error(w, "Missing Arrival Airport", http.StatusBadRequest)
		return
	}

	flights, err := a.FlightsRepo.FlightsByDepartureArrival(depAirport, arrAirport)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(flights)
	_ = err
}

func (a *App) FlightHandler(w http.ResponseWriter, r *http.Request) {
	flightCode := chi.URLParam(r, "flightNumber")
	if flightCode == "" {
		http.Error(w, "Invalid Flight Number", http.StatusBadRequest)
		return
	}

	flight, err := a.FlightsRepo.FlightByCode(flightCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(flight)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
