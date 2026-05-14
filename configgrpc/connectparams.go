// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
)

// ConnectParams holds parameters for gRPC connection establishment,
// including backoff configuration for reconnection attempts.
type ConnectParams struct {
	// MinConnectTimeout is the minimum amount of time we are willing to give a
	// connection to complete. Default is 20 seconds.
	MinConnectTimeout time.Duration `mapstructure:"min_connect_timeout"`

	// Backoff holds parameters for the backoff strategy used during reconnection.
	Backoff BackoffConfig `mapstructure:"backoff"`
}

// NewDefaultConnectParams returns a ConnectParams with sensible defaults.
func NewDefaultConnectParams() ConnectParams {
	return ConnectParams{
		MinConnectTimeout: 20 * time.Second,
		Backoff:           NewDefaultBackoffConfig(),
	}
}

// Validate checks that ConnectParams fields are valid.
func (c ConnectParams) Validate() error {
	if c.MinConnectTimeout < 0 {
		return fmt.Errorf("min_connect_timeout must be non-negative, got %s", c.MinConnectTimeout)
	}
	if err := c.Backoff.Validate(); err != nil {
		return fmt.Errorf("backoff: %w", err)
	}
	return nil
}

// ToDialOption converts ConnectParams into a grpc.DialOption.
func (c ConnectParams) ToDialOption() grpc.DialOption {
	return grpc.WithConnectParams(grpc.ConnectParams{
		MinConnectTimeout: c.MinConnectTimeout,
		Backoff: backoff.Config{
			BaseDelay:  c.Backoff.BaseDelay,
			Multiplier: c.Backoff.Multiplier,
			Jitter:     c.Backoff.Jitter,
			MaxDelay:   c.Backoff.MaxDelay,
		},
	})
}
