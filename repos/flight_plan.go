package repos

import (
	"context"

	"chai/database/sqlc"
)

type FlightPlanRepository interface {
	CreatePlan(userID int32, initialFlightNumber string) (*sqlc.FlightPlanFlight, error)
	AddFlightToPlan(planID int32, flightNumber string) (*sqlc.FlightPlanFlight, error)
	GetPlansForUser(userID int32) ([]sqlc.GetFlightPlansRow, error)
	Exists(planID int32) (bool, error)
	GetPlan(planID int32) ([]sqlc.GetFlightPlanRow, error)
	DeletePlan(planID int32) error
	StepCount(planID int32) (int64, error)
	DeleteFlightFromPlan(stepID int32) error
}

type flightPlanRepositoryImpl struct {
	db *sqlc.Queries
}

func NewFlightPlanRepository(db *sqlc.Queries) FlightPlanRepository {
	return &flightPlanRepositoryImpl{db: db}
}

func (r *flightPlanRepositoryImpl) CreatePlan(userID int32, initialFlightNumber string) (*sqlc.FlightPlanFlight, error) {
	createdPlanStep, err := r.db.CreateFlightPlan(context.Background(), sqlc.CreateFlightPlanParams{
		Flightnumber: initialFlightNumber,
		Users:        userID,
	})
	if err != nil {
		return nil, err
	}

	return &createdPlanStep, nil
}

func (r *flightPlanRepositoryImpl) AddFlightToPlan(planID int32, flightNumber string) (*sqlc.FlightPlanFlight, error) {
	createdPlanStep, err := r.db.PatchFlightPlan(context.Background(), sqlc.PatchFlightPlanParams{
		FlightPlan:   planID,
		FlightNumber: flightNumber,
	})
	if err != nil {
		return nil, err
	}

	return &createdPlanStep, nil
}

func (r *flightPlanRepositoryImpl) GetPlansForUser(userID int32) ([]sqlc.GetFlightPlansRow, error) {
	plans, err := r.db.GetFlightPlans(context.Background(), userID)
	if err != nil {
		return nil, err
	}

	return plans, nil
}

func (r *flightPlanRepositoryImpl) GetPlan(planID int32) ([]sqlc.GetFlightPlanRow, error) {
	plan, err := r.db.GetFlightPlan(context.Background(), planID)
	if err != nil {
		return nil, err
	}

	return plan, nil
}

func (r *flightPlanRepositoryImpl) DeletePlan(planID int32) error {
	err := r.db.DeleteFlightPlan(context.Background(), planID)
	return err
}

func (r *flightPlanRepositoryImpl) DeleteFlightFromPlan(stepID int32) error {
	err := r.db.DeleteFlightPlanStep(context.Background(), stepID)
	return err
}

func (r *flightPlanRepositoryImpl) Exists(planID int32) (bool, error) {
	exists, err := r.db.FlightPlanExists(context.Background(), planID)
	if err != nil {
		return false, err
	}
	return exists, err
}

func (r *flightPlanRepositoryImpl) StepCount(planID int32) (int64, error) {
	count, err := r.db.GetFlightPlanStepCount(context.Background(), planID)
	if err != nil {
		return 0, err
	}
	return count, nil
}
