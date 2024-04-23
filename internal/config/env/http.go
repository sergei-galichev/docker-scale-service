package env

import (
	"docker-scale-service/internal/config"
	"errors"
	"net"
	"os"
)

const (
	httpHostEnvName = "HOST"
	httpPortEnvName = "PORT"
)

var httpCfg *httpConfig

type httpConfig struct {
	host string
	port string
}

func GetHTTPConfig() (config.HTTPConfig, error) {
	if httpCfg == nil {
		c, err := newHTTPConfig()
		if err != nil {
			return nil, err
		}

		httpCfg = c

		return httpCfg, nil
	}
	return httpCfg, nil
}

func newHTTPConfig() (*httpConfig, error) {
	host := os.Getenv(httpHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("env HOST not found")
	}

	port := os.Getenv(httpPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("env PORT not found")
	}

	return &httpConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg *httpConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
