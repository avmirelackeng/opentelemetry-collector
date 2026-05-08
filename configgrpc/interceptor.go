// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"errors"
	"strings"
)

// InterceptorType represents the type of gRPC interceptor.
type InterceptorType string

const (
	// InterceptorTypeUnary represents a unary interceptor.
	InterceptorTypeUnary InterceptorType = "unary"
	// InterceptorTypeStream represents a streaming interceptor.
	InterceptorTypeStream InterceptorType = "stream"
)

// InterceptorSettings holds configuration for gRPC interceptors.
type InterceptorSettings struct {
	// Type specifies whether the interceptor applies to unary or stream RPCs.
	Type InterceptorType `mapstructure:"type"`
	// Enabled controls whether the interceptor is active.
	Enabled bool `mapstructure:"enabled"`
}

// NewDefaultInterceptorSettings returns an InterceptorSettings with default values.
func NewDefaultInterceptorSettings() InterceptorSettings {
	return InterceptorSettings{
		Type:    InterceptorTypeUnary,
		Enabled: true,
	}
}

// Validate checks the InterceptorSettings for configuration errors.
func (i InterceptorSettings) Validate() error {
	switch i.Type {
	case InterceptorTypeUnary, InterceptorTypeStream:
		return nil
	case "":
		return errors.New("interceptor type must not be empty")
	default:
		return errors.New("interceptor type must be one of [unary, stream], got: " + string(i.Type))
	}
}

// String returns the string representation of the InterceptorType.
func (i InterceptorType) String() string {
	return string(i)
}

// IsUnary returns true if the interceptor type is unary.
func (i InterceptorSettings) IsUnary() bool {
	return strings.EqualFold(string(i.Type), string(InterceptorTypeUnary))
}

// IsStream returns true if the interceptor type is stream.
func (i InterceptorSettings) IsStream() bool {
	return strings.EqualFold(string(i.Type), string(InterceptorTypeStream))
}
