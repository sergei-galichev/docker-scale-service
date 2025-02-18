package usecase

import (
	"context"
	services "docker-scale-service/internal/service"
	dockerScale "docker-scale-service/internal/service/docker-scale"
)

type DockerScaleUseCase struct {
	dockerScaleService services.DockerScaleServiceInterface
}

func NewUseCase(ctx context.Context) *DockerScaleUseCase {
	return &DockerScaleUseCase{}
}

func (uc *DockerScaleUseCase) DockerScaleService() services.DockerScaleServiceInterface {
	if uc.dockerScaleService == nil {
		var err error

		uc.dockerScaleService, err = dockerScale.NewService()
		if err != nil {
			panic(err)
		}
	}
	return uc.dockerScaleService
}
