package server

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

func (a *App) AirportsHandler(w http.ResponseWriter, _ *http.Request) {
	airports, err := a.AirportsRepo.GetAll()
	if err != nil {
		log.Error().AnErr("AirportsHandler", err).Msg("Failed to retieve airport")
		http.Error(w, "Failed to get airports", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(airports); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
