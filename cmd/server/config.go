package main

import (
	"github.com/_/_/internal/pkg/logger"
	"github.com/caarlos0/env/v11"
)

type Config struct {
	Host        string `env:"HOST" envDefault:"0.0.0.0"`
	Port        int    `env:"PORT" envDefault:"3000"`
	CertFile    string `env:"CERT_FILE" envDefault:""`
	CertKeyFile string `env:"CERT_KEY_FILE" envDefault:""`

	Environment Environment     `env:"ENVIRONMENT" envDefault:"development"`
	LogLevel    logger.LogLevel `env:"LOG_LEVEL" envDefault:"info"`
}

type Environment string

const (
	EnvironmentDevelopment Environment = "development"
	EnvironmentProduction  Environment = "production"
)

func LoadConfig() (cfg Config, err error) {
	cfg = Config{}
	if err = env.ParseWithOptions(&cfg, env.Options{RequiredIfNoDef: true}); err != nil {
		return
	}

	return
}
