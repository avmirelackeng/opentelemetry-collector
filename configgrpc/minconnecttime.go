// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"errors"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// MinConnectTimeSettings controls the minimum amount of time a connection
// attempt will be given to complete before timing out.
type MinConnectTimeSettings struct {
	// MinConnectTimeout is the minimum amount of time we are willing to give a
	// connection to complete. Default is 20 seconds.
	MinConnectTimeout time.Duration `mapstructure:"min_connect_timeout"`
}

// NewDefaultMinConnectTimeSettings returns a MinConnectTimeSettings with
// default values applied.
func NewDefaultMinConnectTimeSettings() MinConnectTimeSettings {
	return MinConnectTimeSettings{
		MinConnectTimeout: 20 * time.Second,
	}
}

// Validate checks that the MinConnectTimeSettings are valid.
func (m MinConnectTimeSettings) Validate() error {
	if m.MinConnectTimeout < 0 {
		return errors.New("min_connect_timeout must be non-negative")
	}
	return nil
}

// IsDefault returns true if the settings match the default values.
func (m MinConnectTimeSettings) IsDefault() bool {
	defaults := NewDefaultMinConnectTimeSettings()
	return m.MinConnectTimeout == defaults.MinConnectTimeout
}

// ToDialOption converts the MinConnectTimeSettings to a gRPC dial option.
func (m MinConnectTimeSettings) ToDialOption() grpc.DialOption {
	return grpc.WithConnectParams(grpc.ConnectParams{
		KeepaliveParams: keepalive.ClientParameters{
			Time:    m.MinConnectTimeout,
			Timeout: m.MinConnectTimeout,
		},
		MinConnectTimeout: m.MinConnectTimeout,
	})
}
