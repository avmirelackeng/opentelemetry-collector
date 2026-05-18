// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTracingModeConstants(t *testing.T) {
	assert.Equal(t, TracingMode("none"), TracingModeNone)
	assert.Equal(t, TracingMode("enabled"), TracingModeEnabled)
}

func TestNewDefaultTracingSettings(t *testing.T) {
	s := NewDefaultTracingSettings()
	assert.Equal(t, TracingModeNone, s.Mode)
}

func TestTracingSettingsValidate(t *testing.T) {
	tests := []struct {
		name    string
		mode    TracingMode
		wantErr bool
	}{
		{name: "none", mode: TracingModeNone, wantErr: false},
		{name: "enabled", mode: TracingModeEnabled, wantErr: false},
		{name: "invalid", mode: TracingMode("invalid"), wantErr: true},
		{name: "empty", mode: TracingMode(""), wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := TracingSettings{Mode: tt.mode}
			err := s.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestTracingSettingsIsEnabled(t *testing.T) {
	assert.False(t, TracingSettings{Mode: TracingModeNone}.IsEnabled())
	assert.True(t, TracingSettings{Mode: TracingModeEnabled}.IsEnabled())
}

func TestTracingSettingsToDialOptions(t *testing.T) {
	s := NewDefaultTracingSettings()
	opts := s.ToDialOptions()
	assert.Nil(t, opts)

	s.Mode = TracingModeEnabled
	opts = s.ToDialOptions()
	assert.Nil(t, opts)
}
