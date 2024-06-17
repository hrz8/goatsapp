package config

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	AppVersion string
	AppName    string `env:"APP_NAME,default=goatsapp"`
	LogLevel   string `env:"LOG_LEVEL,default=DEBUG"`
}

func New() *Config {
	cfg := &Config{AppVersion: Version}
	if err := envconfig.Process(context.Background(), cfg); err != nil {
		panic("Can't load configuration file")
	}

	return cfg
}
