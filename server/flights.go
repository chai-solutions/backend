package server

import (
	"chai/database/sqlc"
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type FlightRes struct {
	Message string      `json:"message"`
	Data    interface{} `json:"accounts"`
}

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
		FlightRes{
			Message: "Flights retrieved successfully",
			Data:    flights,
		})
	_ = err
}

func (a *App) FlightHandler(w http.ResponseWriter, r *http.Request) {
	idInt, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid Flight ID", http.StatusBadRequest)
	}

	flight, err := a.Queries.GetFlight(context.Background(), int32(idInt))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(
		FlightRes{
			Message: "Flight retrieved successfully",
			Data:    flight,
		})
	_ = err
}