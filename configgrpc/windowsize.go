// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"fmt"

	"google.golang.org/grpc"
)

const (
	// DefaultInitialWindowSize is the default initial window size for gRPC streams.
	DefaultInitialWindowSize = int32(0)
	// DefaultInitialConnWindowSize is the default initial connection window size for gRPC.
	DefaultInitialConnWindowSize = int32(0)
	// MinWindowSize is the minimum allowed window size (1 byte).
	MinWindowSize = int32(1)
)

// WindowSizeSettings configures the initial window size for gRPC streams
// and connections. A value of 0 means use the gRPC default.
type WindowSizeSettings struct {
	// InitialWindowSize sets the initial window size for a stream.
	InitialWindowSize int32 `mapstructure:"initial_window_size"`
	// InitialConnWindowSize sets the initial window size for a connection.
	InitialConnWindowSize int32 `mapstructure:"initial_conn_window_size"`
}

// NewDefaultWindowSizeSettings returns a WindowSizeSettings with default values.
func NewDefaultWindowSizeSettings() WindowSizeSettings {
	return WindowSizeSettings{
		InitialWindowSize:     DefaultInitialWindowSize,
		InitialConnWindowSize: DefaultInitialConnWindowSize,
	}
}

// Validate checks that the window size settings are valid.
func (w WindowSizeSettings) Validate() error {
	if w.InitialWindowSize < 0 {
		return fmt.Errorf("initial_window_size must be non-negative, got %d", w.InitialWindowSize)
	}
	if w.InitialConnWindowSize < 0 {
		return fmt.Errorf("initial_conn_window_size must be non-negative, got %d", w.InitialConnWindowSize)
	}
	return nil
}

// ToDialOptions converts the WindowSizeSettings to gRPC dial options.
func (w WindowSizeSettings) ToDialOptions() []grpc.DialOption {
	var opts []grpc.DialOption
	if w.InitialWindowSize > 0 {
		opts = append(opts, grpc.WithInitialWindowSize(w.InitialWindowSize))
	}
	if w.InitialConnWindowSize > 0 {
		opts = append(opts, grpc.WithInitialConnWindowSize(w.InitialConnWindowSize))
	}
	return opts
}

// ToServerOptions converts the WindowSizeSettings to gRPC server options.
func (w WindowSizeSettings) ToServerOptions() []grpc.ServerOption {
	var opts []grpc.ServerOption
	if w.InitialWindowSize > 0 {
		opts = append(opts, grpc.InitialWindowSize(w.InitialWindowSize))
	}
	if w.InitialConnWindowSize > 0 {
		opts = append(opts, grpc.InitialConnWindowSize(w.InitialConnWindowSize))
	}
	return opts
}
