// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResolverSchemeConstants(t *testing.T) {
	// Verify the known resolver scheme constants have expected string values.
	assert.Equal(t, ResolverScheme("dns"), ResolverSchemeDNS)
	assert.Equal(t, ResolverScheme("passthrough"), ResolverSchemePassthrough)
}

func TestResolverSchemeValidateKnownSchemes(t *testing.T) {
	tests := []struct {
		name    string
		scheme  ResolverScheme
		wantErr bool
	}{
		{
			name:    "dns scheme is valid",
			scheme:  ResolverSchemeDNS,
			wantErr: false,
		},
		{
			name:    "passthrough scheme is valid",
			scheme:  ResolverSchemePassthrough,
			wantErr: false,
		},
		{
			name:    "empty scheme is valid (uses default)",
			scheme:  "",
			wantErr: false,
		},
		{
			name:    "unknown scheme is invalid",
			scheme:  ResolverScheme("unknown"),
			wantErr: true,
		},
		{
			name:    "uppercase DNS is invalid",
			scheme:  ResolverScheme("DNS"),
			wantErr: true,
		},
		{
			name:    "mixed case passthrough is invalid",
			scheme:  ResolverScheme("Passthrough"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.scheme.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestResolverSchemeApplyToEndpointWithSchemes(t *testing.T) {
	tests := []struct {
		name     string
		scheme   ResolverScheme
		endpoint string
		want     string
		wantErr  bool
	}{
		{
			name:     "dns scheme prepended to bare host",
			scheme:   ResolverSchemeDNS,
			endpoint: "localhost:4317",
			want:     "dns:///localhost:4317",
		},
		{
			name:     "passthrough scheme prepended to bare host",
			scheme:   ResolverSchemePassthrough,
			endpoint: "localhost:4317",
			want:     "passthrough:///localhost:4317",
		},
		{
			name:     "empty scheme returns endpoint unchanged",
			scheme:   "",
			endpoint: "localhost:4317",
			want:     "localhost:4317",
		},
		{
			name:     "dns scheme with already-schemed endpoint returns error",
			scheme:   ResolverSchemeDNS,
			endpoint: "dns:///localhost:4317",
			wantErr:  true,
		},
		{
			name:     "passthrough scheme with already-schemed endpoint returns error",
			scheme:   ResolverSchemePassthrough,
			endpoint: "passthrough:///localhost:4317",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.scheme.ApplyToEndpoint(tt.endpoint)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
