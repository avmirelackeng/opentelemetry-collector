// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultDeadlineSettings(t *testing.T) {
	d := NewDefaultDeadlineSettings()
	assert.True(t, d.Enabled)
	assert.Equal(t, time.Duration(0), d.MaxDeadline)
}

func TestDeadlineSettingsValidate(t *testing.T) {
	tests := []struct {
		name    string
		settings DeadlineSettings
		wantErr bool
	}{
		{
			name:     "default is valid",
			settings: NewDefaultDeadlineSettings(),
			wantErr:  false,
		},
		{
			name:     "positive max deadline is valid",
			settings: DeadlineSettings{Enabled: true, MaxDeadline: 5 * time.Second},
			wantErr:  false,
		},
		{
			name:     "negative max deadline is invalid",
			settings: DeadlineSettings{Enabled: true, MaxDeadline: -1 * time.Second},
			wantErr:  true,
		},
		{
			name:     "disabled with zero max deadline is valid",
			settings: DeadlineSettings{Enabled: false, MaxDeadline: 0},
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

func TestDeadlineSettingsIsEnabled(t *testing.T) {
	assert.True(t, DeadlineSettings{Enabled: true}.IsEnabled())
	assert.False(t, DeadlineSettings{Enabled: false}.IsEnabled())
}

func TestDeadlineSettingsToDialOption(t *testing.T) {
	t.Run("enabled returns non-empty option", func(t *testing.T) {
		d := DeadlineSettings{Enabled: true, MaxDeadline: 10 * time.Second}
		opt := d.ToDialOption()
		assert.NotNil(t, opt)
	})

	t.Run("disabled returns empty option", func(t *testing.T) {
		d := DeadlineSettings{Enabled: false}
		opt := d.ToDialOption()
		assert.NotNil(t, opt)
	})
}
