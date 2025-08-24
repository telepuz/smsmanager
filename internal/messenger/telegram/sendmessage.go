package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/telepuz/smsmanager/internal"
)

const message_template = `
*%s*

%s
`

type message struct {
	ChatID              int64  `json:"chat_id"`
	ParseMode           string `json:"parse_mode"`
	DisablePreview      bool   `json:"disable_web_page_preview"`
	DisableNotification bool   `json:"disable_notification"`
	Text                string `json:"text"`
}

func newMessage(chatID int64, title, body string) *message {
	slog.Debug(fmt.Sprintf(
		"NewMessage(): Created new telegram-message: chatID: %v, title: %s, Body: %s",
		chatID,
		title,
		body,
	))
	return &message{
		ChatID:              chatID,
		ParseMode:           "Markdown",
		DisablePreview:      true,
		DisableNotification: false,
		Text: fmt.Sprintf(
			message_template,
			title,
			body,
		),
	}
}

func (t *Telegram) SendMessage(chatID int64, message internal.Message) error {
	m := newMessage(
		chatID,
		message.Phone,
		message.Content,
	)
	payload, err := json.Marshal(m)
	if err != nil {
		return fmt.Errorf(
			"SendMessage: %v",
			err,
		)
	}
	response, err := http.Post(
		fmt.Sprintf(
			"https://api.telegram.org/bot%s/sendMessage",
			t.token,
		),
		"application/json",
		bytes.NewBuffer(payload),
	)
	if err != nil {
		return fmt.Errorf(
			"SendMessage: %v",
			err,
		)
	}
	slog.Debug(fmt.Sprintf(
		"SendMessage: Request complete: Code - %d",
		response.StatusCode,
	))
	defer func(body io.ReadCloser) {
		if err := body.Close(); err != nil {
			slog.Warn("failed to close response body")
		}
	}(response.Body)
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("SendMessage: failed to send request. Status: %q", response.Status)
	}
	slog.Info(fmt.Sprintf(
		"SendMessage: Phone: %s, Content: %s",
		message.Phone,
		message.Content,
	))
	return nil
}
