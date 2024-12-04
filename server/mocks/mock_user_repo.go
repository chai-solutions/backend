package mocks

import (
	"errors"
	"math"
	"sync"

	"chai/database/sqlc"
)

func GetUnusedID(users map[int32]sqlc.User) int32 {
	for i := int32(1); i < math.MaxInt32; i++ {
		if _, exists := users[i]; !exists {
			return i
		}
	}
	panic("no available IDs")
}

type MockUserRepository struct {
	Data map[int32]sqlc.User
	mu   sync.RWMutex
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		Data: make(map[int32]sqlc.User),
	}
}

func (m *MockUserRepository) CreateUser(name string, email string, password string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, user := range m.Data {
		if user.Email == email {
			return errors.New("user already exists")
		}
	}

	newID := GetUnusedID(m.Data)

	m.Data[newID] = sqlc.User{
		ID:       newID,
		Name:     name,
		Email:    email,
		Password: password,
	}

	return nil
}

func (m *MockUserRepository) GetUserByEmail(email string) (*sqlc.User, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, user := range m.Data {
		if user.Email == email {
			return &user, nil
		}
	}

	return nil, errors.New("user not found")
}
