// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"crypto/tls"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTLSVersionConstants(t *testing.T) {
	assert.Equal(t, TLSVersion("1.0"), TLSVersion10)
	assert.Equal(t, TLSVersion("1.1"), TLSVersion11)
	assert.Equal(t, TLSVersion("1.2"), TLSVersion12)
	assert.Equal(t, TLSVersion("1.3"), TLSVersion13)
}

func TestTLSVersionValidate(t *testing.T) {
	tests := []struct {
		name    string
		version TLSVersion
		wantErr bool
	}{
		{name: "valid 1.0", version: TLSVersion10, wantErr: false},
		{name: "valid 1.1", version: TLSVersion11, wantErr: false},
		{name: "valid 1.2", version: TLSVersion12, wantErr: false},
		{name: "valid 1.3", version: TLSVersion13, wantErr: false},
		{name: "invalid empty", version: TLSVersion(""), wantErr: true},
		{name: "invalid version", version: TLSVersion("1.4"), wantErr: true},
		{name: "invalid format", version: TLSVersion("TLSv1.2"), wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.version.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestTLSVersionToUint16(t *testing.T) {
	tests := []struct {
		version  TLSVersion
		expected uint16
	}{
		{version: TLSVersion10, expected: tls.VersionTLS10},
		{version: TLSVersion11, expected: tls.VersionTLS11},
		{version: TLSVersion12, expected: tls.VersionTLS12},
		{version: TLSVersion13, expected: tls.VersionTLS13},
	}
	for _, tt := range tests {
		t.Run(string(tt.version), func(t *testing.T) {
			val, err := tt.version.ToUint16()
			require.NoError(t, err)
			assert.Equal(t, tt.expected, val)
		})
	}

	_, err := TLSVersion("bad").ToUint16()
	require.Error(t, err)
}

func TestTLSVersionString(t *testing.T) {
	assert.Equal(t, "1.2", TLSVersion12.String())
	assert.Equal(t, "1.3", TLSVersion13.String())
}

func TestTLSVersionIsSet(t *testing.T) {
	assert.False(t, TLSVersion("").IsSet())
	assert.True(t, TLSVersion12.IsSet())
	assert.True(t, TLSVersion13.IsSet())
}
