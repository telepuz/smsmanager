package noneexporter

import (
	"log/slog"
)

type NoneExporter struct{}

func New() (*NoneExporter, error) {
	slog.Debug("Use None Exporter")
	return &NoneExporter{}, nil
}

func (n *NoneExporter) Run()                           {}
func (n *NoneExporter) IncMessageReceiveCounter()      {}
func (n *NoneExporter) IncMessageSendCounter()         {}
func (n *NoneExporter) SetDatabaseMessagesGauge(k int) {}
func (n *NoneExporter) IncErrMessageReceiveCounter()   {}
func (n *NoneExporter) IncErrMessageSendCounter()      {}
func (n *NoneExporter) IncErrDatabaseCounter()         {}
