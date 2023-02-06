package handlers

import (
	"my-vpn-shop/db"
	"my-vpn-shop/outline"
)

type Storage interface {
	GetSubscribers() ([]db.Subscriber, error)
	InsertSubscriber(id int64, name string) (db.Subscriber, error)
	InsertAccessKey(id, name, accessUrl string, sub db.Subscriber) (db.AccessKey, error)
	DeleteSubscriber(id int64) error
	GetKeyBySubId(id int64) (db.AccessKey, error)
}

type API interface {
	GetKeys() (outline.AccessKeys, error)
	CreateKey() (outline.AccessKey, error)
	ChangeKeyName(name string, key outline.AccessKey) error
	DeleteKey(key outline.AccessKey) error
	GetAccess(name string) (outline.AccessKey, error)
}

type Handlers struct {
	storage Storage
	api     API
}

func NewHandlers(storage Storage, api API) *Handlers {
	return &Handlers{storage: storage, api: api}
}
