package mocks

import (
	"errors"
	"math"
	"sync"
	"time"

	"chai/database/sqlc"
	"chai/repos"

	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
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

func NewMockUserRepository() repos.UserRepository {
	mockUsers := map[int32]sqlc.User{
		// TODO: add uuids
		1: {
			ID:        1,
			Name:      "Alice",
			Email:     "alice@example.com",
			Password:  "$2a$14$Vl/o5y6wOmIT8uQCDrdp2uO6zkzANM1KmmW/X6jYrMrApF2jdCOre", // password
			CreatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
		},
		2: {
			ID:        2,
			Name:      "Bob",
			Email:     "bob@example.com",
			Password:  "$2a$14$Vl/o5y6wOmIT8uQCDrdp2uO6zkzANM1KmmW/X6jYrMrApF2jdCOre", // password
			CreatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
		},
		3: {
			ID:        3,
			Name:      "Sanjay Ramaswamy",
			Email:     "sanjay@ramaswamy.net",
			Password:  "$2a$14$UfFC/SdbF2cJmvhXdAruTufD20Tfysk7sFTA4uOz7iCKNNIwuUdLW", // ramaswamy123
			CreatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
		},
	}

	return &MockUserRepository{
		Data: mockUsers,
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
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	m.Data[newID] = sqlc.User{
		ID:       newID,
		Name:     name,
		Email:    email,
		Password: string(passwordHash),
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

func (m *MockUserRepository) GetUserByID(userID int32) (*sqlc.User, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, user := range m.Data {
		if user.ID == userID {
			return &user, nil
		}
	}

	return nil, errors.New("user not found")
}
