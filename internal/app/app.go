package app

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/telepuz/smsmanager/internal/config"
	"github.com/telepuz/smsmanager/internal/exporter"
	"github.com/telepuz/smsmanager/internal/messenger"
	"github.com/telepuz/smsmanager/internal/probes"
	"github.com/telepuz/smsmanager/internal/storage"
	"github.com/telepuz/smsmanager/internal/user"
)

type AppContext struct {
	Config      *config.Config
	Storage     storage.Storage
	Users       []user.User
	Messenger   messenger.Messenger
	Exporter    exporter.Exporter
	HealthCheck *probes.ProbeService
}

func (c *AppContext) Run() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ch
		slog.Info("Run(): Got SIGTTERM")
		slog.Info("Run(): Exit...")
		os.Exit(0)
	}()

	go c.HealthCheck.Run()
	go c.Exporter.Run()

	c.HealthCheck.SetReady()
	c.MainLoop()
}
