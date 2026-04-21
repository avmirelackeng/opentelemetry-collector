// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResolverSchemeValidate(t *testing.T) {
	tests := []struct {
		name    string
		scheme  ResolverScheme
		wantErr bool
	}{
		{
			name:    "empty scheme is valid (uses default)",
			scheme:  "",
			wantErr: false,
		},
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
			name:    "unknown scheme is invalid",
			scheme:  ResolverScheme("unknown"),
			wantErr: true,
		},
		{
			name:    "xds scheme is valid",
			scheme:  ResolverSchemeXDS,
			wantErr: false,
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

func TestResolverSchemeApplyToEndpoint(t *testing.T) {
	tests := []struct {
		name     string
		scheme   ResolverScheme
		endpoint string
		want     string
	}{
		{
			name:     "empty scheme does not modify endpoint",
			scheme:   "",
			endpoint: "localhost:4317",
			want:     "localhost:4317",
		},
		{
			name:     "dns scheme prepends dns:///",
			scheme:   ResolverSchemeDNS,
			endpoint: "localhost:4317",
			want:     "dns:///localhost:4317",
		},
		{
			name:     "passthrough scheme prepends passthrough:///",
			scheme:   ResolverSchemePassthrough,
			endpoint: "localhost:4317",
			want:     "passthrough:///localhost:4317",
		},
		{
			name:     "dns scheme with already prefixed endpoint",
			scheme:   ResolverSchemeDNS,
			endpoint: "dns:///localhost:4317",
			want:     "dns:///localhost:4317",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.scheme.ApplyToEndpoint(tt.endpoint)
			assert.Equal(t, tt.want, got)
		})
	}
}
