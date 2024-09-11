package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/_/_/internal/app/hello"
	"github.com/_/_/internal/pkg/logger"
	"github.com/_/_/internal/pkg/mw"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/favicon"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	fiberrecover "github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"go.uber.org/automaxprocs/maxprocs"
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
		Pretty:   cfg.Environment == EnvironmentDevelopment,
	})

	app := fiber.New(fiber.Config{
		StrictRouting:      true,
		EnableIPValidation: true,
	})
	app.Use(mw.NewRealIP())
	app.Use(helmet.New(helmet.Config{HSTSPreloadEnabled: true, HSTSMaxAge: 31536000}))
	app.Use(fiberrecover.New(fiberrecover.Config{EnableStackTrace: cfg.Environment == EnvironmentDevelopment}))
	app.Use(favicon.New())
	app.Use(requestid.New())
	serverLog := log.With("source", "server")
	app.Use(mw.NewLogger(serverLog, slog.LevelInfo))

	clientSvc := hello.New(hello.Options{Logger: log.With("source", "client")})
	app.Get("/", func(c fiber.Ctx) error {
		hello, err := clientSvc.Hello(mw.GetRealIP(c))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
		}

		return c.JSON(hello)
	})
	app.Get(mw.HealthCheckEndpoint, healthcheck.NewHealthChecker())

	g := errgroup.Group{}

	g.Go(func() error {
		addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
		listenConfig := fiber.ListenConfig{
			GracefulContext:       ctx,
			ListenerNetwork:       fiber.NetworkTCP,
			DisableStartupMessage: true,
			CertFile:              cfg.CertFile,
			CertKeyFile:           cfg.CertKeyFile,
			OnShutdownError: func(err error) {
				serverLog.Error("error shutting down", "error", err)
			},
			OnShutdownSuccess: func() {
				serverLog.Info("shutdown successfully")
			},
		}

		if cfg.Environment == EnvironmentDevelopment {
			proto := "http"
			if cfg.CertFile != "" && cfg.CertKeyFile != "" {
				proto = "https"
			}

			fmt.Println(startupLogo + fmt.Sprintf("%s://localhost:%d", proto, cfg.Port))
		} else {
			serverLog.Info("starting server", "address", addr, "environment", cfg.Environment)
		}

		if err := app.Listen(addr, listenConfig); err != nil {
			return err
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		log.Error("error", "error", err)
		os.Exit(1)
	}
}

func init() {
	maxprocs.Set()
}

var startupLogo = `
    _______ __
   / ____(_) /_  ___  _____
  / /_  / / __ \/ _ \/ ___/
 / __/ / / /_/ /  __/ /    
/_/   /_/_.___/\___/_/       

ꕤ URL  `
