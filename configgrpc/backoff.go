// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"errors"
	"time"
)

// BackoffConfig defines the configuration for gRPC connection backoff.
type BackoffConfig struct {
	// BaseDelay is the amount of time to wait before the first retry attempt.
	BaseDelay time.Duration `mapstructure:"base_delay"`

	// Multiplier is the factor with which to multiply backoff after a failed retry.
	// Should ideally be greater than 1.
	Multiplier float64 `mapstructure:"multiplier"`

	// Jitter is the factor with which backoff is randomized.
	Jitter float64 `mapstructure:"jitter"`

	// MaxDelay is the upper bound of backoff delay.
	MaxDelay time.Duration `mapstructure:"max_delay"`
}

// NewDefaultBackoffConfig returns a BackoffConfig with default values.
func NewDefaultBackoffConfig() BackoffConfig {
	return BackoffConfig{
		BaseDelay:  1 * time.Second,
		Multiplier: 1.6,
		Jitter:     0.2,
		MaxDelay:   120 * time.Second,
	}
}

// Validate checks the BackoffConfig for invalid settings.
func (b *BackoffConfig) Validate() error {
	if b.BaseDelay < 0 {
		return errors.New("base_delay must be non-negative")
	}
	if b.Multiplier < 1.0 {
		return errors.New("multiplier must be at least 1.0")
	}
	if b.Jitter < 0 || b.Jitter > 1.0 {
		return errors.New("jitter must be between 0 and 1")
	}
	if b.MaxDelay < 0 {
		return errors.New("max_delay must be non-negative")
	}
	if b.MaxDelay > 0 && b.BaseDelay > b.MaxDelay {
		return errors.New("base_delay must not exceed max_delay")
	}
	return nil
}

// IsDefault returns true if the BackoffConfig matches the default values.
func (b *BackoffConfig) IsDefault() bool {
	defaults := NewDefaultBackoffConfig()
	return b.BaseDelay == defaults.BaseDelay &&
		b.Multiplier == defaults.Multiplier &&
		b.Jitter == defaults.Jitter &&
		b.MaxDelay == defaults.MaxDelay
}
