package subscription

import (
	tg "gopkg.in/telebot.v3"
)

const MinAmount = 87.41 // valid minimal value for telegram payments
const Label = "Актуальная цена за этот месяц"

func GetActualPrice(keysCount int, totalVPNPrice float64) (tg.Price, error) {
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
