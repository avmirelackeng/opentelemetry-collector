// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultChannelzSettings(t *testing.T) {
	s := NewDefaultChannelzSettings()
	assert.False(t, s.Enabled)
	assert.Equal(t, ":0", s.Address)
}

func TestChannelzSettingsValidate(t *testing.T) {
	tests := []struct {
		name    string
		s       *ChannelzSettings
		wantErr bool
	}{
		{
			name:    "nil settings",
			s:       nil,
			wantErr: true,
		},
		{
			name:    "disabled with empty address",
			s:       &ChannelzSettings{Enabled: false, Address: ""},
			wantErr: false,
		},
		{
			name:    "enabled with valid address",
			s:       &ChannelzSettings{Enabled: true, Address: ":9090"},
			wantErr: false,
		},
		{
			name:    "enabled with empty address",
			s:       &ChannelzSettings{Enabled: true, Address: ""},
			wantErr: true,
		},
		{
			name:    "default settings",
			s:       func() *ChannelzSettings { s := NewDefaultChannelzSettings(); return &s }(),
			wantErr: false,
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

func TestChannelzSettingsToDialOption(t *testing.T) {
	t.Run("nil settings returns error", func(t *testing.T) {
		var s *ChannelzSettings
		_, err := s.ToDialOption()
		require.Error(t, err)
	})

	t.Run("invalid settings returns error", func(t *testing.T) {
		s := &ChannelzSettings{Enabled: true, Address: ""}
		_, err := s.ToDialOption()
		require.Error(t, err)
	})

	t.Run("valid settings returns dial option", func(t *testing.T) {
		s := &ChannelzSettings{Enabled: true, Address: ":9090"}
		opt, err := s.ToDialOption()
		require.NoError(t, err)
		assert.NotNil(t, opt)
	})

	t.Run("disabled settings returns dial option", func(t *testing.T) {
		s := NewDefaultChannelzSettings()
		opt, err := s.ToDialOption()
		require.NoError(t, err)
		assert.NotNil(t, opt)
	})
}
