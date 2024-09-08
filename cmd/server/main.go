package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/_/_/internal/app/hello"
	"github.com/_/_/internal/pkg/logger"
	"github.com/_/_/internal/pkg/mw"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	fiberrecover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx := context.Background()
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg, err := LoadConfig()
	if err != nil {
		panic(err)
	}

	log := logger.New(logger.Options{
		LogLevel: cfg.LogLevel,
		Pretty:   cfg.Environment == "development",
	})

	app := fiber.New(fiber.Config{
		StrictRouting:      true,
		Network:            "tcp",
		EnableIPValidation: true,
	})
	app.Use(mw.NewRealIP())
	app.Use(helmet.New(helmet.Config{HSTSPreloadEnabled: true, HSTSMaxAge: 31536000}))
	app.Use(fiberrecover.New(fiberrecover.Config{EnableStackTrace: cfg.Environment == "devleopment"}))
	app.Use(favicon.New())
	app.Use(requestid.New())
	app.Use(healthcheck.New(healthcheck.Config{LivenessEndpoint: "/healthz"}))
	app.Use(mw.NewLogger(log.With("source", "server")))

	clientSvc := hello.New(hello.Options{Logger: log.With("source", "client")})
	app.Get("/", func(c *fiber.Ctx) error {
		hello, err := clientSvc.Hello(mw.GetRealIP(c))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
		}

		return c.JSON(hello)
	})

	g := errgroup.Group{}

	g.Go(func() error {
		if err := app.Listen(":" + cfg.Port); err != nil {
			return err
		}

		return nil
	})

	sg := errgroup.Group{}
	sg.Go(func() error {
		<-ctx.Done()

		if err := app.ShutdownWithTimeout(time.Second * 5); err != nil {
			log.Error(fmt.Sprintf("shutdown: %s\n", err))
			return err
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		log.Error("error", "error", err)
		os.Exit(1)
	}

	if err := sg.Wait(); err != nil {
		log.Error(fmt.Sprintf("shutdown: %s\n", err))
	}

	log.Info("exiting")
}
