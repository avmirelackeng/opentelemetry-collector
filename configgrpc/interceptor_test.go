// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInterceptorTypeConstants(t *testing.T) {
	assert.Equal(t, InterceptorType("unary"), InterceptorTypeUnary)
	assert.Equal(t, InterceptorType("stream"), InterceptorTypeStream)
}

func TestInterceptorSettingsDefaults(t *testing.T) {
	settings := NewDefaultInterceptorSettings()
	assert.Equal(t, InterceptorTypeUnary, settings.Type)
	assert.True(t, settings.Enabled)
}

func TestInterceptorSettingsValidate(t *testing.T) {
	tests := []struct {
		name        string
		settings    InterceptorSettings
		wantErr     bool
		errContains string
	}{
		{
			name:     "valid unary",
			settings: InterceptorSettings{Type: InterceptorTypeUnary, Enabled: true},
			wantErr:  false,
		},
		{
			name:     "valid stream",
			settings: InterceptorSettings{Type: InterceptorTypeStream, Enabled: false},
			wantErr:  false,
		},
		{
			name:        "empty type",
			settings:    InterceptorSettings{Type: "", Enabled: true},
			wantErr:     true,
			errContains: "must not be empty",
		},
		{
			name:        "invalid type",
			settings:    InterceptorSettings{Type: "bidirectional", Enabled: true},
			wantErr:     true,
			errContains: "must be one of",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.settings.Validate()
			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.errContains)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestInterceptorTypeString(t *testing.T) {
	assert.Equal(t, "unary", InterceptorTypeUnary.String())
	assert.Equal(t, "stream", InterceptorTypeStream.String())
}

func TestInterceptorSettingsIsUnaryIsStream(t *testing.T) {
	unary := InterceptorSettings{Type: InterceptorTypeUnary}
	assert.True(t, unary.IsUnary())
	assert.False(t, unary.IsStream())

	stream := InterceptorSettings{Type: InterceptorTypeStream}
	assert.False(t, stream.IsUnary())
	assert.True(t, stream.IsStream())
}
