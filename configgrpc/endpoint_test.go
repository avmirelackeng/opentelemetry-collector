// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEndpointValidate(t *testing.T) {
	tests := []struct {
		name    string
		endpt   Endpoint
		wantErr bool
	}{
		{name: "empty", endpt: "", wantErr: true},
		{name: "host:port", endpt: "localhost:4317", wantErr: false},
		{name: "ip:port", endpt: "127.0.0.1:4317", wantErr: false},
		{name: "scheme://host:port", endpt: "dns://localhost:4317", wantErr: false},
		{name: "bare hostname", endpt: "localhost", wantErr: false},
		{name: "invalid port", endpt: "localhost:notaport", wantErr: false}, // net.SplitHostPort accepts non-numeric ports
		{name: "scheme only", endpt: "dns://", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.endpt.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestEndpointString(t *testing.T) {
	e := Endpoint("localhost:4317")
	assert.Equal(t, "localhost:4317", e.String())
}

func TestEndpointIsSet(t *testing.T) {
	assert.True(t, Endpoint("localhost:4317").IsSet())
	assert.False(t, Endpoint("").IsSet())
}

func TestEndpointWithScheme(t *testing.T) {
	tests := []struct {
		name     string
		endpt    Endpoint
		scheme   ResolverScheme
		expected Endpoint
	}{
		{
			name:     "add dns scheme",
			endpt:    Endpoint("localhost:4317"),
			scheme:   ResolverSchemeDNS,
			expected: Endpoint("dns:///localhost:4317"),
		},
		{
			name:     "add passthrough scheme",
			endpt:    Endpoint("localhost:4317"),
			scheme:   ResolverSchemePassthrough,
			expected: Endpoint("passthrough:///localhost:4317"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.endpt.WithScheme(tt.scheme)
			assert.Equal(t, tt.expected, result)
		})
	}
}
