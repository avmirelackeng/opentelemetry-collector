// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"errors"
	"fmt"
)

// CredentialsType represents the type of credentials to use for a gRPC connection.
type CredentialsType string

const (
	// CredentialsTypeInsecure indicates no transport security.
	CredentialsTypeInsecure CredentialsType = "insecure"
	// CredentialsTypeTLS indicates TLS transport security.
	CredentialsTypeTLS CredentialsType = "tls"
	// CredentialsTypeLocal indicates local credentials (for testing).
	CredentialsTypeLocal CredentialsType = "local"
)

var validCredentialsTypes = map[CredentialsType]struct{}{
	CredentialsTypeInsecure: {},
	CredentialsTypeTLS:      {},
	CredentialsTypeLocal:    {},
}

// Validate checks that the CredentialsType is a known valid value.
func (c CredentialsType) Validate() error {
	if _, ok := validCredentialsTypes[c]; !ok {
		return fmt.Errorf("unknown credentials type %q: must be one of [insecure, tls, local]", c)
	}
	return nil
}

// String returns the string representation of the CredentialsType.
func (c CredentialsType) String() string {
	return string(c)
}

// IsSecure returns true if the credentials type provides transport security.
func (c CredentialsType) IsSecure() bool {
	return c == CredentialsTypeTLS
}

// CredentialsSettings holds configuration for gRPC connection credentials.
type CredentialsSettings struct {
	// Type specifies the credentials type to use.
	Type CredentialsType `mapstructure:"type"`
}

// NewDefaultCredentialsSettings returns CredentialsSettings with default values.
func NewDefaultCredentialsSettings() CredentialsSettings {
	return CredentialsSettings{
		Type: CredentialsTypeInsecure,
	}
}

// Validate checks that the CredentialsSettings are valid.
func (cs *CredentialsSettings) Validate() error {
	if cs == nil {
		return errors.New("credentials settings must not be nil")
	}
	return cs.Type.Validate()
}
