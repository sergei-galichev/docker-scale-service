package config

import (
	"context"
	"flag"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"os"
)

func InitConfig(_ context.Context) error {
	mode := flag.String("mode", "dev", "run mode")

	flag.Parse()

	switch *mode {
	case "dev":
		err := godotenv.Load("./configs/dev.env")
		if err != nil {
			return errors.New("Error loading dev.env file")
		}

		err = os.Setenv("APP_MODE", "dev")
		if err != nil {
			return errors.New("Error setting APP_MODE environment variable")
		}

	case "local":
		err := os.Setenv("APP_MODE", "local")
		if err != nil {
			return errors.New("Error setting APP_MODE environment variable")
		}

	case "prod":
		err := os.Setenv("APP_MODE", "prod")
		if err != nil {
			return errors.New("Error setting APP_MODE environment variable")
		}

	default:
		return errors.New("Invalid mode")
	}

	return nil
}
