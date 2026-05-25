// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"fmt"

	"google.golang.org/grpc"
)

// MaxHeaderListSizeSettings configures the maximum size of header list that the
// client/server is prepared to accept.
type MaxHeaderListSizeSettings struct {
	// MaxHeaderListSize is the maximum size in bytes of header list that the
	// client or server is prepared to accept. Zero means no limit.
	MaxHeaderListSize uint32 `mapstructure:"max_header_list_size"`

	// enabled tracks whether a non-default value has been set.
	enabled bool
}

// NewDefaultMaxHeaderListSizeSettings returns MaxHeaderListSizeSettings with default values.
func NewDefaultMaxHeaderListSizeSettings() MaxHeaderListSizeSettings {
	return MaxHeaderListSizeSettings{}
}

// Validate checks the MaxHeaderListSizeSettings for invalid values.
func (m MaxHeaderListSizeSettings) Validate() error {
	if m.MaxHeaderListSize > 1<<30 {
		return fmt.Errorf("max_header_list_size must not exceed 1GiB (1073741824), got %d", m.MaxHeaderListSize)
	}
	return nil
}

// IsDefault returns true if the settings represent the default (unset) value.
func (m MaxHeaderListSizeSettings) IsDefault() bool {
	return m.MaxHeaderListSize == 0
}

// ToDialOption returns a grpc.DialOption that applies the max header list size
// setting for clients. Returns nil if the setting is at its default value.
func (m MaxHeaderListSizeSettings) ToDialOption() grpc.DialOption {
	if m.IsDefault() {
		return nil
	}
	return grpc.WithMaxHeaderListSize(m.MaxHeaderListSize)
}

// ToServerOption returns a grpc.ServerOption that applies the max header list
// size setting for servers. Returns nil if the setting is at its default value.
func (m MaxHeaderListSizeSettings) ToServerOption() grpc.ServerOption {
	if m.IsDefault() {
		return nil
	}
	return grpc.MaxHeaderListSize(m.MaxHeaderListSize)
}
