package handlers

import (
	tg "gopkg.in/telebot.v3"
	"my-vpn-shop/config"
	"my-vpn-shop/outline"
	"my-vpn-shop/subscription"
)

func HandleVPN(c tg.Context) error {
	api := outline.NewOutlineClient(config.Get().ApiUrl)
	keys, err := api.GetKeys()
	if err != nil {
		return err
	}
	invoice, err := subscription.GetInvoice(len(keys), config.Get().TotalVpnPrice, config.Get().ProviderToken)
	if err != nil {
		return err
	}
	_, err = invoice.Send(c.Bot(), c.Recipient(), nil)
	return err
}
