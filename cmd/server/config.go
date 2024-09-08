package main

import (
	"fmt"

	"github.com/_/_/internal/pkg/logger"
	"github.com/caarlos0/env/v11"
)

type Config struct {
	DomainName  string `env:"DOMAIN_NAME" envDefault:"localhost"`
	Host        string `env:"HOST" envDefault:"0.0.0.0"`
	Port        int    `env:"PORT" envDefault:"3000"`
	CertFile    string `env:"CERT_FILE" envDefault:""`
	CertKeyFile string `env:"CERT_KEY_FILE" envDefault:""`
	URL         string `env:"URL" envDefault:""` // This is set during load time based on the environment

	Environment string          `env:"ENVIRONMENT" envDefault:"development"`
	LogLevel    logger.LogLevel `env:"LOG_LEVEL" envDefault:"info"`
}

func LoadConfig() (cfg Config, err error) {
	cfg = Config{}
	if err = env.Parse(&cfg); err != nil {
		return
	}

	// Set the URL based on the environment
	cfg.URL = fmt.Sprintf("https://%s", cfg.DomainName)
	if cfg.Environment == "development" {
		proto := "http"
		if cfg.CertFile != "" && cfg.CertKeyFile != "" {
			proto = "https"
		}
		cfg.URL = fmt.Sprintf("%s://%s:%d", proto, cfg.DomainName, cfg.Port)
	}

	return
}
