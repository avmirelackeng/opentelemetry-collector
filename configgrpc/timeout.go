// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"errors"
	"time"
)

// TimeoutSettings holds timeout configuration for gRPC connections.
type TimeoutSettings struct {
	// DialTimeout is the timeout for establishing a connection.
	// Defaults to 0 (no timeout).
	DialTimeout time.Duration `mapstructure:"dial_timeout"`

	// RequestTimeout is the per-request timeout.
	// Defaults to 0 (no timeout).
	RequestTimeout time.Duration `mapstructure:"request_timeout"`
}

// NewDefaultTimeoutSettings returns a TimeoutSettings with default values.
func NewDefaultTimeoutSettings() TimeoutSettings {
	return TimeoutSettings{
		DialTimeout:    0,
		RequestTimeout: 0,
	}
}

// Validate checks that timeout values are non-negative.
func (t *TimeoutSettings) Validate() error {
	if t.DialTimeout < 0 {
		return errors.New("dial_timeout must be non-negative")
	}
	if t.RequestTimeout < 0 {
		return errors.New("request_timeout must be non-negative")
	}
	return nil
}

// HasDialTimeout reports whether a dial timeout is configured.
func (t *TimeoutSettings) HasDialTimeout() bool {
	return t.DialTimeout > 0
}

// HasRequestTimeout reports whether a request timeout is configured.
func (t *TimeoutSettings) HasRequestTimeout() bool {
	return t.RequestTimeout > 0
}
