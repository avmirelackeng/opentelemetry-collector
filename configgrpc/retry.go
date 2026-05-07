// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import "fmt"

// RetryPolicy defines the retry policy for gRPC calls.
type RetryPolicy struct {
	// MaxAttempts is the maximum number of attempts, including the original request.
	// Must be greater than 1 to enable retries.
	MaxAttempts uint `mapstructure:"max_attempts"`

	// InitialBackoff is the initial backoff duration in seconds.
	InitialBackoff float64 `mapstructure:"initial_backoff"`

	// MaxBackoff is the maximum backoff duration in seconds.
	MaxBackoff float64 `mapstructure:"max_backoff"`

	// BackoffMultiplier is the multiplier for the backoff duration.
	BackoffMultiplier float64 `mapstructure:"backoff_multiplier"`

	// RetryableStatusCodes is the list of gRPC status codes that are retryable.
	RetryableStatusCodes []string `mapstructure:"retryable_status_codes"`
}

// Validate checks that the RetryPolicy is valid.
func (r *RetryPolicy) Validate() error {
	if r == nil {
		return nil
	}
	if r.MaxAttempts < 2 {
		return fmt.Errorf("max_attempts must be at least 2, got %d", r.MaxAttempts)
	}
	if r.InitialBackoff <= 0 {
		return fmt.Errorf("initial_backoff must be positive, got %v", r.InitialBackoff)
	}
	if r.MaxBackoff <= 0 {
		return fmt.Errorf("max_backoff must be positive, got %v", r.MaxBackoff)
	}
	if r.MaxBackoff < r.InitialBackoff {
		return fmt.Errorf("max_backoff (%v) must be >= initial_backoff (%v)", r.MaxBackoff, r.InitialBackoff)
	}
	if r.BackoffMultiplier <= 0 {
		return fmt.Errorf("backoff_multiplier must be positive, got %v", r.BackoffMultiplier)
	}
	if len(r.RetryableStatusCodes) == 0 {
		return fmt.Errorf("retryable_status_codes must not be empty")
	}
	return nil
}
