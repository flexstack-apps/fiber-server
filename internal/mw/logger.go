package mw

import (
	"app/internal/logger"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func NewLogger(logger *logger.Logger) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if c.Path() == "/healthz" {
			return c.Next()
		}

		c.Locals(LoggerKey, logger)
		if err := c.Next(); err != nil {
			return err
		}

		logger.Info(
			fmt.Sprintf("%s %s", c.Method(), c.Path()),
			"status", c.Response().StatusCode(),
			"ip", GetRealIP(c),
			"duration", time.Since(c.Context().Time()).String(),
		)
		return nil
	}
}

func GetLogger(c *fiber.Ctx) *logger.Logger {
	return c.Locals(LoggerKey).(*logger.Logger)
}

const (
	// LoggerKey is the key used to store the logger in the context
	LoggerKey = "logger"
)
