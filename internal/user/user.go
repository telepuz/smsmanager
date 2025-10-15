package user

import (
	"time"

	"github.com/telepuz/smsmanager/internal"
	"github.com/telepuz/smsmanager/internal/config"
	tguser "github.com/telepuz/smsmanager/internal/user/tgUser"
)

type User interface {
	GetSMSMessenges() ([]internal.Message, error)
	DeleteSMSFromModem(messageID int) error
	ChatID() int64
	Name() string
	IsSendSms() bool
	SendSmsTo() string
	SendSmsText() string
	SendSmsPeriod() time.Duration
	SendSms() error
}

func New(cfg *config.Config) ([]User, error) {
	users := []User{}
	for _, user := range tguser.New(cfg) {
		users = append(users, user)
	}
	return users, nil
}
