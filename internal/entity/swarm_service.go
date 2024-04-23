package entity

import (
	"github.com/docker/docker/api/types/swarm"
)

type SwarmService struct {
	Service *swarm.Service
}
