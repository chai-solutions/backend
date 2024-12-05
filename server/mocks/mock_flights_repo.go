package mocks

import (
	"errors"
	"sync"
	"time"

	"chai/database/sqlc"
	"chai/repos"

	"github.com/jackc/pgx/v5/pgtype"
)

type MockFlightsRepository struct {
	Flights      map[string]sqlc.Flight
	AirportsRepo repos.AirportsRepository
	mu           sync.RWMutex
}

func NewMockFlightsRepository(airportsRepo repos.AirportsRepository) repos.FlightsRepository {
	return &MockFlightsRepository{
		Flights: map[string]sqlc.Flight{
			"UA123": {
				ID:            1,
				FlightNumber:  "UA123",
				Airline:       1,
				DepAirport:    1,
				ArrAirport:    2,
				SchedDepTime:  pgtype.Timestamp{Time: time.Date(2024, time.December, 5, 8, 0, 0, 0, time.UTC), Valid: true},
				SchedArrTime:  pgtype.Timestamp{Time: time.Date(2024, time.December, 5, 10, 30, 0, 0, time.UTC), Valid: true},
				ActualDepTime: pgtype.Timestamp{Time: time.Date(2024, time.December, 5, 8, 5, 0, 0, time.UTC), Valid: true},
				ActualArrTime: pgtype.Timestamp{Time: time.Date(2024, time.December, 5, 10, 45, 0, 0, time.UTC), Valid: true},
			},
			"DL456": {
				ID:            2,
				FlightNumber:  "DL456",
				Airline:       2,
				DepAirport:    1,
				ArrAirport:    3,
				SchedDepTime:  pgtype.Timestamp{Time: time.Date(2024, time.December, 5, 9, 0, 0, 0, time.UTC), Valid: true},
				SchedArrTime:  pgtype.Timestamp{Time: time.Date(2024, time.December, 5, 11, 30, 0, 0, time.UTC), Valid: true},
				ActualDepTime: pgtype.Timestamp{Time: time.Date(2024, time.December, 5, 9, 10, 0, 0, time.UTC), Valid: true},
				ActualArrTime: pgtype.Timestamp{Time: time.Date(2024, time.December, 5, 11, 50, 0, 0, time.UTC), Valid: true},
			},
		},
		AirportsRepo: airportsRepo,
	}
}

func (m *MockFlightsRepository) FlightByCode(flightCode string) (*sqlc.GetFlightRow, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if flight, exists := m.Flights[flightCode]; exists {
		depAirport, err := m.AirportsRepo.GetByID(flight.DepAirport)
		if err != nil {
			return nil, err
		}
		arrAriport, err := m.AirportsRepo.GetByID(flight.ArrAirport)
		if err != nil {
			return nil, err
		}

		return &sqlc.GetFlightRow{
			ID:            flight.ID,
			FlightNumber:  flight.FlightNumber,
			SchedDepTime:  flight.SchedDepTime,
			ActualArrTime: flight.ActualArrTime,
			ActualDepTime: flight.ActualDepTime,
			DepIata:       depAirport.Iata,
			ArrivalIata:   arrAriport.Iata,
			DepName:       depAirport.Name,
			ArrivalName:   depAirport.Iata,
		}, nil
	}

	return nil, errors.New("flight not found")
}

func (m *MockFlightsRepository) FlightsByDepartureArrival(departureCode string, arrivalCode string) ([]sqlc.GetFlightsRow, error) {
	var flights []sqlc.GetFlightsRow

	// This is SUPREMELY inefficient, but for a mock with barely any entries, it's fine.
	for _, flight := range m.Flights {
		depAirport, err := m.AirportsRepo.GetByID(flight.DepAirport)
		if err != nil {
			return nil, err
		}
		arrAriport, err := m.AirportsRepo.GetByID(flight.ArrAirport)
		if err != nil {
			return nil, err
		}

		if depAirport.Iata == departureCode && arrAriport.Iata == arrivalCode {
			flights = append(flights, sqlc.GetFlightsRow{
				ID:            flight.ID,
				FlightNumber:  flight.FlightNumber,
				SchedDepTime:  flight.SchedDepTime,
				ActualArrTime: flight.ActualArrTime,
				ActualDepTime: flight.ActualDepTime,
				ArrivalIata:   arrAriport.Iata,
				ArrivalName:   arrAriport.Name,
				DepIata:       depAirport.Iata,
				DepName:       depAirport.Name,
			})
		}
	}

	return flights, nil
}
