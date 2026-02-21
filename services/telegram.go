package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"railway-assistant/env"
)

func SendTelegramMessage(text string, buttonURL ...string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", env.GetString("TELEGRAM_BOT_TOKEN", ""))
	payload := map[string]interface{}{
		"chat_id":              env.GetString("TELEGRAM_CHAT_ID", ""),
		"text":                 text,
		"parse_mode":           "MarkdownV2",
		"link_preview_options": map[string]bool{"is_disabled": true},
	}

	if len(buttonURL) > 0 && buttonURL[0] != "" {
		payload["reply_markup"] = map[string]interface{}{
			"inline_keyboard": [][]map[string]string{
				{
					{"text": "View Deployment", "url": buttonURL[0]},
				},
			},
		}
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("telegram api returned status: %d, body: %s", resp.StatusCode, string(respBody))
	}

	return nil
}
