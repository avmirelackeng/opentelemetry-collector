// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCredentialsTypeConstants(t *testing.T) {
	assert.Equal(t, CredentialsType("insecure"), CredentialsTypeInsecure)
	assert.Equal(t, CredentialsType("tls"), CredentialsTypeTLS)
	assert.Equal(t, CredentialsType("local"), CredentialsTypeLocal)
}

func TestCredentialsTypeValidate(t *testing.T) {
	tests := []struct {
		name    string
		creds   CredentialsType
		wantErr bool
	}{
		{name: "insecure", creds: CredentialsTypeInsecure, wantErr: false},
		{name: "tls", creds: CredentialsTypeTLS, wantErr: false},
		{name: "local", creds: CredentialsTypeLocal, wantErr: false},
		{name: "unknown", creds: CredentialsType("unknown"), wantErr: true},
		{name: "empty", creds: CredentialsType(""), wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.creds.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestCredentialsTypeString(t *testing.T) {
	assert.Equal(t, "insecure", CredentialsTypeInsecure.String())
	assert.Equal(t, "tls", CredentialsTypeTLS.String())
	assert.Equal(t, "local", CredentialsTypeLocal.String())
}

func TestCredentialsTypeIsSecure(t *testing.T) {
	assert.False(t, CredentialsTypeInsecure.IsSecure())
	assert.True(t, CredentialsTypeTLS.IsSecure())
	assert.False(t, CredentialsTypeLocal.IsSecure())
}

func TestCredentialsSettingsDefaults(t *testing.T) {
	cs := NewDefaultCredentialsSettings()
	assert.Equal(t, CredentialsTypeInsecure, cs.Type)
	require.NoError(t, cs.Validate())
}

func TestCredentialsSettingsValidate(t *testing.T) {
	cs := &CredentialsSettings{Type: CredentialsTypeTLS}
	require.NoError(t, cs.Validate())

	cs.Type = CredentialsType("bad")
	require.Error(t, cs.Validate())

	var nilCS *CredentialsSettings
	require.Error(t, nilCS.Validate())
}
