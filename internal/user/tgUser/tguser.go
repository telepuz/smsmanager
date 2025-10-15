package tguser

import (
	"time"

	"github.com/telepuz/smsmanager/internal"
	"github.com/telepuz/smsmanager/internal/config"
	"github.com/telepuz/smsmanager/internal/modem"
)

type tgUser struct {
	name           string
	chatID         int64
	modem          modem.Modem
	sendSmsEnabled bool
	sendSmsTo      string
	sendSmsText    string
	sendSmsPeriod  time.Duration
}

func (t *tgUser) Name() string {
	return t.name
}

func (t *tgUser) ChatID() int64 {
	return t.chatID
}

func (t *tgUser) IsSendSms() bool {
	return t.sendSmsEnabled
}

func (t *tgUser) SendSmsPeriod() time.Duration {
	return t.sendSmsPeriod
}

func (t *tgUser) SendSmsTo() string {
	return t.sendSmsTo
}
func (t *tgUser) SendSmsText() string {
	return t.sendSmsText
}

func (t *tgUser) GetSMSMessenges() ([]internal.Message, error) {
	return t.modem.GetSMSMessenges()
}

func (t *tgUser) DeleteSMSFromModem(messageID int) error {
	return t.modem.DeleteSMSMessage(messageID)
}

func (t *tgUser) SendSms() error {
	return t.modem.SendSMS(
		t.sendSmsTo,
		t.sendSmsText,
	)
}

func New(cfg *config.Config) []*tgUser {
	users := []*tgUser{}
	for _, user := range cfg.Users {
		users = append(users, &tgUser{
			name:           user.Name,
			chatID:         user.ChatID,
			modem:          modem.New(user.ModemType, user.ModemURL),
			sendSmsEnabled: user.SendSms.Enable,
			sendSmsTo:      user.SendSms.To,
			sendSmsText:    user.SendSms.Text,
			sendSmsPeriod:  user.SendSms.Period,
		})
	}
	return users
}
