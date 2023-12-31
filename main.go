package main

import (
	"fmt"
	tgbotapi "github.com/crocone/tg-bot"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var bot *tgbotapi.BotAPI
var startMenu = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Say hi", "Hello"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Say bay", "Larevedere"),
	),
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".env not load")
	}
	bot, err = tgbotapi.NewBotAPI(os.Getenv("TG_API_BOT_TOKEN")) // Remove the ":="
	if err != nil {
		log.Fatalf("Failed to initialize tg bot api: %v", err)
	}
	// Rest of your code...

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	if err != nil {
		log.Fatalf("Failed to start listnening for updates %v", err)
	}

	for update := range updates {
		if update.CallbackQuery != nil {
			callbacks(update)
		} else if update.Message.IsCommand() {
			commands(update)
		} else {
			//simply message
		}
	}
}
func callbacks(update tgbotapi.Update) {
	data := update.CallbackQuery.Data
	chatId := update.CallbackQuery.From.ID
	firstName := update.CallbackQuery.From.FirstName
	lastName := update.CallbackQuery.From.LastName
	var text string
	switch data {
	case "Hello":
		text = fmt.Sprintf("Hello %v %v", firstName, lastName)
	case "Larevedere":
		text = fmt.Sprintf("Pa pa %v %v", firstName, lastName)
	default:
		text = "Unknown command"
	}

	msg := tgbotapi.NewMessage(chatId, text)
	sendMessage(msg)

}

func commands(update tgbotapi.Update) {
	command := update.Message.Command()
	switch command {
	case "start":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Chose action")
		msg.ReplyMarkup = startMenu
		msg.ParseMode = "Markdown"
		sendMessage(msg)
	}
}
func sendMessage(msg tgbotapi.Chattable) {
	if _, err := bot.Send(msg); err != nil {
		log.Panicf("Send message error %v", err)
	}
}
