package subscription

import (
	"fmt"
	tg "gopkg.in/telebot.v3"
	"log"
	"time"
)

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
