// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"fmt"

	"google.golang.org/grpc"
)

// DefaultReadBufferSize is the default read buffer size in bytes.
// A value of 0 means use the gRPC default.
const DefaultReadBufferSize = 0

// ReadBufferSizeSettings configures the read buffer size for gRPC connections.
// The read buffer size controls the size of the buffer used for reading
// incoming messages. Larger buffers can improve throughput at the cost of
// increased memory usage.
type ReadBufferSizeSettings struct {
	// ReadBufferSize sets the read buffer size in bytes for gRPC connections.
	// A value of 0 means use the gRPC default (currently 32KB).
	ReadBufferSize int `mapstructure:"read_buffer_size"`
}

// NewDefaultReadBufferSizeSettings returns a new ReadBufferSizeSettings with
// default values.
func NewDefaultReadBufferSizeSettings() ReadBufferSizeSettings {
	return ReadBufferSizeSettings{
		ReadBufferSize: DefaultReadBufferSize,
	}
}

// Validate checks that the ReadBufferSizeSettings are valid.
func (r ReadBufferSizeSettings) Validate() error {
	if r.ReadBufferSize < 0 {
		return fmt.Errorf("read_buffer_size must be non-negative, got %d", r.ReadBufferSize)
	}
	return nil
}

// ToDialOptions returns the gRPC dial options for the read buffer size settings.
// If ReadBufferSize is 0, no option is added and the gRPC default is used.
func (r ReadBufferSizeSettings) ToDialOptions() []grpc.DialOption {
	if r.ReadBufferSize == 0 {
		return nil
	}
	return []grpc.DialOption{
		grpc.WithReadBufferSize(r.ReadBufferSize),
	}
}

// ToServerOptions returns the gRPC server options for the read buffer size settings.
// If ReadBufferSize is 0, no option is added and the gRPC default is used.
func (r ReadBufferSizeSettings) ToServerOptions() []grpc.ServerOption {
	if r.ReadBufferSize == 0 {
		return nil
	}
	return []grpc.ServerOption{
		grpc.ReadBufferSize(r.ReadBufferSize),
	}
}
