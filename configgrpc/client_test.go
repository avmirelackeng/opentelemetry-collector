// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultClientConfig(t *testing.T) {
	cfg := NewDefaultClientConfig()
	require.NotNil(t, cfg)

	// Verify timeout defaults
	assert.Equal(t, 10*time.Second, cfg.Timeout.RequestTimeout)

	// Verify compression default
	assert.Equal(t, CompressionNone, cfg.Compression)

	// Verify wait for ready default
	assert.Equal(t, WaitForReadyDisabled, cfg.WaitForReady)
}

func TestClientConfigValidate(t *testing.T) {
	tests := []struct {
		name    string
		modify  func(*ClientConfig)
		wantErr bool
		errMsg  string
	}{
		{
			name:    "valid default config",
			modify:  func(cfg *ClientConfig) {},
			wantErr: false,
		},
		{
			name: "valid with endpoint",
			modify: func(cfg *ClientConfig) {
				cfg.Endpoint = Endpoint("localhost:4317")
			},
			wantErr: false,
		},
		{
			name: "invalid compression",
			modify: func(cfg *ClientConfig) {
				cfg.Compression = CompressionType("invalid")
			},
			wantErr: true,
		},
		{
			name: "invalid wait for ready",
			modify: func(cfg *ClientConfig) {
				cfg.WaitForReady = WaitForReady("invalid")
			},
			wantErr: true,
		},
		{
			name: "negative request timeout",
			modify: func(cfg *ClientConfig) {
				cfg.Timeout.RequestTimeout = -1 * time.Second
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := NewDefaultClientConfig()
			tt.modify(cfg)
			err := cfg.Validate()
			if tt.wantErr {
				require.Error(t, err)
				if tt.errMsg != "" {
					assert.Contains(t, err.Error(), tt.errMsg)
				}
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestClientConfigToDialOptions(t *testing.T) {
	t.Run("default config produces dial options", func(t *testing.T) {
		cfg := NewDefaultClientConfig()
		opts, err := cfg.ToDialOptions()
		require.NoError(t, err)
		assert.NotNil(t, opts)
	})

	t.Run("config with gzip compression", func(t *testing.T) {
		cfg := NewDefaultClientConfig()
		cfg.Compression = CompressionGzip
		opts, err := cfg.ToDialOptions()
		require.NoError(t, err)
		assert.NotNil(t, opts)
	})

	t.Run("config with user agent", func(t *testing.T) {
		cfg := NewDefaultClientConfig()
		cfg.UserAgent = UserAgent("test-agent/1.0")
		opts, err := cfg.ToDialOptions()
		require.NoError(t, err)
		assert.NotNil(t, opts)
	})

	t.Run("invalid config returns error", func(t *testing.T) {
		cfg := NewDefaultClientConfig()
		cfg.Compression = CompressionType("invalid")
		_, err := cfg.ToDialOptions()
		require.Error(t, err)
	})
}
