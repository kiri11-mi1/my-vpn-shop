package handlers

import (
	tg "gopkg.in/telebot.v3"
	"my-vpn-shop/outline"
	"my-vpn-shop/subscription"
	"os"
	"strconv"
)

func HandleVPN(c tg.Context) error {
	api := outline.NewOutlineClient(os.Getenv("VPN_URL_API"))
	totalVpnPrice, err := strconv.ParseFloat(os.Getenv("TOTAL_VPN_PRICE"), 64)
	providerToken := os.Getenv("PROVIDER_TOKEN")
	if err != nil {
		return err
	}
	keys, err := api.GetKeys()
	if err != nil {
		return err
	}
	invoice, err := subscription.GetInvoice(len(keys), totalVpnPrice, providerToken)
	if err != nil {
		return err
	}
	_, err = invoice.Send(c.Bot(), c.Recipient(), nil)
	return err
}
