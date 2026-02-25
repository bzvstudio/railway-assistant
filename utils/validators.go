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

func IsWorkspaceIncluded() bool {
	return env.GetBool("INCLUDE_WORKSPACE", true)
}

func IsStatusIncluded() bool {
	return env.GetBool("INCLUDE_STATUS", true)
}

func IsGitDedailsIncluded() bool {
	return env.GetBool("INCLUDE_BRANCH", true) || env.GetBool("INCLUDE_COMMIT", true) || env.GetBool("INCLUDE_AUTHOR", true)
}

func IsCommitIncluded() bool {
	return env.GetBool("INCLUDE_COMMIT", true)
}

func IsAuthorIncluded() bool {
	return env.GetBool("INCLUDE_AUTHOR", true)
}

func IsBranchIncluded() bool {
	return env.GetBool("INCLUDE_BRANCH", true)
}
