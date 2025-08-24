package promexporter

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/telepuz/smsmanager/internal/config"
)

type PromExporter struct {
	ListenPort  int
	MetricsPath string
}

func New(cfg *config.Config) (*PromExporter, error) {
	return &PromExporter{
		ListenPort:  cfg.Exporter.ListenPort,
		MetricsPath: cfg.Exporter.MetricsPath,
	}, nil
}

func (p *PromExporter) Run() {
	r := prometheus.NewRegistry()
	r.MustRegister(
		MessageReceiveCounter,
		MessageSendCounter,
		DatabaseMessagesGauge,
		ErrMessageReceiveCounter,
		ErrMessageSendCounter,
		ErrDatabaseCounter,
	)
	handler := promhttp.HandlerFor(r, promhttp.HandlerOpts{})
	http.Handle(p.MetricsPath, handler)

	slog.Debug(
		fmt.Sprintf(
			"Starting exporter server on port %d",
			p.ListenPort,
		),
	)
	err := http.ListenAndServe(
		fmt.Sprintf(":%d", p.ListenPort),
		nil,
	)
	if err != nil {
		slog.Error(
			fmt.Sprintf(
				"Exporter http.ListenAndServe: %v",
				err,
			),
		)
	}
}
