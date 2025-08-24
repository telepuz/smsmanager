package tguser

import (
	"github.com/telepuz/smsmanager/internal"
	"github.com/telepuz/smsmanager/internal/config"
	"github.com/telepuz/smsmanager/internal/modem"
)

type tgUser struct {
	name   string
	chatID int64
	modem  modem.Modem
}

func (t *tgUser) Name() string {
	return t.name
}

func (t *tgUser) ChatID() int64 {
	return t.chatID
}

func (t *tgUser) GetSMSMessenges() ([]internal.Message, error) {
	return t.modem.GetSMSMessenges()
}

func (t *tgUser) DeleteSMSFromModem(messageID int) error {
	return t.modem.DeleteSMSMessage(messageID)
}

func New(cfg *config.Config) []*tgUser {
	users := []*tgUser{}
	for _, user := range cfg.Users {
		users = append(users, &tgUser{
			name:   user.Name,
			chatID: user.ChatID,
			modem:  modem.New(user.ModemType, user.ModemURL),
		})
	}
	return users
}
