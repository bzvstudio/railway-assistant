package jobs

import (
	"log"
	"sort"
	"time"

	"railway-assistant/services"
	"railway-assistant/types"
	"railway-assistant/utils"
)

const statusPollInterval = 2 * time.Minute

type pendingUpdate struct {
	Incident types.Incident
	Update   types.IncidentUpdate
}

func StartRailwayStatusPoller() {
	log.Printf("Railway status polling enabled (every %s)\n", statusPollInterval)

	var seen map[string]struct{}

	poll := func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recovered from panic in status polling: %v\n", r)
			}
		}()

		status, err := services.FetchRailwayStatus()
		if err != nil {
			log.Printf("Failed to fetch railway status: %v\n", err)
			return
		}

		fresh, next := newUpdates(seen, status)

		if seen == nil {
			log.Printf("Railway status baseline recorded (%d updates)\n", len(next))
		} else {
			for _, p := range fresh {
				sendStatusUpdate(p)
			}
		}

		seen = next
	}

	poll()

	ticker := time.NewTicker(statusPollInterval)
	defer ticker.Stop()

	for range ticker.C {
		poll()
	}
}

func newUpdates(seen map[string]struct{}, status *types.RailwayStatus) ([]pendingUpdate, map[string]struct{}) {
	next := make(map[string]struct{})
	var fresh []pendingUpdate

	incidents := append(append([]types.Incident{}, status.ActiveIncidents...), status.RecentIncidents...)

	for _, incident := range incidents {
		for _, update := range incident.Updates {
			if _, ok := next[update.ID]; ok {
				continue
			}
			next[update.ID] = struct{}{}

			if seen == nil {
				continue
			}
			if _, ok := seen[update.ID]; !ok {
				fresh = append(fresh, pendingUpdate{Incident: incident, Update: update})
			}
		}
	}

	sort.SliceStable(fresh, func(i, j int) bool {
		return fresh[i].Update.CreatedAt.Before(fresh[j].Update.CreatedAt)
	})

	return fresh, next
}

func sendStatusUpdate(p pendingUpdate) {
	if utils.IsTelegramEnabled() {
		message, incidentURL := utils.PrepareTelegramStatusMessage(p.Incident, p.Update)
		if err := services.SendTelegramMessage(message, "View Incident", incidentURL); err != nil {
			log.Printf("Failed to send telegram status message: %v\n", err)
		}
	}

	if utils.IsSlackEnabled() {
		slackBlocks := utils.PrepareSlackStatusMessage(p.Incident, p.Update)
		if err := services.SendSlackMessage(slackBlocks); err != nil {
			log.Printf("Failed to send slack status message: %v\n", err)
		}
	}
}
