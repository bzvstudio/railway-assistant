# Railway Assistant

[![Deploy on Railway](https://railway.com/button.svg)](https://railway.com/deploy/railway-assistant?referralCode=1uJ_oA)

A production-ready webhook service designed to integrate Railway with Telegram. It acts as a bridge, receiving webhook events from your Railway projects and forwarding them as formatted notifications to a Telegram chat or group.

## Features

- **Real-time Notifications**: Get instant alerts for deployments, build failures, and service crashes via Telegram and Slack.
- **Smart Formatting**: Messages are cleanly formatted with Markdown, including status emojis and direct links to your Railway projects, services, and deployments.
- **Configurable Detail**: Control exactly what information is included in your notifications (Workspace, Branch, Commit, Author, etc.).
- **Zero Dependencies**: Built with Go standard library only. No external frameworks or bloat.
- **Ultra-Lightweight**: Designed for minimal compute resource usage and instant startup.

## Setup Guide

### 1. Create a Telegram Bot or Slack App

**Telegram:**

1.  Open Telegram and search for **@BotFather**.
2.  Send the command `/newbot`.
3.  Follow the prompts to name your bot.
4.  Copy the **HTTP API Token** provided (this is your `TELEGRAM_BOT_TOKEN`).
5.  Start a chat with your new bot (click **Start**) or add it to the group where you want to receive notifications.
6.  Get your **Chat ID**:
    - Forward a message from your chat/group to **@userinfobot**.
    - Or check `https://api.telegram.org/bot<YOUR_TOKEN>/getUpdates` after sending a message to the bot.
    - This ID is your `TELEGRAM_CHAT_ID`.

**Slack:**

1.  Create a Slack App at [api.slack.com/apps](https://api.slack.com/apps).
2.  Enable **Incoming Webhooks** for your app.
3.  Create a new Webhook URL for the channel you want to post to.
4.  Copy the **Webhook URL** (this is your `SLACK_WEBHOOK_URL`).

### 2. Configure Environment Variables

This service is configured entirely via environment variables.

**Required:**

| Variable             | Description                                     |
| :------------------- | :---------------------------------------------- |
| `TELEGRAM_ENABLED`   | Set to `true` to enable Telegram notifications. |
| `TELEGRAM_BOT_TOKEN` | The API token you got from BotFather.           |
| `TELEGRAM_CHAT_ID`   | The ID of the user or group to receive alerts.  |
| `SLACK_ENABLED`      | Set to `true` to enable Slack notifications.    |
| `SLACK_WEBHOOK_URL`  | The Webhook URL from your Slack App.            |
| `PORT`               | The port to listen on (default: `8080`).        |

**Optional (Message Customization):**

Control what information appears in your alerts by setting these to `true` or `false` (default is `true`).

| Variable            | Default | Description                                              |
| :------------------ | :------ | :------------------------------------------------------- |
| `INCLUDE_WORKSPACE` | `true`  | Show the Workspace name.                                 |
| `INCLUDE_STATUS`    | `true`  | Show the deployment status (e.g., INITIALIZING, FAILED). |
| `INCLUDE_BRANCH`    | `true`  | Show the git branch name.                                |
| `INCLUDE_COMMIT`    | `true`  | Show the commit message.                                 |
| `INCLUDE_AUTHOR`    | `true`  | Show the commit author.                                  |

### 3. Deploy & Connect

1.  **Deploy this template** to your own Railway project.
2.  Set the required Environment Variables in your new service.
3.  Once deployed, copy your service's **Public URL** (e.g., `https://railway-assistant.up.railway.app`).
4.  Go to the **Project Settings** of the Railway project you want to monitor.
5.  Navigate to the **Webhooks** section.
6.  Create a new Webhook:
    - **Payload URL**: `https://<YOUR_SERVICE_URL>/railway/alerts`
    - **Event Types**: Select the events you want to be notified about (e.g., `Deployment Failed`, `Deployment Success`, `Volume Alert`, etc.).

## License

MIT

## Example Notification

Here is an example of what a notification looks like in Telegram:

🚄 _Railway Alert_

🔵 _DEPLOYED_

_Details_
📍 My Team Workspace / My Project / Frontend App
🌍 Production
🔄 Status: Success

_Git_
💬 _feat: add new dashboard components_
🌱 main • 👤 username
