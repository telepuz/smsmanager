package telegram

import (
	"log/slog"

	"github.com/telepuz/smsmanager/internal/config"
)

type Telegram struct {
	token string
}

func New(cfg *config.Config) (*Telegram, error) {
	slog.Debug("Creating new telegram messenger...")
	return &Telegram{
		token: cfg.Messenger.Token,
	}, nil
}
