package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"chai/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
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
		if err == pgx.ErrNoRows {
			http.Error(w, "flight number or flight plan not found", http.StatusNotFound)
			return
		}

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
	if err == pgx.ErrNoRows {
		http.Error(w, fmt.Sprintf("no flight with number %s found", body.FlightNumber), http.StatusNotFound)
		return
	}
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
		log.Error().AnErr("GetFlightPlans", err).Msg("Failed to get flight plans")
		http.Error(w, "Failed to get flight plans", http.StatusInternalServerError)
		return
	}

	if len(flightPlans) == 0 {
		_, _ = w.Write([]byte("[]"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(flightPlans); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (a *App) GetFlightPlanHandler(w http.ResponseWriter, r *http.Request) {
	param, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid Flight ID", http.StatusBadRequest)
		return
	}
	flightPlanID := int32(param)

	exists, err := a.FlightPlanRepo.Exists(flightPlanID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to check for existence of flight plan")
		http.Error(w, "Failed to get flight plans", http.StatusInternalServerError)
		return
	}
	if !exists {
		http.Error(w, fmt.Sprintf("no flight plan with id %d exists", flightPlanID), http.StatusInternalServerError)
		return
	}

	flightPlanSteps, err := a.FlightPlanRepo.GetPlan(flightPlanID)
	if err != nil {
		log.Error().AnErr("GetFlightPlan", err).Msg("Failed to get flight plans")
		http.Error(w, "Failed to get flight plans", http.StatusInternalServerError)
		return
	}

	if len(flightPlanSteps) == 0 {
		http.Error(w, "flight plan is empty", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(flightPlanSteps); err != nil {
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
	planID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid Flight Plan ID", http.StatusBadRequest)
		return
	}
	stepID, err := strconv.Atoi(chi.URLParam(r, "stepID"))
	if err != nil {
		http.Error(w, "Invalid Flight Plan Step ID", http.StatusBadRequest)
		return
	}

	stepCount, err := a.FlightPlanRepo.StepCount(int32(planID))
	if err != nil {
		log.Error().AnErr("DeleteFlightPlanStep", err).Msg("Failed to delete flight plan step")
		http.Error(w, "Failed to delete flight plan step", http.StatusInternalServerError)
		return
	}

	if stepCount < 1 {
		http.Error(w, "only one step remaining in plan, cannot delete", http.StatusInternalServerError)
		return
	}

	err = a.FlightPlanRepo.DeleteFlightFromPlan(int32(stepID))
	if err != nil {
		log.Error().AnErr("DeleteFlightPlanStep", err).Msg("Failed to delete flight plan step")
		http.Error(w, "Failed to delete flight plan step", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
