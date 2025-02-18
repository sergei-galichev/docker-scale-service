package service

import (
	"context"
	"docker-scale-service/internal/entity"
	"github.com/docker/docker/api/types/filters"
)

type DockerScaleServiceInterface interface {
	ScaleServiceReplicas(ctx context.Context, service *entity.SwarmService, scaleType int, step uint64) error
	ListServices(ctx context.Context, f filters.Args) ([]*entity.SwarmService, error)
}
