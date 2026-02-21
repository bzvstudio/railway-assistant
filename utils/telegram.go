package utils

import (
	"fmt"
	"railway-assistant/env"
	"railway-assistant/types"
	"strings"
)

func PrepareTelegramMessage(payload types.RailwayAlert) (string, string) {
	var sb strings.Builder
	emoji := GetStatusEmoji(payload)

	sb.WriteString("🚄 *Railway Alert*\n\n")

	eventType := strings.ToUpper(FormatEventType(payload.Type))
	sb.WriteString(fmt.Sprintf("%s *%s*\n\n", emoji, escapeMarkdown(eventType)))

	sb.WriteString("*Details*\n")
	sb.WriteString("📍 ")
	pathParts := []string{}

	if env.GetBool("INCLUDE_WORKSPACE", true) && payload.Resource.Workspace.Name != "" {
		pathParts = append(pathParts, escapeMarkdown(payload.Resource.Workspace.Name))
	}

	pathParts = append(pathParts, escapeMarkdown(payload.Resource.Project.Name))

	if payload.Resource.Service != nil {
		pathParts = append(pathParts, escapeMarkdown(payload.Resource.Service.Name))
	}

	sb.WriteString(strings.Join(pathParts, " / "))

	sb.WriteString(fmt.Sprintf("\n🌍 %s", escapeMarkdown(payload.Resource.Environment.Name)))
	if payload.Resource.Environment.IsEphemeral {
		sb.WriteString(" \\(Ephemeral\\)")
	}

	if env.GetBool("INCLUDE_STATUS", true) && payload.Details.Status != "" {
		sb.WriteString(fmt.Sprintf("\n🔄 Status: %s", escapeMarkdown(TitleCase(payload.Details.Status))))
	}
	sb.WriteString("\n\n")

	var gitSb strings.Builder
	hasGitInfo := false

	if env.GetBool("INCLUDE_COMMIT", true) && payload.Details.CommitMessage != "" {
		gitSb.WriteString(fmt.Sprintf("💬 _%s_\n", escapeMarkdown(payload.Details.CommitMessage)))
		hasGitInfo = true
	}

	metaParts := []string{}

	if env.GetBool("INCLUDE_BRANCH", true) && payload.Details.Branch != "" {
		metaParts = append(metaParts, fmt.Sprintf("🌱 %s", escapeMarkdown(payload.Details.Branch)))
	}

	if env.GetBool("INCLUDE_AUTHOR", true) && payload.Details.CommitAuthor != "" {
		metaParts = append(metaParts, fmt.Sprintf("👤 %s", escapeMarkdown(payload.Details.CommitAuthor)))
	}

	if len(metaParts) > 0 {
		gitSb.WriteString(strings.Join(metaParts, "  •  "))
		gitSb.WriteString("\n")
		hasGitInfo = true
	}

	if hasGitInfo {
		sb.WriteString("*Git*\n")
		sb.WriteString(gitSb.String())
	}

	var deployURL string
	if payload.Resource.Deployment != nil && payload.Resource.Service != nil {
		deployURL = fmt.Sprintf("https://railway.com/project/%s/service/%s?environmentId=%s&deploymentId=%s",
			payload.Resource.Project.ID,
			payload.Resource.Service.ID,
			payload.Resource.Environment.ID,
			payload.Resource.Deployment.ID,
		)
	}

	return sb.String(), deployURL
}

func escapeMarkdown(text string) string {
	replacer := strings.NewReplacer(
		"_", "\\_",
		"*", "\\*",
		"[", "\\[",
		"]", "\\]",
		"(", "\\(",
		")", "\\)",
		"~", "\\~",
		"`", "\\`",
		">", "\\>",
		"#", "\\#",
		"+", "\\+",
		"-", "\\-",
		"=", "\\=",
		"|", "\\|",
		"{", "\\{",
		"}", "\\}",
		".", "\\.",
		"!", "\\!",
		"\\", "\\\\",
	)
	return replacer.Replace(text)
}
