// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"fmt"

	"google.golang.org/grpc"
)

// InitialWindowSizeSettings configures the initial window size for gRPC streams.
// This controls the flow control window size at the stream level.
type InitialWindowSizeSettings struct {
	// InitialWindowSize sets the value for initial window size on a stream.
	// Valid values are in the range [defaultWindowSize, maxWindowSize].
	// A value of 0 means use the default gRPC value.
	InitialWindowSize int32 `mapstructure:"initial_window_size"`
}

// defaultInitialWindowSize is the default gRPC initial window size (64KB).
const defaultInitialWindowSize int32 = 65536

// maxInitialWindowSize is the maximum allowed initial window size (1GB).
const maxInitialWindowSize int32 = 1 << 30

// NewDefaultInitialWindowSizeSettings returns InitialWindowSizeSettings with default values.
func NewDefaultInitialWindowSizeSettings() InitialWindowSizeSettings {
	return InitialWindowSizeSettings{
		InitialWindowSize: 0,
	}
}

// Validate checks that the InitialWindowSizeSettings are valid.
func (s *InitialWindowSizeSettings) Validate() error {
	if s.InitialWindowSize < 0 {
		return fmt.Errorf("initial_window_size must be non-negative, got %d", s.InitialWindowSize)
	}
	if s.InitialWindowSize > maxInitialWindowSize {
		return fmt.Errorf("initial_window_size must be at most %d, got %d", maxInitialWindowSize, s.InitialWindowSize)
	}
	return nil
}

// IsDefault returns true if the InitialWindowSizeSettings uses the default value.
func (s *InitialWindowSizeSettings) IsDefault() bool {
	return s.InitialWindowSize == 0
}

// ToDialOption returns a grpc.DialOption for the initial window size, or nil if default.
func (s *InitialWindowSizeSettings) ToDialOption() grpc.DialOption {
	if s.IsDefault() {
		return nil
	}
	return grpc.WithInitialWindowSize(s.InitialWindowSize)
}

// ToServerOption returns a grpc.ServerOption for the initial window size, or nil if default.
func (s *InitialWindowSizeSettings) ToServerOption() grpc.ServerOption {
	if s.IsDefault() {
		return nil
	}
	return grpc.InitialWindowSize(s.InitialWindowSize)
}
