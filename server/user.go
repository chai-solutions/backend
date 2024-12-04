package server

import (
	"encoding/json"
	"net/http"
	// "chai/middleware"
)

type CreateUserBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a *App) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var body CreateUserBody

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

	err = a.UserRepo.CreateUser(body.Name, body.Email, body.Password)
	if err != nil {
		http.Error(w, "error creating account: "+err.Error(), http.StatusInternalServerError)
	}
}

// func (a *App) UserInfoHandler(w http.ResponseWriter, r *http.Request) {
// 	user := middleware.MustGetUserFromContext(r.Context())
//
// 	err := json.NewEncoder(w).Encode(user)
// 	_ = err
// }
