// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"fmt"

	"google.golang.org/grpc"
)

// MaxConnectionCountSettings configures the maximum number of concurrent
// connections a gRPC server will accept.
type MaxConnectionCountSettings struct {
	// MaxConnectionCount is the maximum number of concurrent connections.
	// A value of 0 means no limit.
	MaxConnectionCount uint32 `mapstructure:"max_connection_count"`
}

// NewDefaultMaxConnectionCountSettings returns a new MaxConnectionCountSettings
// with default values.
func NewDefaultMaxConnectionCountSettings() MaxConnectionCountSettings {
	return MaxConnectionCountSettings{
		MaxConnectionCount: 0,
	}
}

// Validate checks that the MaxConnectionCountSettings is valid.
func (s MaxConnectionCountSettings) Validate() error {
	// No constraints on max connection count; 0 means unlimited.
	return nil
}

// IsDefault returns true if the MaxConnectionCountSettings is the default value.
func (s MaxConnectionCountSettings) IsDefault() bool {
	return s.MaxConnectionCount == 0
}

// ToServerOption returns a grpc.ServerOption that applies the max connection
// count setting. Returns nil if the setting is the default (no limit).
func (s MaxConnectionCountSettings) ToServerOption() (grpc.ServerOption, error) {
	if err := s.Validate(); err != nil {
		return nil, fmt.Errorf("invalid max connection count settings: %w", err)
	}
	if s.IsDefault() {
		return nil, nil
	}
	return grpc.MaxSimultaneousConnections(int(s.MaxConnectionCount)), nil
}
