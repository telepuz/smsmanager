package app

import (
	"fmt"
	"log/slog"
	"time"
)

func (c *AppContext) MainLoop() {
	for {
		for _, u := range c.Users {
			messenges, err := u.GetSMSMessenges()
			if err != nil {
				slog.Error(fmt.Sprintf(
					"Run(): %s",
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
						"Run(): %s",
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
						"Run(): %s",
						err,
					))
					c.Exporter.IncErrDatabaseCounter()
				}

				err = u.DeleteSMSFromModem(m.Index)
				if err != nil {
					slog.Error(fmt.Sprintf(
						"Run(): %s",
						err,
					))
				}
			}
		}

		messageCount, err := c.Storage.GetMessagesCount()
		if err != nil {
			slog.Error(fmt.Sprintf(
				"Run(): %s",
				err,
			))
			c.Exporter.IncErrDatabaseCounter()
		}
		c.Exporter.SetDatabaseMessagesGauge(messageCount)

		slog.Debug(fmt.Sprintf(
			"Run(): Sleeping for %s",
			c.Config.CheckInterval,
		))
		time.Sleep(c.Config.CheckInterval)
	}
}
