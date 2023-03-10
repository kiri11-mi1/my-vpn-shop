package handlers

import (
	"my-vpn-shop/db"
)

type SubscriptionService interface {
	IsConnected(chatId int64) bool
	Connect(chatID int64, username string) (db.AccessKey, error)
	Disconnect(chatId int64) error
	FindKey(chatID int64) (db.AccessKey, error)
	GetCountSubs() (int, error)
	Renew(chatID int64) error
}

type Handlers struct {
	service SubscriptionService
}

func NewHandlers(srv SubscriptionService) *Handlers {
	return &Handlers{service: srv}
}
