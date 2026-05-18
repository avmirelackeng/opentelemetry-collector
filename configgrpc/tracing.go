// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"fmt"

	"google.golang.org/grpc"
)

// TracingMode represents the tracing propagation mode for gRPC connections.
type TracingMode string

const (
	// TracingModeNone disables tracing propagation.
	TracingModeNone TracingMode = "none"
	// TracingModeEnabled enables tracing propagation via interceptors.
	TracingModeEnabled TracingMode = "enabled"
)

// TracingSettings holds configuration for gRPC tracing propagation.
type TracingSettings struct {
	// Mode controls whether tracing is propagated on gRPC calls.
	Mode TracingMode `mapstructure:"mode"`
}

// NewDefaultTracingSettings returns TracingSettings with default values.
func NewDefaultTracingSettings() TracingSettings {
	return TracingSettings{
		Mode: TracingModeNone,
	}
}

// Validate checks that the TracingSettings are valid.
func (t TracingSettings) Validate() error {
	switch t.Mode {
	case TracingModeNone, TracingModeEnabled:
		return nil
	default:
		return fmt.Errorf("invalid tracing mode %q: must be one of [none, enabled]", t.Mode)
	}
}

// IsEnabled returns true if tracing propagation is enabled.
func (t TracingSettings) IsEnabled() bool {
	return t.Mode == TracingModeEnabled
}

// ToDialOptions returns gRPC dial options for the tracing settings.
// When tracing is enabled, unary and stream client interceptors should be
// provided by the caller via otelgrpc or similar.
func (t TracingSettings) ToDialOptions() []grpc.DialOption {
	// Tracing interceptors are injected externally; this method is a placeholder
	// for future integration hooks.
	return nil
}
