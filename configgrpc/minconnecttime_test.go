// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultMinConnectTimeSettings(t *testing.T) {
	settings := NewDefaultMinConnectTimeSettings()
	assert.Equal(t, 20*time.Second, settings.MinConnectTimeout)
}

func TestMinConnectTimeSettingsValidate(t *testing.T) {
	tests := []struct {
		name    string
		setting MinConnectTimeSettings
		wantErr bool
	}{
		{
			name:    "default is valid",
			setting: NewDefaultMinConnectTimeSettings(),
			wantErr: false,
		},
		{
			name:    "zero is valid",
			setting: MinConnectTimeSettings{MinConnectTimeout: 0},
			wantErr: false,
		},
		{
			name:    "positive duration is valid",
			setting: MinConnectTimeSettings{MinConnectTimeout: 5 * time.Second},
			wantErr: false,
		},
		{
			name:    "negative duration is invalid",
			setting: MinConnectTimeSettings{MinConnectTimeout: -1 * time.Second},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.setting.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMinConnectTimeSettingsIsDefault(t *testing.T) {
	defaults := NewDefaultMinConnectTimeSettings()
	assert.True(t, defaults.IsDefault())

	custom := MinConnectTimeSettings{MinConnectTimeout: 5 * time.Second}
	assert.False(t, custom.IsDefault())
}

func TestMinConnectTimeSettingsToDialOption(t *testing.T) {
	settings := NewDefaultMinConnectTimeSettings()
	opt := settings.ToDialOption()
	assert.NotNil(t, opt)
}
