package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {

	bot, err := tgbotapi.NewBotAPI("api")
	if err != nil {
		log.Panic(err)
	}

	_, err = bot.SetWebhook(tgbotapi.NewWebhook("https://your.domain.com/telegram_bot_endpoint"))
	if err != nil {
		log.Fatal(err)
	}
	
	updates := bot.ListenForWebhook("/telegram_bot_endpoint")


	for update := range updates {
		if update.Message == nil {
			continue
		}

		
		if update.Message.Text == "" {
			continue
		}

	
		msg := tgbotapi.NewMessage("api", update.Message.Text)

		
		_, err := bot.Send(msg)
		if err != nil {
			log.Println(err)
		}
	}
}
