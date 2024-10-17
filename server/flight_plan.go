package server

import (
	"chai/middleware"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"chai/database/sqlc"

	"github.com/go-chi/chi/v5"
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

func (a *App) GetFlightPlansHandler(w http.ResponseWriter, r *http.Request) {
	user := middleware.MustGetUserFromContext(r.Context())

	flightPlans, err := a.Queries.GetFlightPlans(context.Background(), user.ID)
	if err != nil {
		log.Error().AnErr("GetFlightPlans", err).Msg("Failed to get flight plans")
		http.Error(w, "Failed to get flight plans", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(flightPlans); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (a *App) GetFlightPlanHandler(w http.ResponseWriter, r *http.Request) {
	user := middleware.MustGetUserFromContext(r.Context())
	idInt, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid Flight ID", http.StatusBadRequest)
		return
	}
	id := int32(idInt)

	flightPlans, err := a.Queries.GetFlightPlan(context.Background(), sqlc.GetFlightPlanParams{
		Users: user.ID,
		ID:    id,
	})
	if err != nil {
		log.Error().AnErr("GetFlightPlans", err).Msg("Failed to get flight plans")
		http.Error(w, "Failed to get flight plans", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(flightPlans); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
