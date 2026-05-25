// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// MaxConnectionAgeGraceSettings configures the grace period after MaxConnectionAge
// before forcibly closing connections on the gRPC server.
type MaxConnectionAgeGraceSettings struct {
	// Grace is the additional time after MaxConnectionAge before the connection
	// is forcibly closed. A zero value means the connection is closed immediately
	// after MaxConnectionAge.
	Grace time.Duration `mapstructure:"grace"`

	// Enabled indicates whether the max connection age grace period is set.
	Enabled bool `mapstructure:"enabled"`
}

// NewDefaultMaxConnectionAgeGraceSettings returns a new MaxConnectionAgeGraceSettings
// with default values (disabled).
func NewDefaultMaxConnectionAgeGraceSettings() MaxConnectionAgeGraceSettings {
	return MaxConnectionAgeGraceSettings{
		Enabled: false,
		Grace:   0,
	}
}

// Validate checks the MaxConnectionAgeGraceSettings for invalid values.
func (s *MaxConnectionAgeGraceSettings) Validate() error {
	if s.Enabled && s.Grace < 0 {
		return fmt.Errorf("max connection age grace must be non-negative, got %s", s.Grace)
	}
	return nil
}

// IsDefault returns true if the grace period is not enabled.
func (s *MaxConnectionAgeGraceSettings) IsDefault() bool {
	return !s.Enabled
}

// ToServerOption converts the settings into a gRPC server option.
// If not enabled, it returns nil.
func (s *MaxConnectionAgeGraceSettings) ToServerOption() grpc.ServerOption {
	if !s.Enabled {
		return nil
	}
	return grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionAgeGrace: s.Grace,
	})
}
