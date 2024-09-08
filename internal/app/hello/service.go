package hello

import (
	"log/slog"
)

type Options struct {
	Logger *slog.Logger
}

type Service struct {
	logger *slog.Logger
}

func New(opt Options) *Service {
	return &Service{logger: opt.Logger}
}

func (s *Service) Hello(ip string) (HelloResult, error) {
	log := s.logger.With("method", "Hello")
	log.Debug("called with", "ip", ip)
	return HelloResult{ClientAddress: ip}, nil
}

type HelloResult struct {
	ClientAddress string `json:"client_address"`
}
