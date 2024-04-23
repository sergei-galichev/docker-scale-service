package service

import (
	"context"
	"docker-scale-service/internal/entity"
)

type DockerScaleServiceInterface interface {
	ScaleServiceReplicas(ctx context.Context, service *entity.SwarmService, scaleType int, step uint64) error
}
