package mocks

import (
	"chai/config"
	"chai/server"
)

func InitializeMockApp() *server.App {
	mockUserRepo := NewMockUserRepository()
	mockSessionRepo := NewMockSessionRepository(mockUserRepo)

	err := mockUserRepo.CreateUser(
		"test@example.com",
		"Test User",
		"password123",
	)
	if err != nil {
		panic("could not initialize user data: " + err.Error())
	}

	cfg := config.AppConfig{}

	app := server.NewApp(cfg, mockUserRepo, mockSessionRepo)
	app.RegisterRoutes()

	return app
}
