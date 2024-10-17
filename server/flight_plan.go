package server

import (
	"chai/middleware"
	"context"
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

func (a *App) CreateFlightPlanHandler(w http.ResponseWriter, r *http.Request) {
	user := middleware.MustGetUserFromContext(r.Context())

	flightPlan, err := a.Queries.CreateFlightPlan(context.Background(), user.ID)
	if err != nil {
		log.Error().AnErr("CreateFlightPlan", err).Msg("Failed to create flight plan")
		http.Error(w, "Failed to create flight plan", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(flightPlan); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
