// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHeadersValidate(t *testing.T) {
	tests := []struct {
		name    string
		headers Headers
		wantErr bool
	}{
		{
			name:    "valid headers",
			headers: Headers{"x-tenant-id": "abc", "authorization": "Bearer token"},
			wantErr: false,
		},
		{
			name:    "empty headers",
			headers: Headers{},
			wantErr: false,
		},
		{
			name:    "nil headers",
			headers: nil,
			wantErr: false,
		},
		{
			name:    "empty key",
			headers: Headers{"": "value"},
			wantErr: true,
		},
		{
			name:    "key with spaces",
			headers: Headers{"bad key": "value"},
			wantErr: true,
		},
		{
			name:    "key with special chars",
			headers: Headers{"x-header!": "value"},
			wantErr: true,
		},
		{
			name:    "valid key with dots and underscores",
			headers: Headers{"x.custom_header": "value"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.headers.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestHeadersToStringMap(t *testing.T) {
	h := Headers{
		"x-tenant-id":   "my-tenant",
		"authorization": "Bearer abc",
	}
	m := h.ToStringMap()
	assert.Equal(t, "my-tenant", m["x-tenant-id"])
	assert.Equal(t, "Bearer abc", m["authorization"])
	assert.Len(t, m, 2)
}

func TestHeadersToStringMapEmpty(t *testing.T) {
	h := Headers{}
	m := h.ToStringMap()
	assert.Empty(t, m)
}
