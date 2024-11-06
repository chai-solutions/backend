package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/rs/zerolog/log"
)

type NotificationPayload struct {
	AppID                  string            `json:"app_id"`
	IncludeExternalUserIDs []string          `json:"include_external_user_ids"`
	Headings               map[string]string `json:"headings"`
	Contents               map[string]string `json:"contents"`
}

func GetAppID() string {
	return os.Getenv("ONESIGNAL_APP_ID")
}

func GetAPIKey() string {
	return os.Getenv("ONESIGNAL_API_KEY")
}

func ConstructNotificationPayload(userIDs []string, heading, content string) NotificationPayload {
	return NotificationPayload{
		AppID:                  GetAppID(),
		IncludeExternalUserIDs: userIDs,
		Headings: map[string]string{
			"en": heading,
		},
		Contents: map[string]string{
			"en": content,
		},
	}
}

func SendNotification(payload NotificationPayload) error {
	apiKey := GetAPIKey()
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
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", apiKey))

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
