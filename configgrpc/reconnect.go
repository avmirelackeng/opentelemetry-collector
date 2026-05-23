// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"errors"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
)

// ReconnectSettings defines the reconnection behavior for a gRPC client.
type ReconnectSettings struct {
	// Enabled controls whether reconnection is enabled.
	Enabled bool `mapstructure:"enabled"`

	// InitialInterval is the initial backoff interval for reconnection attempts.
	InitialInterval time.Duration `mapstructure:"initial_interval"`

	// MaxInterval is the maximum backoff interval for reconnection attempts.
	MaxInterval time.Duration `mapstructure:"max_interval"`

	// Multiplier is the factor by which the backoff interval is multiplied after each attempt.
	Multiplier float64 `mapstructure:"multiplier"`
}

// NewDefaultReconnectSettings returns a ReconnectSettings with default values.
func NewDefaultReconnectSettings() ReconnectSettings {
	return ReconnectSettings{
		Enabled:         true,
		InitialInterval: 1 * time.Second,
		MaxInterval:     120 * time.Second,
		Multiplier:      1.6,
	}
}

// Validate checks the ReconnectSettings for invalid values.
func (r ReconnectSettings) Validate() error {
	if r.InitialInterval < 0 {
		return errors.New("initial_interval must be non-negative")
	}
	if r.MaxInterval < 0 {
		return errors.New("max_interval must be non-negative")
	}
	if r.MaxInterval < r.InitialInterval {
		return errors.New("max_interval must be greater than or equal to initial_interval")
	}
	if r.Multiplier < 1.0 {
		return errors.New("multiplier must be >= 1.0")
	}
	return nil
}

// ToDialOption converts ReconnectSettings to a grpc.DialOption.
func (r ReconnectSettings) ToDialOption() grpc.DialOption {
	if !r.Enabled {
		return grpc.WithConnectParams(grpc.ConnectParams{
			Backoff: backoff.Config{
				BaseDelay:  0,
				Multiplier: 1.0,
				MaxDelay:   0,
			},
		})
	}
	return grpc.WithConnectParams(grpc.ConnectParams{
		Backoff: backoff.Config{
			BaseDelay:  r.InitialInterval,
			Multiplier: r.Multiplier,
			MaxDelay:   r.MaxInterval,
		},
	})
}
