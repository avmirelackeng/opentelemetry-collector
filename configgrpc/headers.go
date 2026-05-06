// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"errors"
	"fmt"
	"strings"
)

// HeaderValue represents a single gRPC metadata header value.
type HeaderValue string

// Headers is a map of gRPC metadata header key-value pairs.
type Headers map[string]HeaderValue

// Validate checks that all header keys are non-empty and contain only
// valid characters as per the gRPC metadata specification.
func (h Headers) Validate() error {
	var errs []error
	for k := range h {
		if k == "" {
			errs = append(errs, errors.New("header key must not be empty"))
			continue
		}
		if err := validateHeaderKey(k); err != nil {
			errs = append(errs, fmt.Errorf("invalid header key %q: %w", k, err))
		}
	}
	return errors.Join(errs...)
}

// ToStringMap converts Headers to a map[string]string suitable for use
// with gRPC metadata.
func (h Headers) ToStringMap() map[string]string {
	m := make(map[string]string, len(h))
	for k, v := range h {
		m[k] = string(v)
	}
	return m
}

// validateHeaderKey checks that a header key contains only lowercase ASCII
// letters, digits, hyphens, and underscores.
func validateHeaderKey(key string) error {
	for _, c := range strings.ToLower(key) {
		if !((c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '-' || c == '_' || c == '.') {
			return fmt.Errorf("character %q is not allowed in header key", c)
		}
	}
	return nil
}
