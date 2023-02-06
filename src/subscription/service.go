package subscription

import (
	"fmt"
	tg "gopkg.in/telebot.v3"
	"my-vpn-shop/db"
)

const MinAmount = 60.00 // valid minimal value for telegram payments
const Label = "Актуальная цена за этот месяц"

func getActualPrice(keysCount int, totalVPNPrice float64) (tg.Price, error) {
	if keysCount == 0 {
		return tg.Price{}, ErrZeroKeysInServer
	}
	if totalVPNPrice < 0 {
		return tg.Price{}, ErrNegativeTotalPrice
	}
	tmp := totalVPNPrice / float64(keysCount)
	if tmp < MinAmount {
		return tg.Price{Label: Label, Amount: int(MinAmount * 100)}, nil
	}

	return tg.Price{Label: Label, Amount: int(tmp * 100.00)}, nil
}

func GetInvoice(keysCount int, totalVpnPrice float64, providerToken string) (tg.Invoice, error) {
	price, err := getActualPrice(keysCount, totalVpnPrice)
	if err != nil {
		return tg.Invoice{}, err
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
	return invoice, nil
}

func IsConnected(subscribers []db.Subscriber, chatId int64) bool {
	for _, sub := range subscribers {
		if sub.ID == chatId {
			return true
		}
	}
	return false
}

func GetName(sub string, id int64) string {
	return fmt.Sprintf("sub_%s_%d", sub, id)
}
