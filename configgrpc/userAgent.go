// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"errors"
	"fmt"
	"strings"
)

// UserAgent represents the user-agent string sent in gRPC metadata.
type UserAgent string

// Validate checks that the UserAgent is a valid non-empty string without
// control characters or invalid characters.
func (ua UserAgent) Validate() error {
	if ua == "" {
		return errors.New("user-agent must not be empty")
	}
	for _, r := range string(ua) {
		if r < 0x20 || r == 0x7F {
			return fmt.Errorf("user-agent contains invalid character: %q", r)
		}
	}
	return nil
}

// IsSet reports whether the UserAgent has been explicitly set.
func (ua UserAgent) IsSet() bool {
	return ua != ""
}

// String returns the string representation of the UserAgent.
func (ua UserAgent) String() string {
	return string(ua)
}

// WithProduct returns a new UserAgent that appends the given product token
// (e.g. "otelcol/0.1.0") to the existing user-agent string.
func (ua UserAgent) WithProduct(product string) UserAgent {
	product = strings.TrimSpace(product)
	if product == "" {
		return ua
	}
	if ua == "" {
		return UserAgent(product)
	}
	return UserAgent(string(ua) + " " + product)
}
