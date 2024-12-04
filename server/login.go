package server

import (
	"encoding/json"
	"net/http"

	"chai/middleware"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginRes struct {
	Authorization string `json:"authorization"`
}

func (a *App) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var body LoginRequestBody

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

	user, err := a.UserRepo.GetUserByEmail(body.Email)
	if err != nil {
		http.Error(w, "error logging in: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.Error(w, "account not found", http.StatusForbidden)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		http.Error(w, "incorrect password", http.StatusForbidden)
	}

	if err != nil {
		log.Warn().Err(err).Send()
		http.Error(w, "session token generation error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Expire 30 days from now.
	token, err := a.SessionRepo.AddSession(user.ID)
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

	err := a.SessionRepo.DeleteSession(user.Token)
	if err != nil {
		log.Warn().Err(err).Send()
		http.Error(w, "session deletion error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
