package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"chai/utils"

	"github.com/rs/zerolog/log"
)

type WebhookPayload struct {
	Type   string `json:"type"`
	Flight struct {
		Number string `json:"number"`
	} `json:"flight"`
}

func (a *App) WebhookHandler(w http.ResponseWriter, r *http.Request) {
	var payload WebhookPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Error().Err(err).Msg("Failed to decode request payload")
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	log.Info().Str("type", payload.Type).Str("flight_number", payload.Flight.Number).Msg("got webhook payload")

	eventType := payload.Type
	flightNumber := payload.Flight.Number

	rows, err := a.FlightsRepo.UsersOnFlight(flightNumber)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user IDs and flight status")
		return
	}

	var userIDs []int32
	var publicUserIDs []string
	for _, row := range rows {
		publicUserIDs = append(publicUserIDs, utils.PGTypeUUIDToString(row.PublicID))
		userIDs = append(userIDs, row.UserID)
	}

	onesignalNotification := getNotificationPayload(publicUserIDs, eventType, flightNumber)

	err = a.NotificationsRepo.CreateNotifications(
		userIDs,
		onesignalNotification.Headings["en"],
		onesignalNotification.Contents["en"],
	)
	if err != nil {
		log.Error().AnErr("CreateNotification", err).Msg("Failed to create notification")
		return
	}

	if err := utils.SendNotification(onesignalNotification); err != nil {
		log.Error().Err(err).Msg("Failed to send notification")
	}

	w.WriteHeader(http.StatusOK)
}

func getNotificationPayload(userIDs []string, eventType string, flightNumber string) utils.NotificationPayload {
	var heading, content string

	switch eventType {
	case "flight/delay":
		heading = "Flight delayed"
		content = fmt.Sprintf("Your flight %s has been delayed", flightNumber)
	case "flight/cancelled":
		heading = "Flight cancelled"
		content = fmt.Sprintf("Your flight %s has been cancelled", flightNumber)
	case "flight/gate-reassignment":
		heading = "Flight gate reassigned"
		content = fmt.Sprintf("Your flight %s has a new gate assignment", flightNumber)
	default:
		panic(fmt.Sprintf("unknown event type: %s", eventType))
	}

	payload := utils.ConstructNotificationPayload(userIDs, heading, content)

	return payload
}
