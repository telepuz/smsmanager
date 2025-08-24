package stdout

import (
	"fmt"
	"log/slog"

	"github.com/telepuz/smsmanager/internal"
)

type Stdout struct{}

func New() (*Stdout, error) {
	return &Stdout{}, nil
}

func (s *Stdout) SendMessage(chatID int64, message internal.Message) error {
	slog.Info(
		fmt.Sprintf(
			"SendMessage - ChatID: %v, Phone: %s, Content: %s",
			chatID,
			message.Phone,
			message.Content,
		),
	)
	return nil
}
