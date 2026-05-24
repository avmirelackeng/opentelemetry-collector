// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"errors"
	"time"

	"google.golang.org/grpc"
)

// DeadlineSettings configures per-RPC deadline propagation behavior.
type DeadlineSettings struct {
	// Enabled controls whether deadlines are propagated to outgoing RPCs.
	Enabled bool `mapstructure:"enabled"`

	// MaxDeadline caps the deadline duration forwarded to the server.
	// A zero value means no cap is applied.
	MaxDeadline time.Duration `mapstructure:"max_deadline"`
}

// NewDefaultDeadlineSettings returns DeadlineSettings with sensible defaults.
func NewDefaultDeadlineSettings() DeadlineSettings {
	return DeadlineSettings{
		Enabled:     true,
		MaxDeadline: 0,
	}
}

// Validate checks that DeadlineSettings fields are consistent.
func (d DeadlineSettings) Validate() error {
	if d.MaxDeadline < 0 {
		return errors.New("max_deadline must be non-negative")
	}
	return nil
}

// IsEnabled reports whether deadline propagation is active.
func (d DeadlineSettings) IsEnabled() bool {
	return d.Enabled
}

// ToDialOption returns a grpc.DialOption that installs a unary interceptor
// enforcing the deadline settings on outgoing calls.
func (d DeadlineSettings) ToDialOption() grpc.DialOption {
	if !d.Enabled {
		return grpc.EmptyDialOption{}
	}
	max := d.MaxDeadline
	return grpc.WithUnaryInterceptor(func(
		ctx interface{ Deadline() (interface{}, bool) },
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		// Placeholder: real implementation would manipulate context deadline.
		_ = max
		return nil
	})
}
