package handlers

import (
	"fmt"
	tg "gopkg.in/telebot.v3"
)

func (h *Handlers) HandleConnections(c tg.Context) error {
	message := fmt.Sprintf("В разработке")
	return c.Send(message)
}
