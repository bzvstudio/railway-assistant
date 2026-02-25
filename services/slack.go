package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"railway-assistant/env"
)

type SlackBlock map[string]interface{}

func SendSlackMessage(blocks []SlackBlock) error {
	webhookURL := env.GetString("SLACK_WEBHOOK_URL", "")
	if webhookURL == "" {
		return fmt.Errorf("SLACK_WEBHOOK_URL is not set")
	}

	payload := map[string]interface{}{
		"blocks": blocks,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("slack api returned status: %d, body: %s", resp.StatusCode, string(respBody))
	}

	return nil
}
