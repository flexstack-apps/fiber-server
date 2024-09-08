package main

import (
	"github.com/_/_/internal/pkg/logger"
	"github.com/caarlos0/env/v11"
)

type Config struct {
	Port        string          `env:"PORT" envDefault:"9000"`
	LogLevel    logger.LogLevel `env:"LOG_LEVEL" envDefault:"info"`
	Environment string          `env:"ENVIRONMENT" envDefault:"development"`
}

func LoadConfig() (cfg Config, err error) {
	cfg = Config{}
	err = env.Parse(&cfg)
	return
}
