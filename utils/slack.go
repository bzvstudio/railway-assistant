package utils

import (
	"fmt"
	"railway-assistant/env"
	"railway-assistant/services"
	"railway-assistant/types"
	"strings"
)

func PrepareSlackMessage(payload types.RailwayAlert) []services.SlackBlock {
	var blocks []services.SlackBlock

	emoji := GetStatusEmoji(payload)
	eventType := strings.ToUpper(FormatEventType(payload.Type))

	headerText := fmt.Sprintf("%s %s", emoji, eventType)

	blocks = append(blocks, services.SlackBlock{
		"type": "header",
		"text": map[string]interface{}{
			"type":  "plain_text",
			"text":  headerText,
			"emoji": true,
		},
	})

	var fields []map[string]string

	projectInfo := fmt.Sprintf("*Project:*\n%s", payload.Resource.Project.Name)
	if payload.Resource.Service != nil {
		projectInfo += fmt.Sprintf(" / %s", payload.Resource.Service.Name)
	}
	fields = append(fields, map[string]string{
		"type": "mrkdwn",
		"text": projectInfo,
	})

	envInfo := fmt.Sprintf("*Environment:*\n%s", payload.Resource.Environment.Name)
	if payload.Resource.Environment.IsEphemeral {
		envInfo += " (Ephemeral)"
	}
	fields = append(fields, map[string]string{
		"type": "mrkdwn",
		"text": envInfo,
	})

	blocks = append(blocks, services.SlackBlock{
		"type":   "section",
		"fields": fields,
	})

	if env.GetBool("INCLUDE_STATUS", true) && payload.Details.Status != "" {
		statusText := fmt.Sprintf("*Status:*\n%s", TitleCase(payload.Details.Status))

		blocks = append(blocks, services.SlackBlock{
			"type": "section",
			"fields": []map[string]string{
				{
					"type": "mrkdwn",
					"text": statusText,
				},
			},
		})
	}

	if env.GetBool("INCLUDE_COMMIT", true) && payload.Details.CommitMessage != "" {
		blocks = append(blocks, services.SlackBlock{
			"type": "divider",
		})

		var lineParts []string

		if env.GetBool("INCLUDE_BRANCH", true) && payload.Details.Branch != "" {
			lineParts = append(lineParts, fmt.Sprintf("🌿 *%s*", payload.Details.Branch))
		}

		lineParts = append(lineParts, fmt.Sprintf("💬 _%s_", payload.Details.CommitMessage))

		if env.GetBool("INCLUDE_AUTHOR", true) && payload.Details.CommitAuthor != "" {
			lineParts = append(lineParts, fmt.Sprintf("👤 *%s*", payload.Details.CommitAuthor))
		}

		gitText := strings.Join(lineParts, "  •  ")

		blocks = append(blocks, services.SlackBlock{
			"type": "section",
			"text": map[string]string{
				"type": "mrkdwn",
				"text": gitText,
			},
		})
	}

	if payload.Resource.Deployment != nil && payload.Resource.Service != nil {
		deployURL := fmt.Sprintf("https://railway.com/project/%s/service/%s?environmentId=%s&deploymentId=%s",
			payload.Resource.Project.ID,
			payload.Resource.Service.ID,
			payload.Resource.Environment.ID,
			payload.Resource.Deployment.ID,
		)

		blocks = append(blocks, services.SlackBlock{
			"type": "actions",
			"elements": []map[string]interface{}{
				{
					"type": "button",
					"text": map[string]interface{}{
						"type":  "plain_text",
						"text":  "View Deployment 🚀",
						"emoji": true,
					},
					"url": deployURL,
				},
			},
		})
	}

	return blocks
}
