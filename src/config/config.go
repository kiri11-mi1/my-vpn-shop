package config

import (
	"github.com/caarlos0/env/v6"
	"log"
)

type Config struct {
	ApiUrl               string  `env:"VPN_URL_API"`
	TotalVpnPrice        float64 `env:"TOTAL_VPN_PRICE"`
	TelegramToken        string  `env:"TELEGRAM_TOKEN"`
	AdminChatId          string  `env:"ADMIN_CHAT_ID"`
	ProviderToken        string  `env:"PROVIDER_TOKEN"`
	PostgresUser         string  `env:"POSTGRES_USER"`
	PostgresPassword     string  `env:"POSTGRES_PASSWORD"`
	PostgresNameDatabase string  `env:"POSTGRES_DB"`
	PostgresHost         string  `env:"POSTGRES_HOST"`
}

var cfg Config

func Get() Config {
	if err := env.Parse(&cfg); err != nil {
		log.Fatalln("Config", err)
	}
	return cfg
}
