// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"errors"
	"fmt"

	"google.golang.org/grpc"
)

// ConcurrencySettings holds configuration for gRPC server concurrency limits.
type ConcurrencySettings struct {
	// MaxConcurrentStreams sets the maximum number of concurrent streams
	// that each server transport can handle. Zero means no limit.
	MaxConcurrentStreams uint32 `mapstructure:"max_concurrent_streams"`

	// ReadBufferSize sets the size of the read buffer (in bytes) for each
	// connection. Zero means the default value will be used.
	ReadBufferSize int `mapstructure:"read_buffer_size"`

	// WriteBufferSize sets the size of the write buffer (in bytes) for each
	// connection. Zero means the default value will be used.
	WriteBufferSize int `mapstructure:"write_buffer_size"`
}

// NewDefaultConcurrencySettings returns a ConcurrencySettings with default values.
func NewDefaultConcurrencySettings() ConcurrencySettings {
	return ConcurrencySettings{
		MaxConcurrentStreams: 0,
		ReadBufferSize:       0,
		WriteBufferSize:      0,
	}
}

// Validate checks the ConcurrencySettings for invalid values.
func (c ConcurrencySettings) Validate() error {
	if c.ReadBufferSize < 0 {
		return errors.New("read_buffer_size must be non-negative")
	}
	if c.WriteBufferSize < 0 {
		return errors.New("write_buffer_size must be non-negative")
	}
	return nil
}

// ToServerOptions converts ConcurrencySettings to a slice of grpc.ServerOption.
func (c ConcurrencySettings) ToServerOptions() ([]grpc.ServerOption, error) {
	if err := c.Validate(); err != nil {
		return nil, fmt.Errorf("invalid concurrency settings: %w", err)
	}

	var opts []grpc.ServerOption
	if c.MaxConcurrentStreams > 0 {
		opts = append(opts, grpc.MaxConcurrentStreams(c.MaxConcurrentStreams))
	}
	if c.ReadBufferSize > 0 {
		opts = append(opts, grpc.ReadBufferSize(c.ReadBufferSize))
	}
	if c.WriteBufferSize > 0 {
		opts = append(opts, grpc.WriteBufferSize(c.WriteBufferSize))
	}
	return opts, nil
}
