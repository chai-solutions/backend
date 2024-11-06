package server

import (
	"chai/utils"
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

}
