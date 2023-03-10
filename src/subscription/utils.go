package subscription

import (
	"fmt"
	tg "gopkg.in/telebot.v3"
	"log"
	"time"
)

func GetActualPrice(subCount int, totalVPNPrice float64) (tg.Price, error) {
	if subCount == 0 {
		return tg.Price{}, ErrZeroKeysInServer
	}
	if totalVPNPrice < 0 {
		return tg.Price{}, ErrNegativeTotalPrice
	}
	tmp := totalVPNPrice / float64(subCount)
	if tmp < MinAmount {
		return tg.Price{Label: Label, Amount: int(MinAmount * 100)}, nil
	}

	return tg.Price{Label: Label, Amount: int(tmp * 100.00)}, nil
}

func GetName(sub string, id int64) string {
	return fmt.Sprintf("sub_%s_%d", sub, id)
}

func IsPayDay(lastPayTime, currentTime time.Time) bool {
	nextPayTime := lastPayTime.AddDate(0, 1, 0)

	monthCompare := nextPayTime.Month() == currentTime.Month()
	dayCompare := nextPayTime.Day() == currentTime.Day()
	yearCompare := nextPayTime.Year() == currentTime.Year()
	return monthCompare && dayCompare && yearCompare
}

func IsTimeOutPay(lastPayTime, currentTime time.Time) bool {
	nextPayTime := lastPayTime.AddDate(0, 1, 0)

	monthCompare := currentTime.Month() >= nextPayTime.Month()
	yearCompare := currentTime.Year() >= nextPayTime.Year()
	dayCompare := currentTime.Day() > nextPayTime.Day()

	return monthCompare && dayCompare && yearCompare
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetInvoice(countSubs int, providerToken string, totalVpnPrice float64) (tg.Invoice, error) {
	price, err := GetActualPrice(countSubs, totalVpnPrice)
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
