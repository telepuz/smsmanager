package telegram

import (
	"github.com/telepuz/smsmanager/internal/config"
)

type Telegram struct {
	token string
}

func New(cfg *config.Config) (*Telegram, error) {
	return &Telegram{
		token: cfg.Messenger.Token,
	}, nil
}
