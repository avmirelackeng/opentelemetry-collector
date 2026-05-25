// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultNumStreamsSettings(t *testing.T) {
	settings := NewDefaultNumStreamsSettings()
	assert.Equal(t, DefaultMaxStreamsPerConn, settings.MaxStreamsPerConn)
	assert.True(t, settings.IsDefault())
}

func TestNumStreamsSettingsValidate(t *testing.T) {
	tests := []struct {
		name     string
		settings NumStreamsSettings
		wantErr  bool
	}{
		{
			name:     "default",
			settings: NewDefaultNumStreamsSettings(),
			wantErr:  false,
		},
		{
			name:     "valid positive",
			settings: NumStreamsSettings{MaxStreamsPerConn: 100},
			wantErr:  false,
		},
		{
			name:     "zero is valid",
			settings: NumStreamsSettings{MaxStreamsPerConn: 0},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.settings.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestNumStreamsSettingsIsDefault(t *testing.T) {
	defaultSettings := NewDefaultNumStreamsSettings()
	assert.True(t, defaultSettings.IsDefault())

	customSettings := NumStreamsSettings{MaxStreamsPerConn: 50}
	assert.False(t, customSettings.IsDefault())
}

func TestNumStreamsSettingsToServerOption(t *testing.T) {
	defaultSettings := NewDefaultNumStreamsSettings()
	assert.Nil(t, defaultSettings.ToServerOption())

	customSettings := NumStreamsSettings{MaxStreamsPerConn: 100}
	opt := customSettings.ToServerOption()
	assert.NotNil(t, opt)
}

func TestNumStreamsSettingsToDialOption(t *testing.T) {
	defaultSettings := NewDefaultNumStreamsSettings()
	assert.Nil(t, defaultSettings.ToDialOption())

	customSettings := NumStreamsSettings{MaxStreamsPerConn: 100}
	opt := customSettings.ToDialOption()
	assert.NotNil(t, opt)
}
