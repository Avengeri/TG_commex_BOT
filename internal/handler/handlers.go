package handler

import (
	"TG_commex_BOT/internal/constans"
	"TG_commex_BOT/internal/keyboard"
	"TG_commex_BOT/internal/model"
	"TG_commex_BOT/internal/service"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type Handle struct {
	bot     *tgbotapi.BotAPI
	update  *tgbotapi.Update
	service *service.UserService
}

// Хендлер для обработки сообщений
func HandleMessage(bot *tgbotapi.BotAPI, user *model.User, update *tgbotapi.Update, service *service.UserService) {
	switch update.Message.Text {
	case constans.BUTTON_REPLY_TEXT_REGISTER:
		exists, err := service.Auth.CheckUserService(user)
		if err != nil {
			log.Printf("Не удалось проверить пользователя", err)
		}
		if !exists {
			err = service.Auth.SetUserService(user)
			if err != nil {
				log.Printf("Ошибка добваление пользователя")
			}
			msg := tgbotapi.NewMessage(user.ChatId, fmt.Sprintf("%s, спасибо за регистрацию! Пользователь под ником %s создан", user.UserName, user.UserName))
			_, err = bot.Send(msg)
			if err != nil {
				log.Printf("Не удалось отправить сообщение")
			}
		} else {
			msg := tgbotapi.NewMessage(user.ChatId, "Пользователь уже существует,пожалуйста авторизируйтесь")
			_, err = bot.Send(msg)
			if err != nil {
				log.Printf("Не удалось отправить сообщение")
			}
		}
	case constans.BUTTON_REPLY_TEXT_AUTHORIZE:
		exists, err := service.Auth.CheckUserService(user)
		if err != nil {
			log.Printf("Не удалось проверить пользователя", err)
		}
		if !exists {
			msg := tgbotapi.NewMessage(user.ChatId, "Наши миньоны не нашли Вас в базе данных, пожалуйста зарегистрируйтесь")
			_, err = bot.Send(msg)
			if err != nil {
				log.Printf("Не удалось отправить сообщение")
			}
		} else {
			keyboard.ShowMenu(bot, user)
		}
	case constans.BUTTON_REPLY_TEXT_INFO:
		msg := tgbotapi.NewMessage(user.ChatId, fmt.Sprintf("Имя пользователя: %s\n ChatId: %d", user.UserName, user.ChatId))
		_, err := bot.Send(msg)
		if err != nil {
			log.Printf("Не удалось отправить сообщение")
		}
	}

}

// Хендлер для обработки callback-запросов
func HandleCallbackQuery(bot *tgbotapi.BotAPI, update *tgbotapi.Update, service *service.UserService) {
	switch update.CallbackQuery.Data {
	}
}

// Хендлер для обработки команд
func HandleCommands(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	switch update.Message.Command() {
	case "start":
		user := model.UserUpdate(update)
		keyboard.ShowStartMessage(bot, user)
	}
}
