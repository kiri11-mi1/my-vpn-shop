package main

import (
	tg "gopkg.in/telebot.v3"
	"log"
	"my-vpn-shop/outline"
	"my-vpn-shop/subscription"
	"os"
	"strconv"
	"time"
)

func main() {
	providerToken := os.Getenv("PROVIDER_TOKEN")
	totalVpnPrice, err := strconv.ParseFloat(os.Getenv("TOTAL_VPN_PRICE"), 64)
	api := outline.NewOutlineClient(os.Getenv("VPN_URL_API"))
	if err != nil {
		log.Fatal(err)
		return
	}
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
		keys, err := api.GetKeys()
		if err != nil {
			return err
		}
		price, err := subscription.GetActualPrice(len(keys), totalVpnPrice)
		if err != nil {
			log.Fatal(err)
			return err
		}
		file := tg.File{FileURL: InvoiceImage}
		invoice := tg.Invoice{
			Title:       InvoiceTitle,
			Description: InvoiceDescription,
			Payload:     InvoicePayload,
			Currency:    InvoiceCurrency,
			Token:       providerToken,
			Prices:      []tg.Price{price},
			Photo:       &tg.Photo{File: file},
		}
		_, err = invoice.Send(b, c.Recipient(), nil)
		return err
	})
	b.Handle(tg.OnCheckout, func(c tg.Context) error {
		return c.Accept()
	})
	b.Start()
}
