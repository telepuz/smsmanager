package stdout

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/telepuz/smsmanager/internal"
)

type Stdout struct {
}

func New() *Stdout {
	return &Stdout{}
}

func (s *Stdout) GetMessagesCount() (int, error) {
	return 0, nil
}

func (s *Stdout) DatabaseClose() error {
	return nil
}

func (s *Stdout) IsItTimeToSendSms(username string, dt time.Duration) (bool, error) {
	return true, nil
}

func (s *Stdout) SaveSendSmsTime(phone_number, username, text string) error {
	return nil
}

func (s *Stdout) SaveMessage(message internal.Message, chatID int64) error {
	slog.Debug(
		fmt.Sprintf(
			"SaveMessage - ChatID: %v, Message: %v",
			chatID,
			message,
		),
	)
	return nil
}
