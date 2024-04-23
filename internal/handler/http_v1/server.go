package http_v1

import (
	"context"
	"docker-scale-service/internal/config"
	"docker-scale-service/internal/config/env"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type server struct {
	app     *fiber.App
	httpCfg config.HTTPConfig
}

func NewHttpServer(_ context.Context) *server {
	app := fiber.New(
		fiber.Config{
			AppName: "Docker Swarm Auto-Scale Service App",
		},
	)

	httpCfg, err := env.GetHTTPConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to get HTTP config: %s", err.Error()))
	}

	loggingCfg, err := env.GetLoggingConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to get HTTP config: %s", err.Error()))
	}

	level := logLevel(loggingCfg.LoggingLevel())
	log.SetLevel(level)

	app.Use(
		recover.New(),
		logger.New(
			logger.Config{
				Format:     "[${time}] [${ip}:${port}] - ${status} - ${latency} - ${method} ${path}\n",
				TimeFormat: "2006/01/02 - 15:04:05",
			},
		),
		cors.New(
			cors.Config{
				AllowHeaders: fiber.HeaderAuthorization,
				AllowOrigins: "*",
			},
		),
		pprof.New(
			pprof.Config{
				Prefix: "/profiler",
			},
		),
	)

	s := &server{
		app:     app,
		httpCfg: httpCfg,
	}

	s.initRoutes()

	return s
}

func logLevel(level string) log.Level {
	switch level {
	case "debug":
		return log.LevelDebug
	case "info":
		return log.LevelInfo
	case "warn":
		return log.LevelWarn
	case "error":
		return log.LevelError
	default:
		return log.LevelTrace
	}
}

func (s *server) initRoutes() {
	main := s.app.Group("/")
	main.Get("/", s.Home)
	main.Get("/ping", s.Ping)

}

// Run is used to start the server
func (s *server) Run() error {
	addr := s.httpCfg.Address()
	go func() {
		err := s.app.Listen(addr)
		if err != nil {
			log.Fatal(err)
		}
	}()

	return nil
}

// Shutdown is used to stop the server
func (s *server) Shutdown() error {
	return s.app.Shutdown()
}
