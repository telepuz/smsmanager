package testmodem

import (
	"fmt"
	"log/slog"

	"github.com/telepuz/smsmanager/internal"
)

type TestModem struct{}

func New() *TestModem {
	return &TestModem{}
}

func (t *TestModem) DeleteSMSMessage(messageID int) error {
	return nil
}

func (t *TestModem) SendSMS(phoneNumber, text string) error {
	slog.Info(fmt.Sprintf(
		"Send SMS to %s, with text: %s",
		phoneNumber,
		text))
	return nil
}

func (t *TestModem) GetSMSMessenges() ([]internal.Message, error) {
	return []internal.Message{
		{
			Index: 1,
			Phone: "+12345678901",
			// TODO: Add escapeSting to tests
			// Content: "'_', '*', '[', ']', '(', ')', '~', '`', '>', '#', '+', '-', '=', '|', '{', '}', '.', '!'",
			Content: "This is test SMS",
			Date:    "1970-01-01 00:00",
		},
	}, nil
}
