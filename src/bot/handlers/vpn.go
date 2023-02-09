package handlers

import (
	"fmt"
	tg "gopkg.in/telebot.v3"
)

var (
	ButtonDisconnect = tg.Btn{Text: "Отключиться от VPN", Unique: "btnDisconnect"}
	ButtonLostKey    = tg.Btn{Text: "Я потерял ключ", Unique: "btnLostKey"}
)

func (h *Handlers) HandleVPN(c tg.Context) error {
	if h.service.IsConnected(c.Chat().ID) {
		buttons := &tg.ReplyMarkup{}
		message := "Вы уже подключены к VPN."
		buttons.Inline(buttons.Row(ButtonDisconnect, ButtonLostKey))
		return c.Send(message, buttons)
	}

	message := "Чтобы пользоваться VPN нужно сразу заплатить деньгу за месяц. Бот сам пришлёт следующий счёт оплаты."
	_, err := c.Bot().Send(c.Recipient(), message)
	if err != nil {
		return err
	}
	invoice, err := h.service.GetInvoice()
	if err != nil {
		return err
	}
	_, err = invoice.Send(c.Bot(), c.Recipient(), nil)
	return err
}

func (h *Handlers) HandleAcceptPayment(c tg.Context) error {
	return c.Accept()
}

func (h *Handlers) HandleSuccessPayment(c tg.Context) error {
	//if h.service.IsConnected(c.Chat().ID) {
	//	// продление подписки
	//}
	key, err := h.service.Connect(c.Chat().ID, c.Sender().Username)
	if err != nil {
		return err
	}
	message := fmt.Sprintf("Ключ доступа для Outline:\n`%s`", key.AccessUrl)
	return c.Send(message, tg.ModeMarkdown)
}

func (h *Handlers) HandleDisconnect(c tg.Context) error {
	if err := h.service.Disconnect(c.Chat().ID); err != nil {
		return err
	}
	return c.Send("Вы отключены от совместного использования VPN")
}

func (h *Handlers) HandleLostKey(c tg.Context) error {
	key, err := h.service.FindKey(c.Chat().ID)
	if err != nil {
		return err
	}
	message := fmt.Sprintf("Твой ключ найден, нажми на него, чтобы скопировать в буфер обмена:\n`%s`", key.AccessUrl)
	return c.Send(message, tg.ModeMarkdown)
}
