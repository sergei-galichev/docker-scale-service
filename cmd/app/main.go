package main

import (
	"context"
	"docker-scale-service/internal/app"
	"docker-scale-service/pkg/logging"
	"os"
	"os/signal"
	"syscall"
)

var (
	ctx     = context.Background()
	signals = []os.Signal{
		syscall.SIGABRT,
		syscall.SIGQUIT,
		syscall.SIGHUP,
		os.Interrupt,
		syscall.SIGTERM,
	}
)

func main() {
	application, err := app.NewApp(ctx)
	if err != nil {
		panic(err)
	}

	err = application.Run()
	if err != nil {
		panic(err)
	}

	shutdown(signals)
}

func shutdown(signals []os.Signal) {
	l := logging.GetLogger()

	ch := make(chan os.Signal, 1)

	signal.Notify(ch, signals...)
	sig := <-ch

	l.Infof("Caught signal: %s. Shutting down...", sig)
}
