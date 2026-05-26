// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// MaxConnectionAgeGraceMillisSettings controls the grace period (in milliseconds)
// after MaxConnectionAge before forcibly closing connections.
//
// This is an alternative to MaxConnectionAgeGraceSettings that uses
// millisecond precision instead of Go's time.Duration.
type MaxConnectionAgeGraceMillisSettings struct {
	// GraceMillis is the additional time (ms) allowed for pending RPCs to complete
	// before forcibly closing connections. Zero means use the default.
	GraceMillis uint32 `mapstructure:"grace_millis"`
}

// NewDefaultMaxConnectionAgeGraceMillisSettings returns a new MaxConnectionAgeGraceMillisSettings
// with default values.
func NewDefaultMaxConnectionAgeGraceMillisSettings() MaxConnectionAgeGraceMillisSettings {
	return MaxConnectionAgeGraceMillisSettings{
		GraceMillis: 0,
	}
}

// Validate checks the MaxConnectionAgeGraceMillisSettings for invalid values.
func (s MaxConnectionAgeGraceMillisSettings) Validate() error {
	if s.GraceMillis > 3_600_000 {
		return fmt.Errorf("grace_millis must not exceed 3600000 (1 hour), got %d", s.GraceMillis)
	}
	return nil
}

// IsDefault returns true if the settings represent the default (zero) value.
func (s MaxConnectionAgeGraceMillisSettings) IsDefault() bool {
	return s.GraceMillis == 0
}

// ToServerOption converts the settings to a grpc.ServerOption.
// Returns nil if the settings are at their default value.
func (s MaxConnectionAgeGraceMillisSettings) ToServerOption() grpc.ServerOption {
	if s.IsDefault() {
		return nil
	}
	return grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionAgeGrace: time.Duration(s.GraceMillis) * time.Millisecond,
	})
}
