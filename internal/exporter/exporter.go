package exporter

import (
	"github.com/telepuz/smsmanager/internal/config"
	"github.com/telepuz/smsmanager/internal/exporter/noneexporter"
	"github.com/telepuz/smsmanager/internal/exporter/promexporter"
)

type Exporter interface {
	Run()
	IncMessageReceiveCounter()
	IncMessageSendCounter()
	SetDatabaseMessagesGauge(n int)
	IncErrMessageReceiveCounter()
	IncErrMessageSendCounter()
	IncErrDatabaseCounter()
}

func New(cfg *config.Config) (Exporter, error) {
	switch cfg.Exporter.Type {
	case "prom":
		return promexporter.New(cfg)
	default:
		return noneexporter.New()
	}
}
