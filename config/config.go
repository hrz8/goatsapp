package config

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	AppVersion       string
	AppVersionNumber [3]uint32
	AppPort          string `env:"APP_PORT,default=1543"`
	AppName          string `env:"APP_NAME,default=Goatsapp"`
	LogLevel         string `env:"LOG_LEVEL,default=DEBUG"`
	DatabaseURL      string `env:"DATABASE_URL"`
}

func New() *Config {
	cfg := &Config{
		AppVersion:       Version,
		AppVersionNumber: VersionNumber,
	}
	if err := envconfig.Process(context.Background(), cfg); err != nil {
		panic("can't load configuration file")
	}

	return cfg
}
