package handlers

import tg "gopkg.in/telebot.v3"

func HandleStart(c tg.Context) error {
	message := "Привет, друг! Тут ты можешь купить подписку на доступ VPN. Стоимость будет рассчитываться в зависимости от подключенных клиентов."
	return c.Send(message)
}
