// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultReconnectSettings(t *testing.T) {
	s := NewDefaultReconnectSettings()
	assert.True(t, s.Enabled)
	assert.Equal(t, 1*time.Second, s.InitialInterval)
	assert.Equal(t, 120*time.Second, s.MaxInterval)
	assert.Equal(t, 1.6, s.Multiplier)
}

func TestReconnectSettingsValidate(t *testing.T) {
	tests := []struct {
		name    string
		s       ReconnectSettings
		wantErr bool
	}{
		{
			name: "valid defaults",
			s:    NewDefaultReconnectSettings(),
		},
		{
			name: "negative initial interval",
			s: ReconnectSettings{
				Enabled:         true,
				InitialInterval: -1 * time.Second,
				MaxInterval:     120 * time.Second,
				Multiplier:      1.6,
			},
			wantErr: true,
		},
		{
			name: "negative max interval",
			s: ReconnectSettings{
				Enabled:         true,
				InitialInterval: 1 * time.Second,
				MaxInterval:     -1 * time.Second,
				Multiplier:      1.6,
			},
			wantErr: true,
		},
		{
			name: "max less than initial",
			s: ReconnectSettings{
				Enabled:         true,
				InitialInterval: 10 * time.Second,
				MaxInterval:     5 * time.Second,
				Multiplier:      1.6,
			},
			wantErr: true,
		},
		{
			name: "multiplier less than 1",
			s: ReconnectSettings{
				Enabled:         true,
				InitialInterval: 1 * time.Second,
				MaxInterval:     120 * time.Second,
				Multiplier:      0.5,
			},
			wantErr: true,
		},
		{
			name: "disabled with zero intervals",
			s: ReconnectSettings{
				Enabled:         false,
				InitialInterval: 0,
				MaxInterval:     0,
				Multiplier:      1.0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.s.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestReconnectSettingsToDialOption(t *testing.T) {
	s := NewDefaultReconnectSettings()
	opt := s.ToDialOption()
	require.NotNil(t, opt)

	disabled := ReconnectSettings{
		Enabled:         false,
		InitialInterval: 0,
		MaxInterval:     0,
		Multiplier:      1.0,
	}
	opt = disabled.ToDialOption()
	require.NotNil(t, opt)
}
