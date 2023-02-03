package config

import (
	"github.com/caarlos0/env/v6"
	"log"
)

type Config struct {
	ApiUrl        string  `env:"VPN_URL_API"`
	TotalVpnPrice float64 `env:"TOTAL_VPN_PRICE"`
	TelegramToken string  `env:"TELEGRAM_TOKEN"`
	AdminChatId   string  `env:"ADMIN_CHAT_ID"`
	ProviderToken string  `env:"PROVIDER_TOKEN"`
}

var cfg Config

func Get() Config {
	if err := env.Parse(&cfg); err != nil {
		log.Fatalln("Config", err)
	}
	return cfg
}
