package app

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/telepuz/smsmanager/internal"
)

func (c *AppContext) MainLoop() {
	slog.Debug("Starting MainLoop...")
	for {
		for _, u := range c.Users {
			messenges, err := u.GetSMSMessenges()
			if err != nil {
				slog.Error(fmt.Sprintf(
					"MainLoop(): %s",
					err,
				))
				c.Exporter.IncErrMessageReceiveCounter()
			}

			chatID := u.ChatID()
			for _, m := range messenges {
				c.Exporter.IncMessageReceiveCounter()

				err = c.Messenger.SendMessage(
					chatID,
					m,
				)
				if err != nil {
					slog.Error(fmt.Sprintf(
						"MainLoop(): %s",
						err,
					))
					c.Exporter.IncErrMessageSendCounter()
				}
				c.Exporter.IncMessageSendCounter()

				err = c.Storage.SaveMessage(
					m,
					chatID,
				)
				if err != nil {
					slog.Error(fmt.Sprintf(
						"MainLoop(): %s",
						err,
					))
					c.Exporter.IncErrDatabaseCounter()
				}

				err = u.DeleteSMSFromModem(m.Index)
				if err != nil {
					slog.Error(fmt.Sprintf(
						"MainLoop(): %s",
						err,
					))
				}
			}

			if u.IsSendSms() {
				ok, err := c.Storage.IsItTimeToSendSms(
					u.Name(),
					u.SendSmsPeriod(),
				)
				if err != nil {
					slog.Error(fmt.Sprintf(
						"MainLoop(): %s",
						err,
					))
					c.Exporter.IncErrDatabaseCounter()
				}
				if ok {
					err = u.SendSms()
					if err != nil {
						slog.Error(fmt.Sprintf(
							"MainLoop(): %s",
							err,
						))
						c.Exporter.IncErrSendSMSCounter()
					}

					err = c.Storage.SaveSendSmsTime(
						u.SendSmsTo(),
						u.Name(),
						u.SendSmsText(),
					)
					if err != nil {
						slog.Error(fmt.Sprintf(
							"MainLoop(): %s",
							err,
						))
						c.Exporter.IncErrDatabaseCounter()
					}

					err = c.Messenger.SendMessage(
						u.ChatID(),
						internal.Message{
							Index: 1,
							Phone: "System",
							Content: fmt.Sprintf(
								"SMS was sent to number %s",
								u.SendSmsTo(),
							),
							Date: "1970-01-01 00:00",
						},
					)
					if err != nil {
						slog.Error(fmt.Sprintf(
							"MainLoop(): %s",
							err,
						))
						c.Exporter.IncErrMessageSendCounter()
					}
					c.Exporter.IncMessageSendCounter()
				}
			}
		}

		messageCount, err := c.Storage.GetMessagesCount()
		if err != nil {
			slog.Error(fmt.Sprintf(
				"MainLoop(): %s",
				err,
			))
			c.Exporter.IncErrDatabaseCounter()
		}
		c.Exporter.SetDatabaseMessagesGauge(messageCount)

		slog.Debug(fmt.Sprintf(
			"MainLoop(): Sleeping for %s",
			c.Config.CheckInterval,
		))
		time.Sleep(c.Config.CheckInterval)
	}
}
