// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc // import "go.opentelemetry.io/collector/config/configgrpc"

import "fmt"

// CompressionType specifies the compression algorithm to use for gRPC connections.
type CompressionType string

const (
	// CompressionNone disables compression.
	CompressionNone CompressionType = "none"
	// CompressionGzip enables gzip compression.
	CompressionGzip CompressionType = "gzip"
	// CompressionSnappy enables snappy compression.
	CompressionSnappy CompressionType = "snappy"
	// CompressionZstd enables zstd compression.
	CompressionZstd CompressionType = "zstd"
)

var validCompressionTypes = map[CompressionType]struct{}{
	CompressionNone:   {},
	CompressionGzip:   {},
	CompressionSnappy: {},
	CompressionZstd:   {},
}

// Validate returns an error if the CompressionType is not a known value.
func (c CompressionType) Validate() error {
	if _, ok := validCompressionTypes[c]; !ok {
		return fmt.Errorf("unknown compression type %q, expected one of: none, gzip, snappy, zstd", string(c))
	}
	return nil
}

// String returns the string representation of the CompressionType.
func (c CompressionType) String() string {
	return string(c)
}

// IsEnabled returns true if compression is enabled (i.e., not CompressionNone).
func (c CompressionType) IsEnabled() bool {
	return c != CompressionNone && c != ""
}
