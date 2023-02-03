package handlers

import tg "gopkg.in/telebot.v3"

func HandleHelp(c tg.Context) error {
	message := "/vpn - Подключить или отключить VPN"
	return c.Send(message)
}
