package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	CheckInterval time.Duration `yaml:"check_interval" env-default:"1m"`
	Logger        Logger        `yaml:"logger"`
	Messenger     Messenger     `yaml:"messenger"`
	Storage       Storage       `yaml:"storage"`
	Users         []User        `yaml:"users"`
	Exporter      Exporter      `yaml:"exporter"`
	HealthCheck   HealthCheck   `yaml:"health_check"`
}

type User struct {
	Name      string  `yaml:"name"`
	ChatID    int64   `yaml:"chat_id"`
	ModemURL  string  `yaml:"modem_url"`
	ModemType string  `yaml:"modem_type" env-default:"huaweie3372"`
	SendSms   SendSms `yaml:"send_sms"`
}

type Logger struct {
	Format string `yaml:"format" env-default:"plaintext"`
	Level  string `yaml:"level" env-default:"info"`
}

type Messenger struct {
	Type  string `yaml:"type" env-default:"telegram"`
	Token string `yaml:"token"`
}

type Storage struct {
	Type       string `yaml:"type" env-default:"sqlite3"`
	DBFilePath string `yaml:"file_path" env-default:"/var/lib/smsmanager/db.sql"`
}

type Exporter struct {
	Type        string `yaml:"type" env-default:"none"`
	ListenPort  int    `yaml:"listen_port"`
	MetricsPath string `yaml:"metrics_path"`
}

type HealthCheck struct {
	Enable     bool `yaml:"enable" env-default:"false"`
	ListenPort int  `yaml:"listen_port" env-default:"3000"`
}

type SendSms struct {
	Enable bool          `yaml:"enable" env-default:"false"`
	Period time.Duration `yaml:"period" env-default:"60d"`
	To     string        `yaml:"to"`
	Text   string        `yalml:"text"`
}

func New(configFile string) (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadConfig(configFile, cfg)
	if err != nil {
		return nil, fmt.Errorf("NewConfig(): config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
