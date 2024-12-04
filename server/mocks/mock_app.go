package mocks

import (
	"chai/config"
	"chai/server"
)

func InitializeMockApp() *server.App {
	mockUserRepo := NewMockUserRepository()

	// TODO: Populate mock data if needed
	mockUserRepo.CreateUser(
		"test@example.com",
		"Test User",
		"password123",
	)

	cfg := config.AppConfig{}

	app := server.NewApp(cfg, mockUserRepo)
	app.RegisterRoutes()

	return app
}
