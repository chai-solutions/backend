package server

import (
	"chai/database/sqlc"
	"context"
	"encoding/json"
	"net/http"
)

type FlightRes struct {
	Message string               `json:"message"`
	Data    []sqlc.GetFlightsRow `json:"accounts"`
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
