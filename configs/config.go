package configs

import (
	"os"
)

const (
	prefix = "prefix"
	botToken = "bot_token"
)

type BotConfig struct {
	Prefix string
	Token string
}
var cfg *BotConfig

func Load() {
	cfg = &BotConfig{
		Prefix: os.Getenv(prefix),
		Token: os.Getenv(botToken),
	}
}

func GetToken() string {
	return cfg.Token
}

func GetPrefix() string {
	return cfg.Prefix
}
