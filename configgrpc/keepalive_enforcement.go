// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"errors"
	"time"

	"google.golang.org/grpc/keepalive"
)

// KeepaliveEnforcementPolicy defines the server-side keepalive enforcement policy.
// It controls how the server handles keepalive pings from clients.
type KeepaliveEnforcementPolicy struct {
	// MinTime is the minimum amount of time a client should wait before sending
	// a keepalive ping. If the client sends pings more frequently, the server
	// will close the connection.
	MinTime time.Duration `mapstructure:"min_time"`

	// PermitWithoutStream controls whether the server allows keepalive pings
	// even when there are no active streams. If false, pings with no active
	// streams will be treated as a violation.
	PermitWithoutStream bool `mapstructure:"permit_without_stream"`
}

// NewDefaultKeepaliveEnforcementPolicy returns a KeepaliveEnforcementPolicy
// with default values matching gRPC library defaults.
func NewDefaultKeepaliveEnforcementPolicy() KeepaliveEnforcementPolicy {
	return KeepaliveEnforcementPolicy{
		MinTime:             5 * time.Minute,
		PermitWithoutStream: false,
	}
}

// Validate checks that the KeepaliveEnforcementPolicy configuration is valid.
func (k KeepaliveEnforcementPolicy) Validate() error {
	if k.MinTime < 0 {
		return errors.New("keepalive enforcement policy min_time must be non-negative")
	}
	return nil
}

// ToEnforcementPolicy converts the KeepaliveEnforcementPolicy to a
// google.golang.org/grpc/keepalive.EnforcementPolicy.
func (k KeepaliveEnforcementPolicy) ToEnforcementPolicy() keepalive.EnforcementPolicy {
	return keepalive.EnforcementPolicy{
		MinTime:             k.MinTime,
		PermitWithoutStream: k.PermitWithoutStream,
	}
}
