package server

// import (
// 	"encoding/json"
// 	"net/http"
//
// 	"github.com/rs/zerolog/log"
// )
//
// type WebhookPayload struct {
// 	Type   string `json:"type"`
// 	Flight struct {
// 		Number string `json:"number"`
// 	} `json:"flight"`
// }
//
// func (a *App) WebhookHandler(w http.ResponseWriter, r *http.Request) {
// 	var payload WebhookPayload
//
// 	log.Info().Msg("Received webhook request")
//
// 	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
// 		log.Error().Err(err).Msg("Failed to decode request payload")
// 		http.Error(w, "Invalid request payload", http.StatusBadRequest)
// 		return
// 	}
//
// 	log.Info().Str("type", payload.Type).Str("flight_number", payload.Flight.Number).Msg("Parsed webhook payload")
//
// 	a.handleFlightEvent(payload.Type, payload.Flight.Number)
//
// 	log.Info().Msg("Successfully handled webhook event")
//
// 	w.WriteHeader(http.StatusOK)
// }
