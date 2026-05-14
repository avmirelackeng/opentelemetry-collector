// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultConcurrencySettings(t *testing.T) {
	c := NewDefaultConcurrencySettings()
	assert.Equal(t, uint32(0), c.MaxConcurrentStreams)
	assert.Equal(t, 0, c.ReadBufferSize)
	assert.Equal(t, 0, c.WriteBufferSize)
}

func TestConcurrencySettingsValidate(t *testing.T) {
	tests := []struct {
		name    string
		settings ConcurrencySettings
		wantErr string
	}{
		{
			name:    "defaults are valid",
			settings: NewDefaultConcurrencySettings(),
		},
		{
			name: "valid positive values",
			settings: ConcurrencySettings{
				MaxConcurrentStreams: 100,
				ReadBufferSize:       4096,
				WriteBufferSize:      4096,
			},
		},
		{
			name: "negative read buffer size",
			settings: ConcurrencySettings{
				ReadBufferSize: -1,
			},
			wantErr: "read_buffer_size must be non-negative",
		},
		{
			name: "negative write buffer size",
			settings: ConcurrencySettings{
				WriteBufferSize: -1,
			},
			wantErr: "write_buffer_size must be non-negative",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.settings.Validate()
			if tt.wantErr != "" {
				require.ErrorContains(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestConcurrencySettingsToServerOptions(t *testing.T) {
	t.Run("defaults produce no options", func(t *testing.T) {
		c := NewDefaultConcurrencySettings()
		opts, err := c.ToServerOptions()
		require.NoError(t, err)
		assert.Empty(t, opts)
	})

	t.Run("all settings produce options", func(t *testing.T) {
		c := ConcurrencySettings{
			MaxConcurrentStreams: 50,
			ReadBufferSize:       8192,
			WriteBufferSize:      8192,
		}
		opts, err := c.ToServerOptions()
		require.NoError(t, err)
		assert.Len(t, opts, 3)
	})

	t.Run("invalid settings return error", func(t *testing.T) {
		c := ConcurrencySettings{
			ReadBufferSize: -1,
		}
		_, err := c.ToServerOptions()
		require.ErrorContains(t, err, "invalid concurrency settings")
	})
}
