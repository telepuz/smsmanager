package logger

import (
	"log/slog"
	"os"

	"github.com/telepuz/smsmanager/internal/config"
)

func ConfigureSlog(c *config.Logger) error {
	level := slog.LevelInfo
	switch c.Level {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		slog.Warn("ConfigureSlog: Unknown log level")
	}
	opts := &slog.HandlerOptions{
		Level: level,
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	switch c.Format {
	case "plaintext":
		logger = slog.New(slog.NewTextHandler(os.Stdout, opts))
	case "json":
		logger = slog.New(slog.NewJSONHandler(os.Stdout, opts))
	default:
		slog.Warn("ConfigureSlog: Unknown log format")
	}

	slog.SetDefault(logger)
	return nil
}
