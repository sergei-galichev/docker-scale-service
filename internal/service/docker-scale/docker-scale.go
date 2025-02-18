package docker_scale

import (
	"context"
	"docker-scale-service/internal/entity"
	"docker-scale-service/pkg/constant"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/swarm"
	"github.com/pkg/errors"
)

const (
	invalidScaleTypeError     = "Service: Invalid scale type: %d"
	failedToScaleServiceError = "Service: [ServiceUpdate] %s"
)

// ScaleServiceReplicas scales service replicas by step in up or down direction.
// It uses ServiceUpdate method of Docker API.
func (s *service) ScaleServiceReplicas(
	ctx context.Context,
	srv *entity.SwarmService,
	scaleType int,
	step uint64,
) error {
	replicas := *srv.Service.Spec.Mode.Replicated.Replicas

	if scaleType == constant.Up {
		*srv.Service.Spec.Mode.Replicated.Replicas = replicas + step
	} else if scaleType == constant.Down {
		*srv.Service.Spec.Mode.Replicated.Replicas = replicas - step
	} else {
		e := errors.New(fmt.Sprintf(invalidScaleTypeError, scaleType))
		s.logger.Error(e)
		return e
	}

	res, err := s.dockerClient.ServiceUpdate(
		ctx,
		srv.Service.ID,
		swarm.Version{
			Index: srv.Service.Meta.Version.Index,
		},
		srv.Service.Spec,
		types.ServiceUpdateOptions{},
	)

	if err != nil {
		e := errors.New(fmt.Sprintf(failedToScaleServiceError, err.Error()))
		s.logger.Error(e)
		return e
	}

	if len(res.Warnings) > 0 {
		s.printAllSwarmResponseWarnings(res.Warnings)
	}

	return nil
}

func (s *service) ListServices(ctx context.Context, f filters.Args) ([]*entity.SwarmService, error) {
	//var serviceList []*entity.SwarmService

	//services, err := s.dockerClient.ServiceList(ctx, types.ServiceListOptions{Filters: f})

	return nil, nil
}

// printAllSwarmResponseWarnings prints all swarm response warnings when updating swarm.Service
func (s *service) printAllSwarmResponseWarnings(warnings []string) {
	for _, warning := range warnings {
		s.logger.Warn(warning)
	}
}
