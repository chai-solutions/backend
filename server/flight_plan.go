package server

import (
	"chai/database/sqlc"
	"chai/middleware"
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type flightPlanBody struct {
	FlightNumber string `json:"flight_number"`
}

func (a *App) PatchFlightPlanHandler(w http.ResponseWriter, r *http.Request) {
	_ = middleware.MustGetUserFromContext(r.Context())
	var params sqlc.PatchFlightPlanParams
	var body flightPlanBody

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "malformed JSON", http.StatusBadRequest)
		return
	}
	planId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid Flight ID", http.StatusBadRequest)
		return
	}
	params.FlightNumber = body.FlightNumber
	params.FlightPlan = int32(planId)

	flight_plan, err := a.Queries.PatchFlightPlan(context.Background(), params)
	if err != nil {
		http.Error(w, "failed to insert flight plan flight", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(flight_plan); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}

func (a *App) CreateFlightPlanHandler(w http.ResponseWriter, r *http.Request) {
	var body flightPlanBody
	var params sqlc.CreateFlightPlanParams
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "malformed JSON", http.StatusBadRequest)
		return
	}
	user := middleware.MustGetUserFromContext(r.Context())
	params.Users = user.ID
	params.FlightNumber = body.FlightNumber

	flight_plan, err := a.Queries.CreateFlightPlan(context.Background(), params)
	if err != nil {
		log.Error().AnErr("CreateFlightPlan", err).Msg("Failed to create flight plan")
		http.Error(w, "Failed to create flight plan", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(flight_plan); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
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
	param, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid Flight ID", http.StatusBadRequest)
		return
	}
	flightPlanID := int32(param)
	params := sqlc.GetFlightPlanParams{
		Users: user.ID,
		ID:    flightPlanID,
	}

	flightPlans, err := a.Queries.GetFlightPlan(context.Background(), params)
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
