// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultBlockingDialSettings(t *testing.T) {
	settings := NewDefaultBlockingDialSettings()
	assert.False(t, settings.Enabled)
}

func TestBlockingDialSettingsValidate(t *testing.T) {
	tests := []struct {
		name     string
		settings BlockingDialSettings
	}{
		{
			name:     "disabled",
			settings: BlockingDialSettings{Enabled: false},
		},
		{
			name:     "enabled",
			settings: BlockingDialSettings{Enabled: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.NoError(t, tt.settings.Validate())
		})
	}
}

func TestBlockingDialSettingsIsEnabled(t *testing.T) {
	disabled := BlockingDialSettings{Enabled: false}
	assert.False(t, disabled.IsEnabled())

	enabled := BlockingDialSettings{Enabled: true}
	assert.True(t, enabled.IsEnabled())
}

func TestBlockingDialSettingsToDialOptions(t *testing.T) {
	t.Run("disabled returns nil", func(t *testing.T) {
		settings := BlockingDialSettings{Enabled: false}
		opts := settings.ToDialOption()
		assert.Nil(t, opts)
	})

	t.Run("enabled returns WithBlock option", func(t *testing.T) {
		settings := BlockingDialSettings{Enabled: true}
		opts := settings.ToDialOption()
		require.Len(t, opts, 1)
		assert.NotNil(t, opts[0])
	})

	t.Run("default settings returns nil", func(t *testing.T) {
		settings := NewDefaultBlockingDialSettings()
		opts := settings.ToDialOption()
		assert.Nil(t, opts)
	})
}
