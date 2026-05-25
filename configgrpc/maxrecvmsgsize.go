// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"fmt"

	"google.golang.org/grpc"
)

// MaxRecvMsgSizeSettings configures the maximum message size in bytes the
// gRPC client or server can receive.
type MaxRecvMsgSizeSettings struct {
	// MaxRecvMsgSizeBytes is the maximum message size in bytes the client/server
	// can receive. 0 means use the gRPC default (4MB for clients, unlimited for servers).
	MaxRecvMsgSizeBytes int `mapstructure:"max_recv_msg_size_bytes"`
}

// NewDefaultMaxRecvMsgSizeSettings returns MaxRecvMsgSizeSettings with default values.
func NewDefaultMaxRecvMsgSizeSettings() MaxRecvMsgSizeSettings {
	return MaxRecvMsgSizeSettings{}
}

// Validate checks the MaxRecvMsgSizeSettings for invalid values.
func (s MaxRecvMsgSizeSettings) Validate() error {
	if s.MaxRecvMsgSizeBytes < 0 {
		return fmt.Errorf("max_recv_msg_size_bytes must be non-negative, got %d", s.MaxRecvMsgSizeBytes)
	}
	return nil
}

// IsDefault returns true if the settings are at their default value.
func (s MaxRecvMsgSizeSettings) IsDefault() bool {
	return s.MaxRecvMsgSizeBytes == 0
}

// ToDialOption returns a grpc.DialOption for the max receive message size.
// Returns nil if the setting is at its default value.
func (s MaxRecvMsgSizeSettings) ToDialOption() grpc.DialOption {
	if s.IsDefault() {
		return nil
	}
	return grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(s.MaxRecvMsgSizeBytes))
}

// ToServerOption returns a grpc.ServerOption for the max receive message size.
// Returns nil if the setting is at its default value.
func (s MaxRecvMsgSizeSettings) ToServerOption() grpc.ServerOption {
	if s.IsDefault() {
		return nil
	}
	return grpc.MaxRecvMsgSize(s.MaxRecvMsgSizeBytes)
}
