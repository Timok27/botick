package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {

	bot, err := tgbotapi.NewBotAPI("...")
	if err != nil {
		log.Panic(err)
	}

	_, err = bot.SetWebhook(tgbotapi.NewWebhook("https://your.domain.com/telegram_bot_endpoint"))
	if err != nil {
		log.Fatal(err)
	}

	// Канал, через который будут приходить обновления от пользователя
	updates := bot.ListenForWebhook("/telegram_bot_endpoint")

	// Получение обновлений из канала и их обработка
	for update := range updates {
		if update.Message == nil { // игнорируем обновления, которые не являются сообщениями
			continue
		}

		// Если сообщение не содержит текста, пропускаем его
		if update.Message.Text == "" {
			continue
		}

		// Пересылаем текст сообщения вам
		msg := tgbotapi.NewMessage(482924915, update.Message.Text)

		// Отправка сообщения вам
		_, err := bot.Send(msg)
		if err != nil {
			log.Println(err)
		}
	}
}
