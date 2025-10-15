package modem

import (
	"github.com/telepuz/smsmanager/internal"
	"github.com/telepuz/smsmanager/internal/modem/huaweie3372"
	"github.com/telepuz/smsmanager/internal/modem/testmodem"
)

type Modem interface {
	GetSMSMessenges() ([]internal.Message, error)
	SendSMS(phoneNumber, text string) error
	DeleteSMSMessage(messageID int) error
}

func New(modemType, modemURL string) Modem {
	switch modemType {
	case "test-modem":
		return testmodem.New()
	default:
		return huaweie3372.New(modemURL)
	}
}
