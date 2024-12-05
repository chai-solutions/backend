package repos

import (
	"context"

	"chai/database/sqlc"
)

type FlightsRepository interface {
	FlightByCode(flightCode string) (*sqlc.GetFlightRow, error)
	FlightsByDepartureArrival(departureCode string, arrivalCode string) ([]sqlc.GetFlightsRow, error)
}

type flightsRepositoryImpl struct {
	db *sqlc.Queries
}

func NewFlightsRepository(db *sqlc.Queries) FlightsRepository {
	return &flightsRepositoryImpl{db: db}
}

func (r *flightsRepositoryImpl) FlightByCode(flightCode string) (*sqlc.GetFlightRow, error) {
	flight, err := r.db.GetFlight(context.Background(), flightCode)
	if err != nil {
		return nil, err
	}

	return &flight, err
}

func (r *flightsRepositoryImpl) FlightsByDepartureArrival(departureCode string, arrivalCode string) ([]sqlc.GetFlightsRow, error) {
	flights, err := r.db.GetFlights(context.Background(), sqlc.GetFlightsParams{
		Dep: departureCode,
		Arr: arrivalCode,
	})
	if err != nil {
		return nil, err
	}

	return flights, nil
}
