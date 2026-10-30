package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"railway-assistant/types"
	"time"
)

const RailwayStatusURL = "https://status.railway.com/api/status"

var statusClient = &http.Client{Timeout: 15 * time.Second}

func FetchRailwayStatus() (*types.RailwayStatus, error) {
	resp, err := statusClient.Get(RailwayStatusURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("railway status api returned status: %d", resp.StatusCode)
	}

	var status types.RailwayStatus
	if err := json.NewDecoder(resp.Body).Decode(&status); err != nil {
		return nil, err
	}

	return &status, nil
}
