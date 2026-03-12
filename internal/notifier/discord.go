package notifier

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func SendDiscordNotification(webhookURL string, message string) error {
	payload := map[string]string {
		"content": message,
	}

	jsonPayload, _ := json.Marshal(payload)

	_, err := http.Post(webhookURL,"application/json",bytes.NewBuffer(jsonPayload))

	return	 err
}

