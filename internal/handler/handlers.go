package handler

import (
	"TG_commex_BOT/internal/keyboard"
	"TG_commex_BOT/internal/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

// Хендлер для обработки сообщений
func HandleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	// Ваш код обработки сообщения здесь
	log.Printf("Received message from user %d: %s", message.From.ID, message.Text)
}

// Хендлер для обработки callback-запросов
func HandleCallbackQuery(bot *tgbotapi.BotAPI, callbackQuery *tgbotapi.CallbackQuery) {
	// Ваш код обработки callback-запроса здесь
	log.Printf("Received callback query from user %d: %s", callbackQuery.From.ID, callbackQuery.Data)
}
func HandleCommands(bot *tgbotapi.BotAPI, message *tgbotapi.Message, update *tgbotapi.Update) {
	switch message.Command() {
	case "start":
		user := model.UserUpdate(update)
		keyboard.ShowStartMessage(bot, user)
	}
}
