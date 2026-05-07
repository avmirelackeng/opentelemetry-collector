// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuthorityConstants(t *testing.T) {
	assert.Equal(t, Authority(""), NoAuthority)
}

func TestAuthorityValidate(t *testing.T) {
	tests := []struct {
		name      string
		authority Authority
		wantErr   bool
	}{
		{name: "empty is valid", authority: NoAuthority, wantErr: false},
		{name: "simple hostname", authority: "example.com", wantErr: false},
		{name: "hostname with port", authority: "example.com:443", wantErr: false},
		{name: "with space", authority: "example .com", wantErr: true},
		{name: "with tab", authority: "example\t.com", wantErr: true},
		{name: "with newline", authority: "example\n.com", wantErr: true},
		{name: "starts with colon", authority: ":example", wantErr: true},
		{name: "ip address", authority: "192.168.1.1", wantErr: false},
		{name: "ip with port", authority: "192.168.1.1:4317", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.authority.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAuthorityIsSet(t *testing.T) {
	assert.False(t, NoAuthority.IsSet())
	assert.True(t, Authority("example.com").IsSet())
}

func TestAuthorityString(t *testing.T) {
	assert.Equal(t, "", NoAuthority.String())
	assert.Equal(t, "example.com", Authority("example.com").String())
	assert.Equal(t, "example.com:443", Authority("example.com:443").String())
}
