package storage

import (
	"github.com/telepuz/smsmanager/internal"
	"github.com/telepuz/smsmanager/internal/config"
	"github.com/telepuz/smsmanager/internal/storage/sqlite"
	"github.com/telepuz/smsmanager/internal/storage/stdout"
)

type Storage interface {
	SaveMessage(message internal.Message, chatID int64) error
	GetMessagesCount() (int, error)
	DatabaseClose() error
}

func New(cfg *config.Config) (Storage, error) {
	switch cfg.Storage.Type {
	case "sqlite3":
		return sqlite.New(cfg)
	default:
		return stdout.New(), nil
	}
}
