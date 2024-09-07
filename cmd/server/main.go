package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"
	"time"

	"app/internal/environment"
	"app/internal/logger"
	"github.com/gofiber/fiber/v2"
	fiberrecover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/segmentio/encoding/json"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx := context.Background()
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config %v", err)
	}

	log := logger.New(logger.Options{
		LogLevel: cfg.LogLevel,
		Pretty:   cfg.Environment == environment.Development,
	}).With("service", "receiver")

	g := errgroup.Group{}
	srv := fiber.New(fiber.Config{
		StrictRouting: true,
		Network:       "tcp",
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
	})

	srv.Use(fiberrecover.New(fiberrecover.Config{EnableStackTrace: cfg.Environment == environment.Development}))
	// if cfg.Environment == environment.Development {
	// 	srv.Use(flogger.New(flogger.Config{}))
	// }
	srv.Get("/", func(c *fiber.Ctx) error {
		c.Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
		c.Set("Pragma", "no-cache")
		return c.Send([]byte("{\"ok\": true}"))
	})
	srv.Get("/health", func(c *fiber.Ctx) error {
		c.Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
		c.Set("Pragma", "no-cache")
		c.Set("Content-Type", "text/plain")
		return c.SendString(".")
	})

	srvShutdown := errgroup.Group{}
	srvShutdown.Go(func() error {
		<-ctx.Done()
		log.Info("shutting down")
		if err := srv.ShutdownWithTimeout(time.Second * 5); err != nil {
			log.Error(fmt.Sprintf("shutdown: %s\n", err))
			return err
		}
		log.Info("successfully shut down")
		return nil
	})

	g.Go(func() error {
		if err := srv.Listen(":" + cfg.Port); err != nil {
			return err
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		log.Error("error", "error", err)
		stop()
	}

	if err := srvShutdown.Wait(); err != nil {
		log.Error(fmt.Sprintf("shutdown: %s\n", err))
	}

	log.Info("exiting")
}
