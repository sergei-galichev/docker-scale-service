package config

import (
	"time"
)

type HTTPConfig interface {
	Address() string
}

type LoggingConfig interface {
	LoggingLevel() string
}

type SwarmConfig interface {
	MinReplicas() uint64
	MaxReplicas() uint64
	ReplicasStep() uint64
	ScaleUpDelay() time.Duration
	ScaleDownDelay() time.Duration
}
