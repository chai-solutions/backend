package server

import (
	"context"
	"encoding/json"
	"net/http"

	"chai/database/sqlc"

	"golang.org/x/crypto/bcrypt"
)

type createUserBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a *App) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var body createUserBody

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "malformed JSON", http.StatusBadRequest)
		return
	}

	if body.Name == "" {
		http.Error(w, "missing name", http.StatusBadRequest)
		return
	}
	if body.Email == "" {
		http.Error(w, "missing email", http.StatusBadRequest)
		return
	}
	if len(body.Password) < 8 {
		http.Error(w, "password length must be at least 8 characters", http.StatusBadRequest)
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
	if err != nil {
		http.Error(w, "password hash error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	createdUser, err := a.Queries.CreateAccount(context.Background(), sqlc.CreateAccountParams{
		Name:         body.Name,
		Email:        body.Email,
		PasswordHash: string(passwordHash),
	})
	if err != nil {
		http.Error(w, "error creating account: "+err.Error(), http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(createdUser)
	_ = err
}
