// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"fmt"

	"google.golang.org/grpc"
)

// MaxSendMsgSizeSettings configures the maximum message size in bytes the client/server
// can send. If 0, the default gRPC value is used (currently math.MaxInt32).
type MaxSendMsgSizeSettings struct {
	// MaxSendMsgSize is the maximum message size in bytes.
	MaxSendMsgSize int `mapstructure:"max_send_msg_size"`
}

// NewDefaultMaxSendMsgSizeSettings returns a new MaxSendMsgSizeSettings with default values.
func NewDefaultMaxSendMsgSizeSettings() MaxSendMsgSizeSettings {
	return MaxSendMsgSizeSettings{}
}

// Validate checks the MaxSendMsgSizeSettings for invalid values.
func (s MaxSendMsgSizeSettings) Validate() error {
	if s.MaxSendMsgSize < 0 {
		return fmt.Errorf("max_send_msg_size must be non-negative, got %d", s.MaxSendMsgSize)
	}
	return nil
}

// IsDefault returns true if the settings represent the default (unset) state.
func (s MaxSendMsgSizeSettings) IsDefault() bool {
	return s.MaxSendMsgSize == 0
}

// ToDialOption returns a grpc.DialOption for the max send message size.
// Returns nil if the setting is default.
func (s MaxSendMsgSizeSettings) ToDialOption() grpc.DialOption {
	if s.IsDefault() {
		return nil
	}
	return grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(s.MaxSendMsgSize))
}

// ToServerOption returns a grpc.ServerOption for the max send message size.
// Returns nil if the setting is default.
func (s MaxSendMsgSizeSettings) ToServerOption() grpc.ServerOption {
	if s.IsDefault() {
		return nil
	}
	return grpc.MaxSendMsgSize(s.MaxSendMsgSize)
}
