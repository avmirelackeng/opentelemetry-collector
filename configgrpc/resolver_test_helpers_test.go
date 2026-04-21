// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestResolverSchemeApplyToEndpointEdgeCases tests additional edge cases for
// ApplyToEndpoint that are not covered in the main test file.
func TestResolverSchemeApplyToEndpointEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		scheme   ResolverScheme
		endpoint string
		want     string
		wantErr  bool
	}{
		{
			name:     "empty endpoint with dns scheme",
			scheme:   ResolverSchemeDNS,
			endpoint: "",
			wantErr:  true,
		},
		{
			name:     "endpoint already has dns scheme",
			scheme:   ResolverSchemeDNS,
			endpoint: "dns:///example.com:4317",
			want:     "dns:///example.com:4317",
		},
		{
			name:     "endpoint already has passthrough scheme",
			scheme:   ResolverSchemePassthrough,
			endpoint: "passthrough:///example.com:4317",
			want:     "passthrough:///example.com:4317",
		},
		{
			name:     "ipv6 address with dns scheme",
			scheme:   ResolverSchemeDNS,
			endpoint: "[::1]:4317",
			want:     "dns:///[::1]:4317",
		},
		{
			name:     "ipv4 address with passthrough scheme",
			scheme:   ResolverSchemePassthrough,
			endpoint: "127.0.0.1:4317",
			want:     "passthrough:///127.0.0.1:4317",
		},
		{
			name:     "hostname only without port",
			scheme:   ResolverSchemeDNS,
			endpoint: "example.com",
			want:     "dns:///example.com",
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

// TestResolverSchemeString verifies the string representation of ResolverScheme values.
func TestResolverSchemeString(t *testing.T) {
	tests := []struct {
		scheme ResolverScheme
		want   string
	}{
		{ResolverSchemeDNS, "dns"},
		{ResolverSchemePassthrough, "passthrough"},
	}

	for _, tt := range tests {
		t.Run(string(tt.scheme), func(t *testing.T) {
			assert.Equal(t, tt.want, string(tt.scheme))
		})
	}
}
