// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// IdleTimeoutSettings configures the maximum amount of time a connection may
// exist idle before it is closed. This applies to both client and server
// connections.
type IdleTimeoutSettings struct {
	// Timeout is the duration after which an idle connection is closed.
	// A zero value disables idle timeout.
	Timeout time.Duration `mapstructure:"timeout"`
}

// NewDefaultIdleTimeoutSettings returns IdleTimeoutSettings with default values.
func NewDefaultIdleTimeoutSettings() IdleTimeoutSettings {
	return IdleTimeoutSettings{
		Timeout: 0,
	}
}

// Validate checks that the IdleTimeoutSettings are valid.
func (s *IdleTimeoutSettings) Validate() error {
	if s.Timeout < 0 {
		return fmt.Errorf("idle timeout must be non-negative, got %s", s.Timeout)
	}
	return nil
}

// IsEnabled returns true if idle timeout is configured (non-zero).
func (s *IdleTimeoutSettings) IsEnabled() bool {
	return s.Timeout > 0
}

// ToDialOption returns a grpc.DialOption that configures the idle timeout on
// the client side via keepalive parameters.
func (s *IdleTimeoutSettings) ToDialOption() grpc.DialOption {
	if !s.IsEnabled() {
		return grpc.EmptyDialOption{}
	}
	return grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:    s.Timeout,
		Timeout: s.Timeout,
	})
}

// ToServerOption returns a grpc.ServerOption that configures the idle timeout
// on the server side via keepalive parameters.
func (s *IdleTimeoutSettings) ToServerOption() grpc.ServerOption {
	if !s.IsEnabled() {
		return grpc.EmptyServerOption{}
	}
	return grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: s.Timeout,
	})
}
