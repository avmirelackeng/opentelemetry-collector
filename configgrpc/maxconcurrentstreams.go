// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"fmt"

	"google.golang.org/grpc"
)

// MaxConcurrentStreamsSettings configures the maximum number of concurrent
// streams that a gRPC server will allow on a single connection.
type MaxConcurrentStreamsSettings struct {
	// MaxStreams is the maximum number of concurrent streams allowed.
	// A value of 0 means no limit is applied.
	MaxStreams uint32 `mapstructure:"max_streams"`
}

// NewDefaultMaxConcurrentStreamsSettings returns a MaxConcurrentStreamsSettings
// with default values. By default, no limit is applied.
func NewDefaultMaxConcurrentStreamsSettings() MaxConcurrentStreamsSettings {
	return MaxConcurrentStreamsSettings{
		MaxStreams: 0,
	}
}

// Validate checks that the MaxConcurrentStreamsSettings is valid.
func (s MaxConcurrentStreamsSettings) Validate() error {
	// MaxStreams is a uint32, so it cannot be negative.
	// Any value is technically valid; 0 means unlimited.
	return nil
}

// IsDefault returns true if the settings represent the default (no limit).
func (s MaxConcurrentStreamsSettings) IsDefault() bool {
	return s.MaxStreams == 0
}

// ToServerOption converts the settings to a grpc.ServerOption.
// If MaxStreams is 0, nil is returned (no option applied).
func (s MaxConcurrentStreamsSettings) ToServerOption() (grpc.ServerOption, error) {
	if err := s.Validate(); err != nil {
		return nil, fmt.Errorf("invalid max concurrent streams settings: %w", err)
	}
	if s.IsDefault() {
		return nil, nil
	}
	return grpc.MaxConcurrentStreams(s.MaxStreams), nil
}
