// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"fmt"

	"google.golang.org/grpc"
)

// DefaultMaxStreamsPerConn is the default maximum number of concurrent streams per connection.
const DefaultMaxStreamsPerConn uint32 = 0 // 0 means use gRPC default

// NumStreamsSettings configures the maximum number of concurrent streams per connection.
type NumStreamsSettings struct {
	// MaxStreamsPerConn sets the maximum number of concurrent streams per
	// connection. A value of 0 means use the gRPC default.
	MaxStreamsPerConn uint32 `mapstructure:"max_streams_per_conn"`
}

// NewDefaultNumStreamsSettings returns NumStreamsSettings with default values.
func NewDefaultNumStreamsSettings() NumStreamsSettings {
	return NumStreamsSettings{
		MaxStreamsPerConn: DefaultMaxStreamsPerConn,
	}
}

// Validate checks that NumStreamsSettings is valid.
func (n NumStreamsSettings) Validate() error {
	if n.MaxStreamsPerConn > 0 && n.MaxStreamsPerConn < 1 {
		return fmt.Errorf("max_streams_per_conn must be 0 (default) or a positive integer, got %d", n.MaxStreamsPerConn)
	}
	return nil
}

// IsDefault returns true if MaxStreamsPerConn is set to the default value.
func (n NumStreamsSettings) IsDefault() bool {
	return n.MaxStreamsPerConn == DefaultMaxStreamsPerConn
}

// ToDialOption converts NumStreamsSettings to a gRPC dial option.
func (n NumStreamsSettings) ToDialOption() grpc.DialOption {
	if n.IsDefault() {
		return nil
	}
	return grpc.WithMaxHeaderListSize(n.MaxStreamsPerConn)
}

// ToServerOption converts NumStreamsSettings to a gRPC server option.
func (n NumStreamsSettings) ToServerOption() grpc.ServerOption {
	if n.IsDefault() {
		return nil
	}
	return grpc.MaxConcurrentStreams(n.MaxStreamsPerConn)
}
