package utils

import (
	"fmt"
	"railway-assistant/types"
	"strings"
)

const statusMessageLimit = 2500

func GetStatusEmoji(payload types.RailwayAlert) string {
	status := strings.ToUpper(payload.Details.Status)
	severity := strings.ToUpper(payload.Severity)

	if status == "SUCCESS" {
		return "🟢"
	}

	switch severity {
	case "ERROR":
		return "🔴"
	case "WARNING", "WARN":
		return "🟠"
	case "INFO":
		return "🔵"
	}

	eventType := strings.ToLower(payload.Type)
	if strings.Contains(eventType, "fail") || strings.Contains(eventType, "error") || strings.Contains(eventType, "crash") {
		return "🔴"
	}
	if strings.Contains(eventType, "success") {
		return "🟢"
	}

	return "🔵"
}

func FormatEventType(t string) string {
	upper := strings.ToUpper(t)
	text := strings.ReplaceAll(upper, "_", " ")
	text = strings.ReplaceAll(text, ".", " ")

	if strings.HasPrefix(upper, "DEPLOYMENT") || strings.HasPrefix(upper, "DEPLOY") {
		if strings.Contains(upper, "REDEPLOY") {
			return "Redeployed"
		}
		if strings.Contains(upper, "DEPLOYED") {
			return "Deployed"
		}
		if strings.Contains(upper, "BUILDING") {
			return "Building"
		}
		if strings.Contains(upper, "DEPLOYING") {
			return "Deploying"
		}
		if upper == "DEPLOY" {
			return "Deployed"
		}
	}

	return TitleCase(text)
}

func TitleCase(s string) string {
	s = strings.ToLower(s)
	parts := strings.Fields(s)
	for i, p := range parts {
		if len(p) > 0 {
			parts[i] = strings.ToUpper(p[:1]) + p[1:]
		}
	}
	return strings.Join(parts, " ")
}

func GetIncidentStatusEmoji(status string) string {
	switch strings.ToUpper(status) {
	case "TRIAGE", "INVESTIGATING":
		return "🟠"
	case "IDENTIFIED":
		return "🟡"
	case "MONITORING":
		return "🔵"
	case "RESOLVED":
		return "🟢"
	case "POSTMORTEM":
		return "📝"
	}

	return "🔵"
}

func GetImpactEmoji(impact string) string {
	switch strings.ToUpper(impact) {
	case "MAJOR_OUTAGE":
		return "🔴"
	case "PARTIAL_OUTAGE":
		return "🟠"
	case "DEGRADED_PERFORMANCE":
		return "🟡"
	}

	return "🟢"
}

func GetIncidentURL(incident types.Incident) string {
	return fmt.Sprintf("https://status.railway.com/incident/%s", incident.Slug)
}

func TruncateText(text string, limit int) string {
	runes := []rune(text)
	if len(runes) <= limit {
		return text
	}
	return strings.TrimSpace(string(runes[:limit])) + "…"
}

func FormatComponent(c types.IncidentComponent) (string, string) {
	name := c.Name
	if c.GroupName != "" {
		name = fmt.Sprintf("%s / %s", c.GroupName, c.Name)
	}
	return name, TitleCase(strings.ReplaceAll(c.Impact, "_", " "))
}
