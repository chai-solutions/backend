package server

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"time"

	"chai/database/sqlc"
	"chai/middleware"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type loginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginRes struct {
	Authorization string `json:"authorization"`
}

func (a *App) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var body loginReq

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "malformed JSON", http.StatusBadRequest)
		return
	}

	if body.Email == "" {
		http.Error(w, "missing email", http.StatusBadRequest)
		return
	}
	if body.Password == "" {
		http.Error(w, "missing password", http.StatusBadRequest)
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
	if err != nil {
		http.Error(w, "password hash error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	log.Info().Msg(string(passwordHash))

	users, err := a.Queries.SelectAccountByEmail(context.Background(), body.Email)
	if err != nil {
		http.Error(w, "error logging in: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if len(users) == 0 {
		http.Error(w, "account not found", http.StatusForbidden)
		return
	}
	user := users[0]

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		http.Error(w, "incorrect password", http.StatusForbidden)
	}

	token, err := generateSessionToken()
	if err != nil {
		log.Warn().Err(err).Send()
		http.Error(w, "session token generation error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Expire 30 days from now.
	expiryDate := time.Now().Add(24 * time.Hour * 30)
	_, err = a.Queries.CreateSession(context.Background(), sqlc.CreateSessionParams{
		UserID: user.ID,
		Token:  token,
		ExpiryTime: pgtype.Timestamp{
			Time:  expiryDate,
			Valid: true,
		},
	})
	log.Info().Int32("userID", user.ID).Send()
	if err != nil {
		log.Warn().Err(err).Send()
		http.Error(w, "session insertion error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(loginRes{
		Authorization: token,
	})
	_ = err
}

func (a *App) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	user := middleware.MustGetUserFromContext(r.Context())

	err := a.Queries.DeleteSession(context.Background(), user.Token)
	if err != nil {
		log.Warn().Err(err).Send()
		http.Error(w, "session deletion error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func generateSessionToken() (string, error) {
	bytes := make([]byte, 42)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}
