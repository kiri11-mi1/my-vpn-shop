package handlers

import (
	"fmt"
	tg "gopkg.in/telebot.v3"
)

func (h *Handlers) HandleConnections(c tg.Context) error {
	count, err := h.service.GetCountSubs()
	if err != nil {
		return nil
	}
	message := fmt.Sprintf("Количество подключений на данный момент: %d", count)
	return c.Send(message)
}
