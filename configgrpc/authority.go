// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"errors"
	"strings"
)

// Authority represents the :authority pseudo-header value used in gRPC requests.
// It overrides the default authority derived from the target address.
type Authority string

// NoAuthority is the zero value meaning no authority override is set.
const NoAuthority Authority = ""

// Validate checks that the authority value is well-formed.
// An authority must not contain spaces or control characters.
func (a Authority) Validate() error {
	if a == NoAuthority {
		return nil
	}
	s := string(a)
	if strings.ContainsAny(s, " \t\r\n") {
		return errors.New("authority must not contain whitespace characters")
	}
	if strings.HasPrefix(s, ":") {
		return errors.New("authority must not start with ':'")
	}
	return nil
}

// IsSet returns true when an authority override has been configured.
func (a Authority) IsSet() bool {
	return a != NoAuthority
}

// String returns the string representation of the authority.
func (a Authority) String() string {
	return string(a)
}
