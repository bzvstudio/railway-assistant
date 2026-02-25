package utils

import "railway-assistant/env"

func IsTelegramEnabled() bool {
	return env.GetBool("TELEGRAM_ENABLED", false) && env.GetString("TELEGRAM_BOT_TOKEN", "") != "" && env.GetString("TELEGRAM_CHAT_ID", "") != ""
}

func IsSlackEnabled() bool {
	return env.GetBool("SLACK_ENABLED", false) && env.GetString("SLACK_WEBHOOK_URL", "") != ""
}

func IsAtLeastOneProviderEnabled() bool {
	return IsTelegramEnabled() || IsSlackEnabled()
}
