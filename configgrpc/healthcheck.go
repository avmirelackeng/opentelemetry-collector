// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"errors"
	"time"

	"google.golang.org/grpc/health/grpc_health_v1"
)

// HealthCheckService represents the gRPC health check service name.
type HealthCheckService string

const (
	// HealthCheckServiceDefault is the default (empty) service name used for health checks.
	HealthCheckServiceDefault HealthCheckService = ""
)

// HealthCheckSettings configures the gRPC health checking protocol.
type HealthCheckSettings struct {
	// Enabled controls whether health checking is enabled.
	Enabled bool `mapstructure:"enabled"`

	// ServiceName is the service name to use for health checks.
	// An empty string checks the overall server health.
	ServiceName HealthCheckService `mapstructure:"service_name"`

	// Interval is how frequently to perform the health check.
	Interval time.Duration `mapstructure:"interval"`
}

// NewDefaultHealthCheckSettings returns HealthCheckSettings with default values.
func NewDefaultHealthCheckSettings() HealthCheckSettings {
	return HealthCheckSettings{
		Enabled:     false,
		ServiceName: HealthCheckServiceDefault,
		Interval:    10 * time.Second,
	}
}

// Validate checks the HealthCheckSettings for invalid configurations.
func (h HealthCheckSettings) Validate() error {
	if !h.Enabled {
		return nil
	}
	if h.Interval <= 0 {
		return errors.New("health check interval must be positive when health checking is enabled")
	}
	return nil
}

// ToHealthCheckRequest converts the settings into a gRPC HealthCheckRequest.
func (h HealthCheckSettings) ToHealthCheckRequest() *grpc_health_v1.HealthCheckRequest {
	return &grpc_health_v1.HealthCheckRequest{
		Service: string(h.ServiceName),
	}
}
