// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"errors"
	"time"

	"google.golang.org/grpc/keepalive"
)

// KeepaliveBackoffConfig defines the backoff configuration for keepalive reconnection.
type KeepaliveBackoffConfig struct {
	// BaseDelay is the amount of time to wait before the first reconnection attempt.
	BaseDelay time.Duration `mapstructure:"base_delay"`
	// Multiplier is the factor by which the backoff period increases after each retry.
	Multiplier float64 `mapstructure:"multiplier"`
	// Jitter is the fraction of the backoff time to randomize.
	Jitter float64 `mapstructure:"jitter"`
	// MaxDelay is the upper bound of backoff delay.
	MaxDelay time.Duration `mapstructure:"max_delay"`
}

// NewDefaultKeepaliveBackoffConfig returns a KeepaliveBackoffConfig with default values.
func NewDefaultKeepaliveBackoffConfig() KeepaliveBackoffConfig {
	return KeepaliveBackoffConfig{
		BaseDelay:  1.0 * time.Second,
		Multiplier: 1.6,
		Jitter:     0.2,
		MaxDelay:   120 * time.Second,
	}
}

// Validate checks that the KeepaliveBackoffConfig is valid.
func (k KeepaliveBackoffConfig) Validate() error {
	if k.BaseDelay < 0 {
		return errors.New("base_delay must be non-negative")
	}
	if k.Multiplier < 1.0 {
		return errors.New("multiplier must be >= 1.0")
	}
	if k.Jitter < 0 || k.Jitter > 1.0 {
		return errors.New("jitter must be between 0 and 1")
	}
	if k.MaxDelay < 0 {
		return errors.New("max_delay must be non-negative")
	}
	if k.MaxDelay > 0 && k.MaxDelay < k.BaseDelay {
		return errors.New("max_delay must be >= base_delay")
	}
	return nil
}

// ToBackoffConfig converts KeepaliveBackoffConfig to a keepalive.BackoffConfig.
func (k KeepaliveBackoffConfig) ToBackoffConfig() keepalive.BackoffConfig {
	return keepalive.BackoffConfig{
		BaseDelay:  k.BaseDelay,
		Multiplier: k.Multiplier,
		Jitter:     k.Jitter,
		MaxDelay:   k.MaxDelay,
	}
}
