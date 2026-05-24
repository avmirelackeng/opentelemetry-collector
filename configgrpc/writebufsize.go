// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"fmt"

	"google.golang.org/grpc"
)

// WriteBufferSizeSettings configures the write buffer size for gRPC connections.
// The write buffer size determines how much data can be batched before being
// sent over the network.
type WriteBufferSizeSettings struct {
	// WriteBufferSize sets the write buffer size in bytes for the gRPC connection.
	// A value of 0 uses the gRPC default (32 KiB).
	WriteBufferSize int `mapstructure:"write_buffer_size"`
}

// NewDefaultWriteBufferSizeSettings returns a WriteBufferSizeSettings with default values.
func NewDefaultWriteBufferSizeSettings() WriteBufferSizeSettings {
	return WriteBufferSizeSettings{
		WriteBufferSize: 0,
	}
}

// Validate checks the WriteBufferSizeSettings for invalid values.
func (s WriteBufferSizeSettings) Validate() error {
	if s.WriteBufferSize < 0 {
		return fmt.Errorf("write_buffer_size must be non-negative, got %d", s.WriteBufferSize)
	}
	return nil
}

// ToDialOptions returns the gRPC dial options for the write buffer size settings.
func (s WriteBufferSizeSettings) ToDialOptions() []grpc.DialOption {
	if s.WriteBufferSize > 0 {
		return []grpc.DialOption{grpc.WithWriteBufferSize(s.WriteBufferSize)}
	}
	return nil
}

// ToServerOptions returns the gRPC server options for the write buffer size settings.
func (s WriteBufferSizeSettings) ToServerOptions() []grpc.ServerOption {
	if s.WriteBufferSize > 0 {
		return []grpc.ServerOption{grpc.WriteBufferSize(s.WriteBufferSize)}
	}
	return nil
}
