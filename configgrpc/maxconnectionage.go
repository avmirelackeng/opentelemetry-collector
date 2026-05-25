// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"errors"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// MaxConnectionAgeSettings configures the maximum age of a server-side gRPC connection.
// After MaxAge, the server will send a GoAway and close the connection.
// After MaxAgeGrace, the connection is forcibly closed.
type MaxConnectionAgeSettings struct {
	// MaxAge is the maximum duration a connection may exist before the server
	// sends a GoAway. Zero means no limit.
	MaxAge time.Duration `mapstructure:"max_age"`

	// MaxAgeGrace is the additional grace period after MaxAge during which
	// the server waits for RPCs to complete before forcibly closing. Zero means no grace.
	MaxAgeGrace time.Duration `mapstructure:"max_age_grace"`
}

// NewDefaultMaxConnectionAgeSettings returns MaxConnectionAgeSettings with default values.
func NewDefaultMaxConnectionAgeSettings() MaxConnectionAgeSettings {
	return MaxConnectionAgeSettings{}
}

// Validate checks that the MaxConnectionAgeSettings are valid.
func (m MaxConnectionAgeSettings) Validate() error {
	if m.MaxAge < 0 {
		return errors.New("max_age must be non-negative")
	}
	if m.MaxAgeGrace < 0 {
		return errors.New("max_age_grace must be non-negative")
	}
	if m.MaxAge == 0 && m.MaxAgeGrace > 0 {
		return errors.New("max_age_grace requires max_age to be set")
	}
	return nil
}

// IsDefault returns true when no max connection age is configured.
func (m MaxConnectionAgeSettings) IsDefault() bool {
	return m.MaxAge == 0 && m.MaxAgeGrace == 0
}

// ToServerOption converts MaxConnectionAgeSettings to a grpc.ServerOption.
// Returns nil if the settings are default.
func (m MaxConnectionAgeSettings) ToServerOption() grpc.ServerOption {
	if m.IsDefault() {
		return nil
	}
	return grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionAge:      m.MaxAge,
		MaxConnectionAgeGrace: m.MaxAgeGrace,
	})
}
