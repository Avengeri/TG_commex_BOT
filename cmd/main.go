package main

import (
	"TG_commex_BOT/internal/handler"
	"TG_commex_BOT/internal/repository/postgres"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	if err := godotenv.Load("./.env"); err != nil {
		log.Fatalf("error loading environment variable: %s", err.Error())
	}

	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("Postgres creation error: %s", err.Error())
	}
	ok, err := postgres.CheckDBConn(db)
	if err != nil {
		log.Fatalf("Connection to the database could not be established: %s", err.Error())
	}
	fmt.Println(ok)

	bot, err := tgbotapi.NewBotAPI("7151876669:AAHj9OA8TSMRYvIWTqS44en-m3sttRSuhSY")
	if err != nil {
		log.Fatalf("Connection to BOT API failed: %s", err.Error())
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatalf("Could not get updates from chan: %s", err.Error())
	}
	for update := range updates {
		if update.Message != nil && !update.Message.IsCommand() {
			handler.HandleMessage(bot, update.Message)
		} else if update.CallbackQuery != nil {
			handler.HandleCallbackQuery(bot, update.CallbackQuery)
		} else if update.Message.IsCommand() {
			handler.HandleCommands(bot, update.Message, &update)
		}
	}
}
