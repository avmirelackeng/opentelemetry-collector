// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultMaxConcurrentStreamsSettings(t *testing.T) {
	s := NewDefaultMaxConcurrentStreamsSettings()
	assert.Equal(t, uint32(0), s.MaxStreams)
	assert.True(t, s.IsDefault())
}

func TestMaxConcurrentStreamsSettingsValidate(t *testing.T) {
	tests := []struct {
		name    string
		streams uint32
		wantErr bool
	}{
		{name: "zero (unlimited)", streams: 0, wantErr: false},
		{name: "one stream", streams: 1, wantErr: false},
		{name: "hundred streams", streams: 100, wantErr: false},
		{name: "max uint32", streams: ^uint32(0), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := MaxConcurrentStreamsSettings{MaxStreams: tt.streams}
			err := s.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMaxConcurrentStreamsSettingsIsDefault(t *testing.T) {
	assert.True(t, MaxConcurrentStreamsSettings{MaxStreams: 0}.IsDefault())
	assert.False(t, MaxConcurrentStreamsSettings{MaxStreams: 1}.IsDefault())
	assert.False(t, MaxConcurrentStreamsSettings{MaxStreams: 100}.IsDefault())
}

func TestMaxConcurrentStreamsSettingsToServerOption(t *testing.T) {
	t.Run("default returns nil option", func(t *testing.T) {
		s := NewDefaultMaxConcurrentStreamsSettings()
		opt, err := s.ToServerOption()
		require.NoError(t, err)
		assert.Nil(t, opt)
	})

	t.Run("non-zero returns server option", func(t *testing.T) {
		s := MaxConcurrentStreamsSettings{MaxStreams: 50}
		opt, err := s.ToServerOption()
		require.NoError(t, err)
		assert.NotNil(t, opt)
	})

	t.Run("max uint32 returns server option", func(t *testing.T) {
		s := MaxConcurrentStreamsSettings{MaxStreams: ^uint32(0)}
		opt, err := s.ToServerOption()
		require.NoError(t, err)
		assert.NotNil(t, opt)
	})
}
