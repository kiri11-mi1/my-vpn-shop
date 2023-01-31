package main

import (
	tg "gopkg.in/telebot.v3"
	"log"
	"os"
	"time"
)

func main() {
	providerToken := os.Getenv("PROVIDER_TOKEN")
	pref := tg.Settings{
		Token:  os.Getenv("TELEGRAM_TOKEN"),
		Poller: &tg.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tg.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(b.Me.Username, "start working...")

	b.Handle("/start", func(c tg.Context) error {
		return c.Send(START)
	})

	b.Handle("/buy", func(c tg.Context) error {
		// TODO: add method for get actual price vpn
		prices := []tg.Price{{Label: "Актуальная цена за этот месяц", Amount: 10000}}
		file := tg.File{FileURL: InvoiceImage}
		invoice := tg.Invoice{
			Title:       InvoiceTitle,
			Description: InvoiceDescription,
			Payload:     InvoicePayload,
			Currency:    InvoiceCurrency,
			Token:       providerToken,
			Prices:      prices,
			Photo:       &tg.Photo{File: file},
		}

		_, err := invoice.Send(b, c.Recipient(), nil)
		return err
	})

	b.Start()
}
