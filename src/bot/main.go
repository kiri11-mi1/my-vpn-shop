package main

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

func main() {
	bot, err := tg.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	providerToken := os.Getenv("PROVIDER_TOKEN")
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tg.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		var response tg.Chattable
		switch update.Message.Command() {
		case "start":
			response = tg.NewMessage(update.Message.Chat.ID, START)
		case "buy":
			var prices = []tg.LabeledPrice{{Label: "Цена за месяц", Amount: 100}}
			response = tg.NewInvoice(
				update.Message.Chat.ID,
				InvoiceTitle,
				InvoiceDescription,
				InvoicePayload,
				providerToken,
				StartParameter,
				InvoiceCurrency,
				prices,
			)
		case "help":
			response = tg.NewMessage(update.Message.Chat.ID, HELP)
		default:
			response = tg.NewMessage(update.Message.Chat.ID, NotKnownCommand)
		}
		if _, err := bot.Send(response); err != nil {
			log.Panic(err)
		}
	}
}
