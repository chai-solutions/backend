package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"chai/config"

	"github.com/rs/zerolog/log"
)

type NotificationPayload struct {
	AppID          string            `json:"app_id"`
	TargetChannel  string            `json:"target_channel"`
	Headings       map[string]string `json:"headings"`
	Contents       map[string]string `json:"contents"`
	IncludeAliases includeAliasType  `json:"include_aliases"`
}

type includeAliasType struct {
	ExternalID []string `json:"external_id"`
}

func ConstructNotificationPayload(userIDs []string, heading, content string) NotificationPayload {
	cfg := config.GetConfig()
	return NotificationPayload{
		AppID:         cfg.OneSignalAppID,
		TargetChannel: "push",
		Headings: map[string]string{
			"en": heading,
		},
		Contents: map[string]string{
			"en": content,
		},
		IncludeAliases: includeAliasType{
			ExternalID: userIDs,
		},
	}
}

func SendNotification(payload NotificationPayload) error {
	oneSignalAPIKey := config.GetConfig().OneSignalAPIKey

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal notification payload: %w", err)
	}

	req, err := http.NewRequest("POST", "https://onesignal.com/api/v1/notifications", bytes.NewBuffer(body))
	if err != nil {
		log.Error().Err(err).Msg("Failed to create request")
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", oneSignalAPIKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("Failed to send request")
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Error().Int("status_code", resp.StatusCode).Msg("Failed to send notification")
		return fmt.Errorf("failed to send notification, status code: %d", resp.StatusCode)
	}

	log.Info().Msg("Notification sent successfully")
	return nil
}
