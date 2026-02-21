package utils

import "railway-assistant/env"

func IsTelegramEnabled() bool {
	return env.GetBool("TELEGRAM_ENABLED", false) && env.GetString("TELEGRAM_BOT_TOKEN", "") != "" && env.GetString("TELEGRAM_CHAT_ID", "") != ""
}
