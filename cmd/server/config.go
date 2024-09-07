package main

import (
	"app/internal/environment"
	"app/internal/logger"
	"github.com/caarlos0/env/v11"
)

type Config struct {
	Port        string                  `env:"PORT" envDefault:"4000"`
	LogLevel    logger.LogLevel         `env:"LOG_LEVEL" envDefault:"info"`
	Environment environment.Environment `env:"ENVIRONMENT" envDefault:"development"`
}

func LoadConfig() (cfg Config, err error) {
	cfg = Config{}
	err = env.Parse(&cfg)
	return
}
