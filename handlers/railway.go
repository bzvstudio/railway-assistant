package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"railway-assistant/services"
	"railway-assistant/types"
	"railway-assistant/utils"
)

func RailwayAlertsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var payload types.RailwayAlert
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if payload.Type == "" {
		http.Error(w, "Invalid Payload: missing type", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered from panic in alert processing: %v\n", r)
			}
		}()

		message, deployURL := utils.PrepareTelegramMessage(payload)
		fmt.Printf("Sending Telegram message:\n%s\nURL: %s\n", message, deployURL)

		if err := services.SendTelegramMessage(message, deployURL); err != nil {
			fmt.Printf("Failed to send telegram message: %v\n", err)
		}
	}()
}
