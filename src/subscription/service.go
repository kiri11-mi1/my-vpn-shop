package subscription

import (
	tg "gopkg.in/telebot.v3"
	"log"
	"my-vpn-shop/config"
	"my-vpn-shop/db"
	"my-vpn-shop/outline"
	"time"
)

const MinAmount = 60.00 // valid minimal value for telegram payments
const Label = "Актуальная цена за этот месяц"

type Storage interface {
	GetSubscribers() ([]db.Subscriber, error)
	InsertSubscriber(id int64, name string) (db.Subscriber, error)
	InsertAccessKey(id, name, accessUrl string, sub db.Subscriber) (db.AccessKey, error)
	DeleteSubscriber(id int64) error
	GetKeyBySubId(id int64) (db.AccessKey, error)
	GetSubByID(id int64) (db.Subscriber, error)
	UpdateSubscriberPayedAt(id int64) error
}

type API interface {
	GetKeys() (outline.AccessKeys, error)
	CreateKey() (outline.AccessKey, error)
	ChangeKeyName(name string, key outline.AccessKey) error
	DeleteKey(key outline.AccessKey) error
	GetAccess(name string) (outline.AccessKey, error)
}

type Service struct {
	storage Storage
	api     API
}

func NewSubscriptionService(storage Storage, api API) *Service {
	return &Service{storage: storage, api: api}
}

func (s *Service) GetInvoice() (tg.Invoice, error) {
	subs, err := s.storage.GetSubscribers()
	if err != nil {
		return tg.Invoice{}, err
	}
	price, err := GetActualPrice(len(subs)+1, config.Get().TotalVpnPrice)
	if err != nil {
		return tg.Invoice{}, err
	}
	file := tg.File{FileURL: InvoiceImage}
	invoice := tg.Invoice{
		Title:       InvoiceTitle,
		Description: InvoiceDescription,
		Payload:     InvoicePayload,
		Currency:    InvoiceCurrency,
		Token:       config.Get().ProviderToken,
		Prices:      []tg.Price{price},
		Photo:       &tg.Photo{File: file},
	}
	return invoice, nil
}

func (s *Service) IsConnected(chatId int64) bool {
	if _, err := s.storage.GetSubByID(chatId); err != nil {
		return false
	}
	return true
}

func (s *Service) Connect(chatID int64, username string) (db.AccessKey, error) {
	newSub, err := s.storage.InsertSubscriber(chatID, username)
	if err != nil {
		return db.AccessKey{}, err
	}
	log.Println("Add sub: ", newSub.ID, newSub.Name)

	createdKey, err := s.api.GetAccess(GetName(newSub.Name, newSub.ID))
	if err != nil {
		return db.AccessKey{}, err
	}
	log.Println("Create access key", createdKey.Id, createdKey.Name, "for sub", newSub.Name, newSub.ID)
	dbKey, err := s.storage.InsertAccessKey(createdKey.Id, createdKey.Name, createdKey.AccessUrl, newSub)
	if err != nil {
		return db.AccessKey{}, err
	}
	log.Println("Add key in db:", dbKey.Name, dbKey.ID, dbKey.Subscriber.Name)

	return dbKey, nil
}

func (s *Service) Disconnect(chatID int64) error {
	key, err := s.storage.GetKeyBySubId(chatID)
	if err != nil {
		return err
	}
	if err := s.api.DeleteKey(outline.AccessKey{Id: key.ID}); err != nil {
		return err
	}
	if err := s.storage.DeleteSubscriber(chatID); err != nil {
		return err
	}
	return nil
}

func (s *Service) FindKey(chatID int64) (db.AccessKey, error) {
	key, err := s.storage.GetKeyBySubId(chatID)
	if err != nil {
		return db.AccessKey{}, err
	}
	return key, nil
}

func (s *Service) GetCountSubs() (int, error) {
	subs, err := s.storage.GetSubscribers()
	if err != nil {
		return -1, err
	}
	return len(subs), nil
}

func (s *Service) Renew(chatID int64) error {
	if err := s.storage.UpdateSubscriberPayedAt(chatID); err != nil {
		return err
	}
	return nil
}

func (s *Service) CheckPayDateTask(bot *tg.Bot) {
	for true {
		subs, err := s.storage.GetSubscribers()
		CheckError(err)
		if len(subs) == 0 {
			continue
		}
		for _, sub := range subs {
			recipient := &tg.Chat{ID: sub.ID}
			now := time.Now()
			if IsPayDay(sub.PayedAt, now) {
				message := "Пора платить за VPN"
				_, err := bot.Send(recipient, message)
				CheckError(err)

				invoice, err := s.GetInvoice()
				CheckError(err)

				_, err = invoice.Send(bot, recipient, nil)
				CheckError(err)
			}
			if IsTimeOutPay(sub.PayedAt, now) {
				message := "Вы просрочили оплату. Поэтому я отключаю вас от использования VPN :c"
				_, err := bot.Send(recipient, message)
				CheckError(err)
				CheckError(s.Disconnect(sub.ID))
			}
		}
		time.Sleep(24 * time.Hour)
	}
}
