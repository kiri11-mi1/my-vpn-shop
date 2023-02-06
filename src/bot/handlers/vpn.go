package handlers

import (
	tg "gopkg.in/telebot.v3"
	"log"
	"my-vpn-shop/config"
	"my-vpn-shop/outline"
	"my-vpn-shop/subscription"
)

var (
	ButtonDisconnect = tg.Btn{Text: "Отключиться от VPN", Unique: "btnDisconnect"}
	ButtonLostKey    = tg.Btn{Text: "Я потерял ключ", Unique: "btnLostKey"}
)

func (h *Handlers) HandleVPN(c tg.Context) error {
	subs, err := h.storage.GetSubscribers()
	if err != nil {
		return err
	}

	if !subscription.IsConnected(subs, c.Chat().ID) {
		message := "Чтобы пользоваться VPN нужно сразу заплатить деньгу за месяц. Бот сам пришлёт следующий счёт оплаты."
		_, err := c.Bot().Send(c.Recipient(), message)
		if err != nil {
			return err
		}
		invoice, err := subscription.GetInvoice(len(subs)+1, config.Get().TotalVpnPrice, config.Get().ProviderToken)
		if err != nil {
			return err
		}
		_, err = invoice.Send(c.Bot(), c.Recipient(), nil)
		return err
	}
	buttons := &tg.ReplyMarkup{}
	message := "Вы уже подключены к VPN."
	buttons.Inline(buttons.Row(ButtonDisconnect, ButtonLostKey))
	return c.Send(message, buttons)
}

func (h *Handlers) HandleAcceptPayment(c tg.Context) error {
	return c.Accept()
}

func (h *Handlers) HandleSuccessPayment(c tg.Context) error {
	newSub, err := h.storage.InsertSubscriber(c.Chat().ID, c.Sender().Username)
	if err != nil {
		return err
	}
	log.Println("Add sub: ", newSub.ID, newSub.Name)
	createdKey, err := h.api.GetAccess(subscription.GetName(newSub.Name, newSub.ID))
	if err != nil {
		return err
	}
	log.Println("Create access key", createdKey.Id, createdKey.Name, "for sub", newSub.Name, newSub.ID)
	dbKey, err := h.storage.InsertAccessKey(createdKey.Id, createdKey.Name, createdKey.AccessUrl, newSub)
	if err != nil {
		return err
	}
	log.Println("Add key in db:", dbKey.Name, dbKey.ID, dbKey.Subscriber.Name)
	message := "Ключ доступа для Outline: " + dbKey.AccessUrl
	return c.Send(message)
}

func (h *Handlers) HandleDisconnect(c tg.Context) error {
	key, err := h.storage.GetKeyBySubId(c.Sender().ID)
	if err != nil {
		return err
	}
	if err := h.api.DeleteKey(outline.AccessKey{Id: key.ID}); err != nil {
		return err
	}
	if err := h.storage.DeleteSubscriber(c.Sender().ID); err != nil {
		return err
	}
	return c.Send("Вы отключены от совместного использования VPN")
}

func (h *Handlers) HandleLostKey(c tg.Context) error {
	key, err := h.storage.GetKeyBySubId(c.Sender().ID)
	if err != nil {
		return err
	}
	message := "Твой ключ найден: " + key.AccessUrl
	return c.Send(message)
}
