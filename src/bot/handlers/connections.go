package handlers

import (
	"fmt"
	tg "gopkg.in/telebot.v3"
)

func (h *Handlers) HandleConnections(c tg.Context) error {
	subs, err := h.storage.GetSubscribers()
	if err != nil {
		return nil
	}
	message := fmt.Sprintf("Количество подключений на данный момент: %d", len(subs))
	return c.Send(message)
}
