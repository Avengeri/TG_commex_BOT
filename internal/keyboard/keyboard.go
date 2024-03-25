package keyboard

import (
	"TG_commex_BOT/internal/constans"
	"TG_commex_BOT/internal/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func newKeyboardRow(buttonTexts ...string) []tgbotapi.KeyboardButton {
	var buttons []tgbotapi.KeyboardButton

	for _, text := range buttonTexts {
		buttons = append(buttons, tgbotapi.NewKeyboardButton(text))
	}
	return buttons
}

func newInlineKeyboard(buttonText, buttonCode string) []tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(buttonText, buttonCode))
}

func ShowStartMessage(bot *tgbotapi.BotAPI, user *model.User) {

	msg := tgbotapi.NewMessage(user.ChatId, "Есть аккаунт-авторизуйся, нет - заведи")
	replyRow := newKeyboardRow(constans.BUTTON_REPLY_TEXT_AUTORIZE)
	replyRow_2 := newKeyboardRow(constans.BUTTON_REPLY_TEXT_LOGIN)

	replyKeyboard := tgbotapi.NewReplyKeyboard(replyRow, replyRow_2)

	msg.ReplyMarkup = replyKeyboard

	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Не удалось отправить сообщение")
	}
}
