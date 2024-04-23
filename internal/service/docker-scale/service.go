package docker_scale

import (
	services "docker-scale-service/internal/service"
	"docker-scale-service/pkg/logging"
	"fmt"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
	"time"
)

var (
	_ services.DockerScaleServiceInterface = (*service)(nil)
)

const (
	clientCreateError = "Service: %s"
)

type service struct {
	dockerClient *client.Client

	logger *logging.Logger
}

// NewService creates new Docker Scale Service instance
func NewService() (*service, error) {
	logger := logging.GetLogger()

	c, err := client.NewClientWithOpts(
		client.WithHost("unix:///var/run/docker.sock"),
		client.WithVersion(""),
		client.WithTimeout(time.Second*15),
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		return nil, errors.New(fmt.Sprintf(clientCreateError, err.Error()))
	}

	return &service{
		dockerClient: c,
		logger:       logger,
	}, nil
}
