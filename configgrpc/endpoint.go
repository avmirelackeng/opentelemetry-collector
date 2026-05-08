// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

// Endpoint represents a gRPC endpoint address.
type Endpoint string

// Validate checks that the endpoint is a valid host:port or scheme://host:port address.
func (e Endpoint) Validate() error {
	if e == "" {
		return errors.New("endpoint must not be empty")
	}
	s := string(e)
	// If the endpoint contains a scheme, strip it before parsing.
	if idx := strings.Index(s, "://"); idx >= 0 {
		s = s[idx+3:]
	}
	// Allow bare hostnames without port (e.g. "localhost").
	if !strings.Contains(s, ":") {
		if s == "" {
			return fmt.Errorf("endpoint %q has empty host", e)
		}
		return nil
	}
	host, port, err := net.SplitHostPort(s)
	if err != nil {
		return fmt.Errorf("endpoint %q is invalid: %w", e, err)
	}
	if host == "" && port == "" {
		return fmt.Errorf("endpoint %q has empty host and port", e)
	}
	return nil
}

// String returns the string representation of the endpoint.
func (e Endpoint) String() string {
	return string(e)
}

// WithScheme returns the endpoint with the given resolver scheme prepended,
// replacing any existing scheme.
func (e Endpoint) WithScheme(scheme ResolverScheme) Endpoint {
	return Endpoint(scheme.ApplyToEndpoint(string(e)))
}

// IsSet reports whether the endpoint has been set to a non-empty value.
func (e Endpoint) IsSet() bool {
	return e != ""
}
