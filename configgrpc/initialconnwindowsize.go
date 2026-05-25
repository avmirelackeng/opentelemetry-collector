// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"fmt"

	"google.golang.org/grpc"
)

const (
	// defaultInitialConnWindowSize is the default initial window size for a connection.
	// A value of 0 means use the gRPC default.
	defaultInitialConnWindowSize = int32(0)

	// minInitialConnWindowSize is the minimum allowed initial connection window size.
	minInitialConnWindowSize = int32(0)

	// maxInitialConnWindowSize is the maximum allowed initial connection window size (2^31 - 1).
	maxInitialConnWindowSize = int32(1<<31 - 1)
)

// InitialConnWindowSizeSettings configures the initial window size for a connection.
// This controls how much data the sender can transmit before receiving an acknowledgement.
// Larger values can improve throughput on high-latency connections.
type InitialConnWindowSizeSettings struct {
	// Size is the initial window size for a connection in bytes.
	// Must be between 0 and 2^31-1. A value of 0 uses the gRPC default.
	Size int32 `mapstructure:"size"`
}

// NewDefaultInitialConnWindowSizeSettings returns InitialConnWindowSizeSettings with default values.
func NewDefaultInitialConnWindowSizeSettings() InitialConnWindowSizeSettings {
	return InitialConnWindowSizeSettings{
		Size: defaultInitialConnWindowSize,
	}
}

// Validate checks the InitialConnWindowSizeSettings for invalid values.
func (s InitialConnWindowSizeSettings) Validate() error {
	if s.Size < minInitialConnWindowSize {
		return fmt.Errorf("initial connection window size must be non-negative, got %d", s.Size)
	}
	if s.Size > maxInitialConnWindowSize {
		return fmt.Errorf("initial connection window size must not exceed %d, got %d", maxInitialConnWindowSize, s.Size)
	}
	return nil
}

// IsDefault returns true if the initial connection window size is set to the default value.
func (s InitialConnWindowSizeSettings) IsDefault() bool {
	return s.Size == defaultInitialConnWindowSize
}

// ToDialOption converts the InitialConnWindowSizeSettings to a gRPC dial option.
// Returns nil if the size is set to the default value (0).
func (s InitialConnWindowSizeSettings) ToDialOption() grpc.DialOption {
	if s.IsDefault() {
		return nil
	}
	return grpc.WithInitialConnWindowSize(s.Size)
}

// ToServerOption converts the InitialConnWindowSizeSettings to a gRPC server option.
// Returns nil if the size is set to the default value (0).
func (s InitialConnWindowSizeSettings) ToServerOption() grpc.ServerOption {
	if s.IsDefault() {
		return nil
	}
	return grpc.InitialConnWindowSize(s.Size)
}
