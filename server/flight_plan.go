package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"chai/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type flightPlanBody struct {
	FlightNumber string `json:"flightNumber"`
}

func (a *App) PatchFlightPlanHandler(w http.ResponseWriter, r *http.Request) {
	var body flightPlanBody

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "malformed JSON", http.StatusBadRequest)
		return
	}

	planID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid Flight ID", http.StatusBadRequest)
		return
	}

	flightPlan, err := a.FlightPlanRepo.AddFlightToPlan(int32(planID), body.FlightNumber)
	if err != nil {
		http.Error(w, "failed to insert flight plan flight", http.StatusInternalServerError)
		log.Error().Err(err).Msg("failed to patch flight plan")
		return
	}

	if err := json.NewEncoder(w).Encode(flightPlan); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}

func (a *App) CreateFlightPlanHandler(w http.ResponseWriter, r *http.Request) {
	var body flightPlanBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "malformed JSON", http.StatusBadRequest)
		return
	}
	user := middleware.MustGetUserFromContext(r.Context())

	flightPlan, err := a.FlightPlanRepo.CreatePlan(user.ID, body.FlightNumber)
	if err != nil {
		log.Error().AnErr("CreateFlightPlan", err).Msg("Failed to create flight plan")
		http.Error(w, "Failed to create flight plan", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(flightPlan); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (a *App) GetFlightPlansHandler(w http.ResponseWriter, r *http.Request) {
	user := middleware.MustGetUserFromContext(r.Context())

	flightPlans, err := a.FlightPlanRepo.GetPlansForUser(user.ID)
	if err != nil {
		log.Error().AnErr("GetFlightPlan", err).Msg("Failed to get flight plans")
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

	flightPlans, err := a.FlightPlanRepo.GetPlan(user.ID, flightPlanID)
	if err != nil {
		log.Error().AnErr("GetFlightPlan", err).Msg("Failed to get flight plans")
		http.Error(w, "Failed to get flight plans", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(flightPlans); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (a *App) DeleteFlightPlan(w http.ResponseWriter, r *http.Request) {
	planID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid Flight Plan ID", http.StatusBadRequest)
		return
	}

	err = a.FlightPlanRepo.DeletePlan(int32(planID))
	if err != nil {
		log.Error().AnErr("DeleteFlightPlan", err).Msg("Failed to delete flight plan")
		http.Error(w, "Failed to delete flight plan", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (a *App) DeleteFlightPlanStep(w http.ResponseWriter, r *http.Request) {
	stepID, err := strconv.Atoi(chi.URLParam(r, "stepID"))
	if err != nil {
		http.Error(w, "Invalid Flight Plan Step ID", http.StatusBadRequest)
		return
	}

	err = a.FlightPlanRepo.DeleteFlightFromPlan(int32(stepID))
	log.Error().AnErr("DeleteFlightPlanStep", err).Msg("Failed to delete flight plan step")
	if err != nil {
		http.Error(w, "Failed to delete flight plan step", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
