package repos

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"time"

	"chai/database/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

type SessionRepository interface {
	AddSession(userID int32) (string, error)
	GetUserFromSessionContext(token string) (*sqlc.GetUserFromSessionContextRow, error)
	DeleteSession(token string) error
}

type sessionRepositoryImpl struct {
	db *sqlc.Queries
}

func NewSessionRepository(db *sqlc.Queries) SessionRepository {
	return &sessionRepositoryImpl{db: db}
}

func generateSessionToken() (string, error) {
	bytes := make([]byte, 42)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

func (r *sessionRepositoryImpl) AddSession(userID int32) (string, error) {
	token, err := generateSessionToken()
	if err != nil {
		return "", err
	}
	expiryDate := time.Now().Add(24 * time.Hour * 30)

	bruh, err := r.db.CreateSession(context.Background(),
		sqlc.CreateSessionParams{
			UserID: userID,
			Token:  token,
			ExpiryTime: pgtype.Timestamp{
				Time:  expiryDate,
				Valid: true,
			},
		})

	return bruh.Token, err
}

func (r *sessionRepositoryImpl) GetUserFromSessionContext(token string) (*sqlc.GetUserFromSessionContextRow, error) {
	users, err := r.db.GetUserFromSessionContext(context.Background(), token)
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, nil
	}

	return &users[0], nil
}

func (r *sessionRepositoryImpl) DeleteSession(token string) error {
	err := r.db.DeleteSession(context.Background(), token)
	return err
}
