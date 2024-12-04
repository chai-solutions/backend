package mocks

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"sync"

	"chai/database/sqlc"
	"chai/repos"
)

type MockSessionRepository struct {
	Sessions map[string]sqlc.GetUserFromSessionContextRow
	UserRepo repos.UserRepository
	mu       sync.RWMutex
}

func NewMockSessionRepository(userRepo repos.UserRepository) *MockSessionRepository {
	return &MockSessionRepository{
		Sessions: make(map[string]sqlc.GetUserFromSessionContextRow),
		UserRepo: userRepo,
	}
}

func generateSessionToken() (string, error) {
	bytes := make([]byte, 42)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

// AddSession adds a session to the mock repository
func (m *MockSessionRepository) AddSession(userID int32) (string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	user, err := m.UserRepo.GetUserByID(userID)
	if err != nil || user == nil {
		return "", err
	}

	token, err := generateSessionToken()
	if err != nil {
		return "", err
	}

	m.Sessions[token] = sqlc.GetUserFromSessionContextRow{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		PublicID:  user.PublicID,
		Token:     token,
	}

	return token, nil
}

// GetUserFromSessionContext retrieves a user by session token
func (m *MockSessionRepository) GetUserFromSessionContext(token string) (*sqlc.GetUserFromSessionContextRow, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	session, exists := m.Sessions[token]
	if !exists {
		return nil, errors.New("session not found")
	}

	user, err := m.GetUserFromSessionContext(token)
	if err != nil || user == nil {
		return nil, errors.New("user not found")
	}

	return &session, nil
}

func (m *MockSessionRepository) DeleteSession(token string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.Sessions[token]; !exists {
		return nil
	}

	delete(m.Sessions, token)
	return nil
}
