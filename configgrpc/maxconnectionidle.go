// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"fmt"
	"time"

	"google.golang.org/grpc/keepalive"
)

// MaxConnectionIdleSettings configures the maximum amount of time a connection
// may be idle before it is closed. This maps to gRPC's keepalive server
// parameter MaxConnectionIdle.
type MaxConnectionIdleSettings struct {
	// MaxConnectionIdle is the maximum duration a connection may be idle
	// before the server sends a GoAway. Zero means no limit.
	MaxConnectionIdle time.Duration `mapstructure:"max_connection_idle"`
}

// NewDefaultMaxConnectionIdleSettings returns MaxConnectionIdleSettings with
// default values (no limit).
func NewDefaultMaxConnectionIdleSettings() MaxConnectionIdleSettings {
	return MaxConnectionIdleSettings{
		MaxConnectionIdle: 0,
	}
}

// Validate checks that MaxConnectionIdleSettings is valid.
func (s MaxConnectionIdleSettings) Validate() error {
	if s.MaxConnectionIdle < 0 {
		return fmt.Errorf("max_connection_idle must be non-negative, got %s", s.MaxConnectionIdle)
	}
	return nil
}

// IsDefault returns true if MaxConnectionIdle is set to the zero value (no limit).
func (s MaxConnectionIdleSettings) IsDefault() bool {
	return s.MaxConnectionIdle == 0
}

// ToServerOption returns a keepalive.ServerParameters populated with the
// MaxConnectionIdle value, suitable for use with grpc.KeepaliveParams.
func (s MaxConnectionIdleSettings) ToServerOption() keepalive.ServerParameters {
	return keepalive.ServerParameters{
		MaxConnectionIdle: s.MaxConnectionIdle,
	}
}
