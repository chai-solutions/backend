package mocks

import (
	"chai/config"
	"chai/server"
)

func InitializeMockApp() *server.App {
	mockUserRepo := NewMockUserRepository()
	mockSessionRepo := NewMockSessionRepository(mockUserRepo)
	mockAirportsRepo := NewMockAirportsRepository()
	mockFlightsRepo := NewMockFlightsRepository(mockAirportsRepo)

	err := mockUserRepo.CreateUser(
		"test@example.com",
		"Test User",
		"password123",
	)
	if err != nil {
		panic("could not initialize user data: " + err.Error())
	}

	cfg := config.AppConfig{}

	app := server.NewApp(cfg, server.Repositories{
		UserRepo:     mockUserRepo,
		SessionRepo:  mockSessionRepo,
		AirportsRepo: mockAirportsRepo,
		FlightsRepo:  mockFlightsRepo,
	})
	app.RegisterRoutes()

	return app
}
