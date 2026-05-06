// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc // import "go.opentelemetry.io/collector/configgrpc"

import (
	"errors"
	"time"
)

// KeepaliveClientParameters defines keepalive parameters for gRPC clients.
type KeepaliveClientParameters struct {
	// Time is the duration after which the client pings the server if no activity.
	Time time.Duration `mapstructure:"time"`

	// Timeout is the duration the client waits for a response before closing.
	Timeout time.Duration `mapstructure:"timeout"`

	// PermitWithoutStream allows pings even when there are no active streams.
	PermitWithoutStream bool `mapstructure:"permit_without_stream"`
}

// Validate checks that KeepaliveClientParameters values are valid.
func (k KeepaliveClientParameters) Validate() error {
	if k.Time < 0 {
		return errors.New("keepalive time must be non-negative")
	}
	if k.Timeout < 0 {
		return errors.New("keepalive timeout must be non-negative")
	}
	return nil
}

// KeepaliveServerParameters defines keepalive parameters for gRPC servers.
type KeepaliveServerParameters struct {
	// MaxConnectionIdle is the max time a connection can be idle before being closed.
	MaxConnectionIdle time.Duration `mapstructure:"max_connection_idle"`

	// MaxConnectionAge is the max duration a connection may exist before being closed.
	MaxConnectionAge time.Duration `mapstructure:"max_connection_age"`

	// MaxConnectionAgeGrace is the grace period after MaxConnectionAge before forcibly closing.
	MaxConnectionAgeGrace time.Duration `mapstructure:"max_connection_age_grace"`

	// Time is the duration after which the server pings the client if no activity.
	Time time.Duration `mapstructure:"time"`

	// Timeout is the duration the server waits for a response before closing.
	Timeout time.Duration `mapstructure:"timeout"`
}

// Validate checks that KeepaliveServerParameters values are valid.
func (k KeepaliveServerParameters) Validate() error {
	if k.MaxConnectionIdle < 0 {
		return errors.New("keepalive max_connection_idle must be non-negative")
	}
	if k.MaxConnectionAge < 0 {
		return errors.New("keepalive max_connection_age must be non-negative")
	}
	if k.MaxConnectionAgeGrace < 0 {
		return errors.New("keepalive max_connection_age_grace must be non-negative")
	}
	if k.Time < 0 {
		return errors.New("keepalive time must be non-negative")
	}
	if k.Timeout < 0 {
		return errors.New("keepalive timeout must be non-negative")
	}
	return nil
}
