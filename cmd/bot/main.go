package main

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"log"
	"telegram_bot/pkg/telegram"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("1726047056:AAEM-58SLSYOq6E2QBmEdz56_Go4nbsWRys")
	if err != nil {
		log.Panic(err)
	}
	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}
