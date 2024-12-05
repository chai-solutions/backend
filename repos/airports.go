package repos

import (
	"context"

	"chai/database/sqlc"
)

type AirportsRepository interface {
	GetAll() ([]sqlc.Airport, error)
	GetByID(id int32) (*sqlc.Airport, error)
}

type airportsRepositoryImpl struct {
	db *sqlc.Queries
}

func NewAirportsRepository(db *sqlc.Queries) AirportsRepository {
	return &airportsRepositoryImpl{db: db}
}

func (r *airportsRepositoryImpl) GetAll() ([]sqlc.Airport, error) {
	airports, err := r.db.GetAllAirports(context.Background())
	if err != nil {
		return nil, err
	}

	return airports, nil
}

func (r *airportsRepositoryImpl) GetByID(id int32) (*sqlc.Airport, error) {
	airport, err := r.db.GetAirportByID(context.Background(), id)
	if err != nil {
		return nil, err
	}

	return &airport, nil
}
