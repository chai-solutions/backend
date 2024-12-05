package mocks

import (
	"sync"

	"chai/database/sqlc"
	"chai/repos"
)

type MockAirportsRepository struct {
	Airports map[int32]sqlc.Airport
	mu       sync.RWMutex
}

func NewMockAirportsRepository() repos.AirportsRepository {
	mockAirports := map[int32]sqlc.Airport{
		1: {
			ID:   1,
			Iata: "OAK",
			Name: "Oakland International Airport",
		},
		2: {
			ID:   2,
			Iata: "SFO",
			Name: "San Francisco International Airport",
		},
		3: {
			ID:   3,
			Iata: "SJC",
			Name: "San-Jose Mineta International Airport",
		},
	}

	return &MockAirportsRepository{
		Airports: mockAirports,
	}
}

func (m *MockAirportsRepository) GetAll() ([]sqlc.Airport, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	var airports []sqlc.Airport
	for _, v := range m.Airports {
		airports = append(airports, v)
	}

	return airports, nil
}

func (m *MockAirportsRepository) GetByID(id int32) (*sqlc.Airport, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if airport, exists := m.Airports[id]; exists {
		return &airport, nil
	}

	return nil, nil
}
