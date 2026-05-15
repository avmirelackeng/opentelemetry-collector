// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestNewDefaultClientConfig(t *testing.T) {
	cfg := NewDefaultClientConfig()
	require.NotNil(t, cfg)

	// Verify default values are set
	assert.NotNil(t, cfg.Timeout)
	assert.NotNil(t, cfg.Keepalive)
	assert.NotNil(t, cfg.Backoff)
	assert.NotNil(t, cfg.ConnectParams)
}

func TestClientConfigValidate(t *testing.T) {
	tests := []struct {
		name    string
		cfg     func() *ClientConfig
		wantErr bool
		errMsg  string
	}{
		{
			name: "default config is valid",
			cfg:  NewDefaultClientConfig,
		},
		{
			name: "valid endpoint",
			cfg: func() *ClientConfig {
				cfg := NewDefaultClientConfig()
				cfg.Endpoint = Endpoint("localhost:4317")
				return cfg
			},
		},
		{
			name: "valid endpoint with scheme",
			cfg: func() *ClientConfig {
				cfg := NewDefaultClientConfig()
				cfg.Endpoint = Endpoint("dns:///localhost:4317")
				return cfg
			},
		},
		{
			name: "invalid timeout",
			cfg: func() *ClientConfig {
				cfg := NewDefaultClientConfig()
				cfg.Timeout.DialTimeout = -1 * time.Second
				return cfg
			},
			wantErr: true,
		},
		{
			name: "invalid compression",
			cfg: func() *ClientConfig {
				cfg := NewDefaultClientConfig()
				cfg.Compression = CompressionType("invalid")
				return cfg
			},
			wantErr: true,
		},
		{
			name: "invalid balancer",
			cfg: func() *ClientConfig {
				cfg := NewDefaultClientConfig()
				cfg.BalancerName = BalancerName("invalid")
				return cfg
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cfg().Validate()
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
		assert.IsType(t, []grpc.DialOption{}, opts)
	})

	t.Run("config with compression", func(t *testing.T) {
		cfg := NewDefaultClientConfig()
		cfg.Compression = CompressionGzip
		opts, err := cfg.ToDialOptions()
		require.NoError(t, err)
		assert.NotEmpty(t, opts)
	})

	t.Run("config with user agent", func(t *testing.T) {
		cfg := NewDefaultClientConfig()
		cfg.UserAgent = UserAgent("test-agent/1.0")
		opts, err := cfg.ToDialOptions()
		require.NoError(t, err)
		assert.NotEmpty(t, opts)
	})

	t.Run("config with authority", func(t *testing.T) {
		cfg := NewDefaultClientConfig()
		cfg.Authority = Authority("example.com")
		opts, err := cfg.ToDialOptions()
		require.NoError(t, err)
		assert.NotEmpty(t, opts)
	})
}
