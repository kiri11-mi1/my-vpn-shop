package main

import (
	"fmt"
	_ "github.com/lib/pq"
	tg "gopkg.in/telebot.v3"
	"log"
	"my-vpn-shop/bot/handlers"
	"my-vpn-shop/config"
	"my-vpn-shop/db"
	"my-vpn-shop/outline"
	"my-vpn-shop/storage"
	"my-vpn-shop/subscription"
	"time"
)

func main() {
	pref := tg.Settings{
		Token:  config.Get().TelegramToken,
		Poller: &tg.LongPoller{Timeout: 10 * time.Second},
	}
	b, err := tg.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s sslmode=disable",
		config.Get().PostgresUser,
		config.Get().PostgresPassword,
		config.Get().PostgresNameDatabase,
		config.Get().PostgresHost,
	)
	pgClient, err := db.Connect("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("Connected to database successfully")
	if err := pgClient.CreateTables(); err != nil {
		log.Fatal(err)
		return
	}
	pgDB := pgClient.Client()
	pgStorage := storage.NewSQlDB(pgDB)

	outlineAPI := outline.NewOutlineClient(config.Get().ApiUrl)
	subscriptionService := subscription.NewSubscriptionService(pgStorage, outlineAPI)
	handlerManager := handlers.NewHandlers(subscriptionService)

	b.Handle("/start", handlers.HandleStart)
	b.Handle("/vpn", handlerManager.HandleVPN)
	b.Handle(&handlers.ButtonDisconnect, handlerManager.HandleDisconnect)
	b.Handle(&handlers.ButtonLostKey, handlerManager.HandleLostKey)
	b.Handle(tg.OnPayment, handlerManager.HandleSuccessPayment)
	b.Handle(tg.OnCheckout, handlerManager.HandleAcceptPayment)
	b.Handle("/help", handlers.HandleHelp)
	b.Handle("/connections", handlerManager.HandleConnections)

	log.Println(b.Me.Username, "start working...")
	b.Start()
}
