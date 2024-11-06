package server

import (
	"context"
	"fmt"

	"chai/database/sqlc"
	"chai/utils"

	"github.com/rs/zerolog/log"
)

func getNotificationPayload(userIDs []string, eventType string, flightNumber string) utils.NotificationPayload {
	var heading, content string

	switch eventType {
	case "flight/delay":
		heading = fmt.Sprintf("Flight %s delayed", flightNumber)
		content = fmt.Sprintf("Your flight %s has been delayed", flightNumber)
	case "flight/cancelled":
		heading = fmt.Sprintf("Flight %s cancelled", flightNumber)
		content = fmt.Sprintf("Your flight %s has been cancelled", flightNumber)
	case "flight/gate-reassignment":
		heading = fmt.Sprintf("Flight %s gate reassigned", flightNumber)
		content = fmt.Sprintf("Your flight %s has a new gate assignment", flightNumber)
	default:
		panic(fmt.Sprintf("unknown event type: %s", eventType))
	}

	payload := utils.ConstructNotificationPayload(userIDs, heading, content)

	return payload
}

func (a *App) sendPushNotification(userIDs []string, eventType, flightNumber string) {
	payload := getNotificationPayload(userIDs, eventType, flightNumber)

	if err := utils.SendNotification(payload); err != nil {
		log.Error().AnErr("SendNotification", err).Msg("Failed to send notification")
	}
}

func (a *App) handleFlightEvent(eventType, flightNumber string) {
	rows, err := a.Queries.GetUsersByFlightNumber(context.Background(), flightNumber)
	if err != nil {
		log.Error().AnErr("GetUserIDsByFlightNumber", err).Msg("Failed to get user IDs and flight status")
		return
	}

	var userIDs []string
	for _, row := range rows {
		userIDs = append(userIDs, utils.PGTypeUUIDToString(row.PublicID))
	}

	_, err = a.Queries.CreateNotification(context.Background(), sqlc.CreateNotificationParams{
		EventType:    eventType,
		FlightNumber: flightNumber,
	})
	if err != nil {
		log.Error().AnErr("CreateNotification", err).Msg("Failed to create notification")
		return
	}

	a.sendPushNotification(userIDs, eventType, flightNumber)
}
