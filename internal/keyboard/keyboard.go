package keyboard

import (
	"TG_commex_BOT/internal/constans"
	"TG_commex_BOT/internal/model"
	"fmt"
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
	replyRow := newKeyboardRow(constans.BUTTON_REPLY_TEXT_AUTHORIZE)
	replyRow_2 := newKeyboardRow(constans.BUTTON_REPLY_TEXT_REGISTER)
	replyRow_3 := newKeyboardRow(constans.BUTTON_REPLY_TEXT_INFO)

	replyKeyboard := tgbotapi.NewReplyKeyboard(replyRow, replyRow_2, replyRow_3)

	msg.ReplyMarkup = replyKeyboard

	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Не удалось отправить сообщение")
	}
}

func ShowMenu(bot *tgbotapi.BotAPI, user *model.User) {

	msg := tgbotapi.NewMessage(user.ChatId, fmt.Sprintf("Ну привет %s! Выбери что-нибудь", user.UserName))
	replyRow := newKeyboardRow(constans.BUTTON_REPLY_TEXT_PRICE_5_MIN)
	replyRow_2 := newKeyboardRow(constans.BUTTON_REPLY_TEXT_PRICE_1_HOUR)
	replyRow_3 := newKeyboardRow(constans.BUTTON_REPLY_TEXT_PRICE_1_DAY)

	replyKeyboard := tgbotapi.NewReplyKeyboard(replyRow, replyRow_2, replyRow_3)

	msg.ReplyMarkup = replyKeyboard

	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Не удалось отправить сообщение")
	}
}
