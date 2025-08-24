package messenger

import (
	"github.com/telepuz/smsmanager/internal"
	"github.com/telepuz/smsmanager/internal/config"
	"github.com/telepuz/smsmanager/internal/messenger/stdout"
	"github.com/telepuz/smsmanager/internal/messenger/telegram"
)

type Messenger interface {
	SendMessage(chatID int64, message internal.Message) error
}

func New(cfg *config.Config) (Messenger, error) {
	switch cfg.Messenger.Type {
	case "telegram":
		return telegram.New(cfg)
	default:
		return stdout.New()
	}
}
