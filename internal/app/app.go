package app

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		sig := <-ch
		slog.Info(fmt.Sprintf("Run(): Received signal: %v", sig))
		slog.Info("Run(): Shutdown signal received. Waiting for graceful stop...")

		cancel()
		time.Sleep(3 * time.Second)
		os.Exit(0)
	}()

	go c.HealthCheck.Run()
	go c.Exporter.Run()

	c.HealthCheck.SetReady()
	c.MainLoop(ctx)
}
