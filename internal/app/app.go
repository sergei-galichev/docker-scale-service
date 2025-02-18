package app

import (
	"context"
	"docker-scale-service/internal/config"
	"docker-scale-service/internal/handler/http_v1"
	"docker-scale-service/internal/usecase"
	"docker-scale-service/pkg/logging"
	"github.com/pkg/errors"
)

type App struct {
	logger *logging.Logger

	httpServer http_v1.ServerHttp

	uc *usecase.DockerScaleUseCase
}

// NewApp creates new App instance
func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return a, nil
}

// Run runs application
func (app *App) Run() error {
	err := app.runHTTPServer()
	if err != nil {
		app.logger.Error("Failed to run HTTP server: ", err)
		return errors.WithStack(err)
	}

	return nil
}

// initDeps initializes dependencies of the application
func (app *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		app.initConfig,
		app.initLogger,
		app.initHTTPServer,
	}

	for _, fn := range inits {
		if err := fn(ctx); err != nil {
			return err
		}
	}

	return nil
}

// initConfig initializes config of the application
func (app *App) initConfig(ctx context.Context) error {
	err := config.InitConfig(ctx)
	if err != nil {
		return err
	}

	return nil
}

// initLogger initializes logger of the application
func (app *App) initLogger(ctx context.Context) error {
	err := logging.Init(ctx)
	if err != nil {
		return err
	}

	app.logger = logging.GetLogger()

	return nil
}

// initUseCase initializes use case of the application
func (app *App) initUseCase(ctx context.Context) error {
	app.uc = usecase.NewUseCase(ctx)
	return nil
}

// initHTTPServer initializes HTTP server
func (app *App) initHTTPServer(ctx context.Context) error {
	app.httpServer = http_v1.NewHttpServer(ctx, app.uc)
	return nil
}

// runHTTPServer runs HTTP server
func (app *App) runHTTPServer() error {
	return app.httpServer.Run()
}

// stopHTTPServer stops HTTP server
func (app *App) stopHTTPServer() error {
	return app.httpServer.Shutdown()
}
