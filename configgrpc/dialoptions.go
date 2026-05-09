// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"fmt"

	"google.golang.org/grpc"
)

// DialOption represents a named gRPC dial option.
type DialOption string

const (
	// DialOptionBlock makes the Dial call block until the connection is established.
	DialOptionBlock DialOption = "block"

	// DialOptionInsecure disables transport security for the connection.
	DialOptionInsecure DialOption = "insecure"

	// DialOptionNoProxy disables proxy for the connection.
	DialOptionNoProxy DialOption = "no_proxy"
)

// DialOptionsSettings holds configuration for additional gRPC dial options.
type DialOptionsSettings struct {
	// Options is a list of named dial options to apply.
	Options []DialOption `mapstructure:"options"`
}

// NewDefaultDialOptionsSettings returns a DialOptionsSettings with default values.
func NewDefaultDialOptionsSettings() DialOptionsSettings {
	return DialOptionsSettings{
		Options: []DialOption{},
	}
}

// Validate checks that all specified dial options are known.
func (d DialOptionsSettings) Validate() error {
	for _, opt := range d.Options {
		switch opt {
		case DialOptionBlock, DialOptionInsecure, DialOptionNoProxy:
			// valid
		default:
			return fmt.Errorf("unknown dial option %q", opt)
		}
	}
	return nil
}

// ToGRPCDialOptions converts the settings into a slice of grpc.DialOption.
func (d DialOptionsSettings) ToGRPCDialOptions() ([]grpc.DialOption, error) {
	if err := d.Validate(); err != nil {
		return nil, err
	}
	var opts []grpc.DialOption
	for _, opt := range d.Options {
		switch opt {
		case DialOptionBlock:
			opts = append(opts, grpc.WithBlock()) //nolint:staticcheck
		case DialOptionInsecure:
			opts = append(opts, grpc.WithInsecure()) //nolint:staticcheck
		case DialOptionNoProxy:
			opts = append(opts, grpc.WithNoProxy())
		}
	}
	return opts, nil
}
