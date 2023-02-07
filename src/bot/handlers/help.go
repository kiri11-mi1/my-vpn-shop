package handlers

import tg "gopkg.in/telebot.v3"

func HandleHelp(c tg.Context) error {
	message := "" +
		"/vpn - Подключить или отключить VPN\n" +
		"/connections - Количество подключенных пользователей\n" +
		"/howuse - Инструкция использования ключа доступа"
	return c.Send(message)
}
