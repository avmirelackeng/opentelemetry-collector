// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultRateLimitSettings(t *testing.T) {
	s := NewDefaultRateLimitSettings()
	assert.Equal(t, uint32(0), s.MaxConcurrentStreams)
	assert.Equal(t, uint32(0), s.MaxConnectionsPerSecond)
}

func TestRateLimitSettingsValidate(t *testing.T) {
	tests := []struct {
		name    string
		s       RateLimitSettings
		wantErr bool
	}{
		{name: "defaults valid", s: NewDefaultRateLimitSettings(), wantErr: false},
		{name: "max streams valid", s: RateLimitSettings{MaxConcurrentStreams: 100}, wantErr: false},
		{name: "max conns valid", s: RateLimitSettings{MaxConnectionsPerSecond: 50}, wantErr: false},
		{name: "max streams too large", s: RateLimitSettings{MaxConcurrentStreams: 1000001}, wantErr: true},
		{name: "max conns too large", s: RateLimitSettings{MaxConnectionsPerSecond: 1000001}, wantErr: true},
		{name: "boundary streams valid", s: RateLimitSettings{MaxConcurrentStreams: 1000000}, wantErr: false},
		{name: "boundary conns valid", s: RateLimitSettings{MaxConnectionsPerSecond: 1000000}, wantErr: false},
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

func TestRateLimitSettingsIsEnabled(t *testing.T) {
	assert.False(t, NewDefaultRateLimitSettings().IsEnabled())
	assert.True(t, RateLimitSettings{MaxConcurrentStreams: 10}.IsEnabled())
	assert.True(t, RateLimitSettings{MaxConnectionsPerSecond: 10}.IsEnabled())
	assert.True(t, RateLimitSettings{MaxConcurrentStreams: 10, MaxConnectionsPerSecond: 10}.IsEnabled())
}

func TestRateLimitSettingsToServerOptions(t *testing.T) {
	t.Run("disabled returns empty", func(t *testing.T) {
		opts, err := NewDefaultRateLimitSettings().ToServerOptions()
		require.NoError(t, err)
		assert.Empty(t, opts)
	})

	t.Run("max concurrent streams", func(t *testing.T) {
		s := RateLimitSettings{MaxConcurrentStreams: 100}
		opts, err := s.ToServerOptions()
		require.NoError(t, err)
		assert.Len(t, opts, 1)
	})

	t.Run("max connections per second", func(t *testing.T) {
		s := RateLimitSettings{MaxConnectionsPerSecond: 50}
		opts, err := s.ToServerOptions()
		require.NoError(t, err)
		assert.Len(t, opts, 1)
	})

	t.Run("both limits set", func(t *testing.T) {
		s := RateLimitSettings{MaxConcurrentStreams: 100, MaxConnectionsPerSecond: 50}
		opts, err := s.ToServerOptions()
		require.NoError(t, err)
		assert.Len(t, opts, 2)
	})

	t.Run("invalid returns error", func(t *testing.T) {
		s := RateLimitSettings{MaxConcurrentStreams: 9999999}
		_, err := s.ToServerOptions()
		require.Error(t, err)
	})
}
