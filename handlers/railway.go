package handlers

import (
	"encoding/json"
	"log"
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

	if !utils.IsAtLeastOneProviderEnabled() {
		http.Error(w, "At least one notification provider must be enabled", http.StatusBadRequest)
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
				log.Printf("Recovered from panic in alert processing: %v\n", r)
			}
		}()

		if utils.IsTelegramEnabled() {
			message, deployURL := utils.PrepareTelegramMessage(payload)
			if err := services.SendTelegramMessage(message, deployURL); err != nil {
				log.Printf("Failed to send telegram message: %v\n", err)
			}
		}

		if utils.IsSlackEnabled() {
			slackBlocks := utils.PrepareSlackMessage(payload)
			if err := services.SendSlackMessage(slackBlocks); err != nil {
				log.Printf("Failed to send slack message: %v\n", err)
			}
		}
	}()
}
