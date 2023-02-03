package handlers

import tg "gopkg.in/telebot.v3"

func HandleStart(c tg.Context) error {
	message := "Привет! Тут ты можешь купить подписку на мой впн. Стоимость будет рассчитываться в зависимости от подключенных клиентов."
	return c.Send(message)
}
