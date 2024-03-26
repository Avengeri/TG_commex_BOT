package model

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type User struct {
	Id           int64
	UserName     string
	PasswordHash string
	ChatId       int64
}

func UserUpdate(update *tgbotapi.Update) *User {
	userNameUpdate := update.Message.From.UserName
	chatIdUpdate := update.Message.Chat.ID

	log.Printf("Имя пользователя: %s\n Chat ID: %d\n", userNameUpdate, chatIdUpdate)

	user := &User{
		UserName: userNameUpdate,
		ChatId:   chatIdUpdate,
	}
	return user
}
