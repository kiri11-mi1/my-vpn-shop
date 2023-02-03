package main

import (
	tg "gopkg.in/telebot.v3"
	"log"
	"my-vpn-shop/bot/handlers"
	"os"
	"time"
)

func main() {

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

	b.Handle("/start", handlers.HandleStart)
	b.Handle("/vpn", handlers.HandleVPN)
	b.Handle(tg.OnCheckout, func(c tg.Context) error {
		return c.Accept()
	})
	b.Handle("/help", handlers.HandleHelp)
	b.Handle("/connections", handlers.HandleConnections)
	b.Start()
}
