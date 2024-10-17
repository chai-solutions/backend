package server

import (
	"context"
	"encoding/json"
	"net/http"
)

func (a *App) AirportsHandler(w http.ResponseWriter, _ *http.Request) {
	airports, err := a.Queries.GetAllAirports(context.Background())
	if err != nil {
		http.Error(w, "Failed to get airports", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(airports); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
