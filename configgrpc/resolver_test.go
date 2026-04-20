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
		{name: "empty is valid", scheme: "", wantErr: false},
		{name: "dns", scheme: ResolverSchemeDNS, wantErr: false},
		{name: "passthrough", scheme: ResolverSchemePassthrough, wantErr: false},
		{name: "ipv4", scheme: ResolverSchemeIPv4, wantErr: false},
		{name: "ipv6", scheme: ResolverSchemeIPv6, wantErr: false},
		{name: "invalid", scheme: "xds", wantErr: true},
		// xds is not supported; only dns, passthrough, ipv4, ipv6 are allowed
		{name: "invalid unix", scheme: "unix", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.scheme.Validate()
			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), "unsupported resolver scheme")
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
		{name: "empty scheme no-op", scheme: "", endpoint: "localhost:4317", want: "localhost:4317"},
		{name: "dns scheme applied", scheme: ResolverSchemeDNS, endpoint: "localhost:4317", want: "dns:///localhost:4317"},
		{name: "passthrough scheme applied", scheme: ResolverSchemePassthrough, endpoint: "localhost:4317", want: "passthrough:///localhost:4317"},
		{name: "already has scheme", scheme: ResolverSchemeDNS, endpoint: "dns:///localhost:4317", want: "dns:///localhost:4317"},
		// Verify ipv4 scheme is correctly prefixed as well
		{name: "ipv4 scheme applied", scheme: ResolverSchemeIPv4, endpoint: "127.0.0.1:4317", want: "ipv4:///127.0.0.1:4317"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.scheme.ApplyToEndpoint(tt.endpoint)
			assert.Equal(t, tt.want, got)
		})
	}
}
