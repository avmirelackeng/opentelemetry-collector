// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"crypto/tls"
	"fmt"
)

// TLSVersion represents a TLS version string.
type TLSVersion string

const (
	// TLSVersion10 represents TLS version 1.0.
	TLSVersion10 TLSVersion = "1.0"
	// TLSVersion11 represents TLS version 1.1.
	TLSVersion11 TLSVersion = "1.1"
	// TLSVersion12 represents TLS version 1.2.
	TLSVersion12 TLSVersion = "1.2"
	// TLSVersion13 represents TLS version 1.3.
	TLSVersion13 TLSVersion = "1.3"
)

// tlsVersionMap maps TLSVersion strings to crypto/tls constants.
var tlsVersionMap = map[TLSVersion]uint16{
	TLSVersion10: tls.VersionTLS10,
	TLSVersion11: tls.VersionTLS11,
	TLSVersion12: tls.VersionTLS12,
	TLSVersion13: tls.VersionTLS13,
}

// Validate checks that the TLSVersion is a known, supported version.
func (v TLSVersion) Validate() error {
	if _, ok := tlsVersionMap[v]; !ok {
		return fmt.Errorf("unsupported TLS version %q; valid values are 1.0, 1.1, 1.2, 1.3", v)
	}
	return nil
}

// ToUint16 converts the TLSVersion to the corresponding crypto/tls constant.
// Returns 0 and an error if the version is unknown.
func (v TLSVersion) ToUint16() (uint16, error) {
	val, ok := tlsVersionMap[v]
	if !ok {
		return 0, fmt.Errorf("unsupported TLS version %q", v)
	}
	return val, nil
}

// String returns the string representation of the TLSVersion.
func (v TLSVersion) String() string {
	return string(v)
}

// IsSet reports whether the TLSVersion has been explicitly configured.
func (v TLSVersion) IsSet() bool {
	return v != ""
}
