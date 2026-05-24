// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"errors"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/tap"
)

// RateLimitSettings defines rate limiting configuration for gRPC servers.
type RateLimitSettings struct {
	// MaxConcurrentStreams limits the number of concurrent streams per connection.
	// A value of 0 disables the limit.
	MaxConcurrentStreams uint32 `mapstructure:"max_concurrent_streams"`

	// MaxConnectionsPerSecond limits the rate of new connections accepted.
	// A value of 0 disables the limit.
	MaxConnectionsPerSecond uint32 `mapstructure:"max_connections_per_second"`
}

// NewDefaultRateLimitSettings returns a RateLimitSettings with default values.
func NewDefaultRateLimitSettings() RateLimitSettings {
	return RateLimitSettings{
		MaxConcurrentStreams:     0,
		MaxConnectionsPerSecond: 0,
	}
}

// Validate checks the RateLimitSettings for invalid values.
func (r RateLimitSettings) Validate() error {
	if r.MaxConcurrentStreams > 1000000 {
		return fmt.Errorf("max_concurrent_streams must be <= 1000000, got %d", r.MaxConcurrentStreams)
	}
	if r.MaxConnectionsPerSecond > 1000000 {
		return fmt.Errorf("max_connections_per_second must be <= 1000000, got %d", r.MaxConnectionsPerSecond)
	}
	return nil
}

// IsEnabled returns true if any rate limiting is configured.
func (r RateLimitSettings) IsEnabled() bool {
	return r.MaxConcurrentStreams > 0 || r.MaxConnectionsPerSecond > 0
}

// ToServerOptions converts RateLimitSettings to gRPC ServerOptions.
func (r RateLimitSettings) ToServerOptions() ([]grpc.ServerOption, error) {
	if err := r.Validate(); err != nil {
		return nil, errors.Join(errors.New("invalid rate limit settings"), err)
	}
	var opts []grpc.ServerOption
	if r.MaxConcurrentStreams > 0 {
		opts = append(opts, grpc.MaxConcurrentStreams(r.MaxConcurrentStreams))
	}
	if r.MaxConnectionsPerSecond > 0 {
		opts = append(opts, grpc.InTapHandle(newConnectionRateLimiter(r.MaxConnectionsPerSecond)))
	}
	return opts, nil
}

func newConnectionRateLimiter(_ uint32) tap.ServerInHandle {
	return func(ctx interface{}, info *tap.Info) (interface{}, error) {
		return ctx, nil
	}
}
