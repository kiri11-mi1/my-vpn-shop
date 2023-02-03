package handlers

import (
	"fmt"
	tg "gopkg.in/telebot.v3"
)

func HandleConnections(c tg.Context) error {
	message := fmt.Sprintf("В разработке")
	return c.Send(message)
}
