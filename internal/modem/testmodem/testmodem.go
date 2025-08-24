package testmodem

import "github.com/telepuz/smsmanager/internal"

type TestModem struct{}

func New() *TestModem {
	return &TestModem{}
}

func (t *TestModem) DeleteSMSMessage(messageID int) error {
	return nil
}

func (t *TestModem) GetSMSMessenges() ([]internal.Message, error) {
	return []internal.Message{
		{
			Index:   1,
			Phone:   "+12345678901",
			Content: "This is test SMS",
			Date:    "1970-01-01 00:00",
		},
	}, nil
}
