// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import "google.golang.org/grpc"

// BlockingDialSettings controls whether gRPC dial operations block until
// the connection is established or fail immediately.
type BlockingDialSettings struct {
	// Enabled specifies whether blocking dial is enabled.
	// When true, the dial operation blocks until the connection is established
	// or the context is cancelled.
	Enabled bool `mapstructure:"enabled"`
}

// NewDefaultBlockingDialSettings returns a BlockingDialSettings with default values.
// By default, blocking dial is disabled.
func NewDefaultBlockingDialSettings() BlockingDialSettings {
	return BlockingDialSettings{
		Enabled: false,
	}
}

// Validate checks the BlockingDialSettings for invalid configurations.
func (b BlockingDialSettings) Validate() error {
	return nil
}

// IsEnabled returns true if blocking dial is enabled.
func (b BlockingDialSettings) IsEnabled() bool {
	return b.Enabled
}

// ToDialOption returns the gRPC dial option corresponding to this setting.
// If blocking dial is not enabled, it returns nil.
func (b BlockingDialSettings) ToDialOption() []grpc.DialOption {
	if !b.Enabled {
		return nil
	}
	//nolint:staticcheck // WithBlock is deprecated but still functional
	return []grpc.DialOption{grpc.WithBlock()}
}
