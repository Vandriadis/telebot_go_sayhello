# Telegram bot

This is a simple Telegram bot written in Go using the  tg-bot  library. The bot responds to commands and inline queries, and can send messages with inline keyboards.

## Usage

To use the bot, you need to create a Telegram bot and get an API token. You can follow the [official instructions](https://core.telegram.org/bots#6-botfather) to create a bot and get a token.

Once you have the token, create a  .env  file in the root directory of the project with the following content:

```TG_API_BOT_TOKEN=<your_token_here>```

Then, run the bot with the following command:
```
go run main.go
```

The bot will listen for updates and respond to commands and inline queries.

## Features

The bot currently supports the following features:

-  /start  command: displays a start menu with inline buttons.
- Inline buttons: the bot responds to inline buttons with custom messages.
- Unknown commands: the bot responds with a default message for unknown commands.
