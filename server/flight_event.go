package server

import (
	"chai/database/sqlc"
	"chai/utils"
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
)

func (a *App) sendPushNotification(userIDs []string, eventType, flightNumber string) {
	heading := fmt.Sprintf("Flight %s %s", flightNumber, eventType)
	content := fmt.Sprintf("Your flight %s has been %s", flightNumber, eventType)

	payload := utils.ConstructNotificationPayload(userIDs, heading, content)
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
		userIDs = append(userIDs, fmt.Sprintf(
			"%s",
			row.PublicID.Bytes,
		))
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
