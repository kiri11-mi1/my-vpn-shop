package subscription

import tg "gopkg.in/telebot.v3"

func GetActualPrice(keysCount int, totalVPNPrice float64) (tg.Price, error) {
	if keysCount == 0 {
		return tg.Price{}, ZeroKeysInServer
	}
	var (
		tmp    = totalVPNPrice / float64(keysCount)
		result = tg.Price{Label: "Актуальная цена за этот месяц", Amount: int(tmp * 100.00)}
	)

	return result, nil
}
