package env

import (
	"docker-scale-service/internal/config"
	"errors"
	"os"
)

const (
	loggingLevelEnvName = "LOG_LEVEL"
)

var loggingCfg *loggingConfig

type loggingConfig struct {
	level string
}

func GetLoggingConfig() (config.LoggingConfig, error) {
	if loggingCfg == nil {
		c, err := newLoggingConfig()
		if err != nil {
			return nil, err
		}

		loggingCfg = c

		return loggingCfg, nil
	}
	return loggingCfg, nil
}

func newLoggingConfig() (*loggingConfig, error) {

	level := os.Getenv(loggingLevelEnvName)
	if len(level) == 0 {
		return nil, errors.New("env LOG_LEVEL not found")
	}

	return &loggingConfig{
		level: level,
	}, nil
}

func (cfg *loggingConfig) LoggingLevel() string {
	return cfg.level
}
