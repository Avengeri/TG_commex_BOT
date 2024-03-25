package main

import (
	srv "TG_commex_BOT"
	"TG_commex_BOT/internal/repository/postgres"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

func main() {
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

	wh := tgbotapi.NewWebhook("http://localhost:8080" + bot.Token)

	_, err = bot.SetWebhook(wh)
	if err != nil {
		log.Fatalf("Failed to set webhook: %s", err.Error())
	}

	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}
	updates := bot.ListenForWebhook("/" + bot.Token)

	server := new(srv.Server)

	port := os.Getenv("SRV_PORT")

	localServer := fmt.Sprintf("the handlergrpc is running on: http://localhost:%s/", port)
	localServer = fmt.Sprintf("the handlergrpc is running on: http://localhost:%s/", port)
	fmt.Println(localServer)

	localPingPong := fmt.Sprintf("ping pong handlergrpc: http://localhost:%s/ping", port)
	fmt.Println(localPingPong)

	go func() {
		err = server.Run(port)
		if err != nil {
			log.Fatalf("the handlergrpc could not be started: %s", err.Error())
		}
	}()
	for update := range updates {
		log.Printf("%+v\n", update)
	}
}
