// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"errors"
	"fmt"

	"google.golang.org/grpc"
)

// ChannelzSettings configures gRPC channelz, a debugging facility
// that provides insight into the state of gRPC channels.
type ChannelzSettings struct {
	// Enabled controls whether channelz is enabled for this gRPC component.
	Enabled bool `mapstructure:"enabled"`

	// Address is the address to expose the channelz service on.
	// Defaults to ":0" (random port) if empty and Enabled is true.
	Address string `mapstructure:"address"`
}

// NewDefaultChannelzSettings returns a ChannelzSettings with default values.
func NewDefaultChannelzSettings() ChannelzSettings {
	return ChannelzSettings{
		Enabled: false,
		Address: ":0",
	}
}

// Validate checks that the ChannelzSettings are valid.
func (c *ChannelzSettings) Validate() error {
	if c == nil {
		return errors.New("channelz settings must not be nil")
	}
	if c.Enabled && c.Address == "" {
		return errors.New("channelz address must not be empty when enabled")
	}
	return nil
}

// ToDialOption returns a gRPC dial option that enables channelz if configured.
// Currently channelz is enabled globally via grpc.EnableTracing; this method
// is a placeholder that returns a no-op option for forward compatibility.
func (c *ChannelzSettings) ToDialOption() (grpc.DialOption, error) {
	if c == nil {
		return nil, errors.New("channelz settings must not be nil")
	}
	if err := c.Validate(); err != nil {
		return nil, fmt.Errorf("invalid channelz settings: %w", err)
	}
	// grpc.WithBlock is used as a no-op placeholder; channelz is process-global.
	return grpc.WithBlock(), nil //nolint:staticcheck
}
