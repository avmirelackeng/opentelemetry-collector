// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"fmt"

	"google.golang.org/grpc"
)

// SendRecvBufferSizeSettings configures the TCP send and receive buffer sizes
// for gRPC connections. A value of 0 means use the system default.
type SendRecvBufferSizeSettings struct {
	// ReadBufferSize sets the size of the reading buffer in bytes.
	ReadBufferSize int `mapstructure:"read_buffer_size"`
	// WriteBufferSize sets the size of the writing buffer in bytes.
	WriteBufferSize int `mapstructure:"write_buffer_size"`
}

// NewDefaultSendRecvBufferSizeSettings returns SendRecvBufferSizeSettings with
// default values (0 = system default).
func NewDefaultSendRecvBufferSizeSettings() SendRecvBufferSizeSettings {
	return SendRecvBufferSizeSettings{}
}

// Validate checks the SendRecvBufferSizeSettings for invalid values.
func (s SendRecvBufferSizeSettings) Validate() error {
	if s.ReadBufferSize < 0 {
		return fmt.Errorf("read_buffer_size must be non-negative, got %d", s.ReadBufferSize)
	}
	if s.WriteBufferSize < 0 {
		return fmt.Errorf("write_buffer_size must be non-negative, got %d", s.WriteBufferSize)
	}
	return nil
}

// ToDialOptions converts the settings to gRPC dial options.
func (s SendRecvBufferSizeSettings) ToDialOptions() []grpc.DialOption {
	var opts []grpc.DialOption
	if s.ReadBufferSize > 0 {
		opts = append(opts, grpc.WithReadBufferSize(s.ReadBufferSize))
	}
	if s.WriteBufferSize > 0 {
		opts = append(opts, grpc.WithWriteBufferSize(s.WriteBufferSize))
	}
	return opts
}

// ToServerOptions converts the settings to gRPC server options.
func (s SendRecvBufferSizeSettings) ToServerOptions() []grpc.ServerOption {
	var opts []grpc.ServerOption
	if s.ReadBufferSize > 0 {
		opts = append(opts, grpc.ReadBufferSize(s.ReadBufferSize))
	}
	if s.WriteBufferSize > 0 {
		opts = append(opts, grpc.WriteBufferSize(s.WriteBufferSize))
	}
	return opts
}
