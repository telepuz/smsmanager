package probes

import (
	"fmt"
	"log/slog"
	"net/http"
	"sync/atomic"

	"github.com/telepuz/smsmanager/internal/config"
)

type ProbeService struct {
	isReady     atomic.Bool
	listen_port int
	enable      bool
}

func New(cfg *config.Config) *ProbeService {
	return &ProbeService{
		listen_port: cfg.HealthCheck.ListenPort,
		enable:      cfg.HealthCheck.Enable,
	}
}

func (p *ProbeService) SetReady() {
	p.isReady.Store(true)
}

func (p *ProbeService) SetUnReady() {
	p.isReady.Store(false)
}

func (p *ProbeService) Run() {
	if p.enable {
		http.HandleFunc("/livez", p.livenessHandler)
		http.HandleFunc("/readyz", p.readinessHandler)

		slog.Debug(
			fmt.Sprintf(
				"Starting health-check server on :%v",
				p.listen_port,
			),
		)

		http.ListenAndServe(
			fmt.Sprintf(":%v", p.listen_port),
			nil,
		)
	}
}

func (p *ProbeService) livenessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Liveness OK\n")
}

func (p *ProbeService) readinessHandler(w http.ResponseWriter, r *http.Request) {
	if p.isReady.Load() {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Readiness OK\n")
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, "Readiness NOT OK: Initializing...\n")
	}
}
