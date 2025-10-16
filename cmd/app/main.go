package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/telepuz/smsmanager/internal/app"
	"github.com/telepuz/smsmanager/internal/config"
	"github.com/telepuz/smsmanager/internal/exporter"
	"github.com/telepuz/smsmanager/internal/logger"
	"github.com/telepuz/smsmanager/internal/messenger"
	"github.com/telepuz/smsmanager/internal/probes"
	"github.com/telepuz/smsmanager/internal/storage"
	"github.com/telepuz/smsmanager/internal/user"
)

var (
	configFile  *string
	versionFlag *bool
	version     = "devel"
	gitRevision = "devel"
	buildDate   = "devel"
)

func init() {
	configFile = flag.String("config_file", "/etc/smsmanager/smsmanager.yml", "Config filename")
	versionFlag = flag.Bool("version", false, "Print version")
}

func main() {
	flag.Parse()

	if *versionFlag {
		fmt.Printf(
			"Name: %s\nVersion: %s\ngitRevision: %s\nbuildDate: %s\n",
			"SMSmanager",
			version,
			gitRevision,
			buildDate)
		os.Exit(0)
	}

	cfg, err := config.New(*configFile)
	if err != nil {
		slog.Error(
			fmt.Sprintf("main(): %s", err))
		os.Exit(1)
	}
	slog.Debug("main(): Read config file")

	err = logger.ConfigureSlog(&cfg.Logger)
	if err != nil {
		slog.Error(
			fmt.Sprintf("main(): %s", err))
		os.Exit(1)
	}
	slog.Debug("main(): Configured slog")
	slog.Debug(fmt.Sprintf("main(): Read configs: %+v", cfg))

	storage, err := storage.New(cfg)
	if err != nil {
		slog.Error(
			fmt.Sprintf("main(): %s", err))
		os.Exit(1)
	}
	defer func() {
		err = storage.DatabaseClose()
		if err != nil {
			slog.Error(
				fmt.Sprintf("main(): %s", err))
			os.Exit(1)
		}
	}()

	messenger, err := messenger.New(cfg)
	if err != nil {
		slog.Error(
			fmt.Sprintf("main(): %s", err))
		os.Exit(1)
	}

	users, err := user.New(cfg)
	if err != nil {
		slog.Error(
			fmt.Sprintf("main(): %s", err))
		os.Exit(1)
	}

	exporter, err := exporter.New(cfg)
	if err != nil {
		slog.Error(
			fmt.Sprintf("main(): %s", err))
		os.Exit(1)
	}

	hc := probes.New(cfg)

	app := app.AppContext{
		Config:      cfg,
		Storage:     storage,
		Messenger:   messenger,
		Users:       users,
		Exporter:    exporter,
		HealthCheck: hc,
	}

	slog.Info("main(): Starting app...")
	app.Run()
	slog.Info("main(): Exit")
}
