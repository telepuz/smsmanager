package app

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/telepuz/smsmanager/internal"
)

func sleepContext(ctx context.Context, delay time.Duration) {
	select {
	case <-ctx.Done():
	case <-time.After(delay):
	}
}

func (c *AppContext) MainLoop(ctx context.Context) {
	slog.Debug("Starting MainLoop...")
	defer slog.Debug("Stopping MainLoop...")

	for {
		select {
		case <-ctx.Done():
			return
		default:
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
			sleepContext(ctx, c.Config.CheckInterval)
		}
	}
}
