package server

import (
	"chai/middleware"
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

func (a *App) NotificationHandler(w http.ResponseWriter, r *http.Request) {

	user := middleware.MustGetUserFromContext(r.Context())

	notifications, err := a.NotificationsRepo.GetUserNotifications(user.ID)
	if err != nil {
		log.Error().AnErr("GetUserNotifications", err).Msg("Failed to get notifications")
		http.Error(w, "Failed to get notifications", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(notifications); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}

}
