// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultServerConfig(t *testing.T) {
	cfg := NewDefaultServerConfig()
	assert.Equal(t, NewDefaultMaxMessageSizeSettings(), cfg.MaxRecvMsgSizeMiB)
	assert.Equal(t, NewDefaultInterceptorSettings(), cfg.Interceptors)
	assert.Nil(t, cfg.TLSSetting)
	assert.Nil(t, cfg.Keepalive)
}

func TestServerConfigValidate(t *testing.T) {
	tests := []struct {
		name    string
		cfg     ServerConfig
		wantErr bool
	}{
		{
			name: "valid default config",
			cfg:  NewDefaultServerConfig(),
		},
		{
			name: "valid config with endpoint",
			cfg: ServerConfig{
				Endpoint:          Endpoint("localhost:4317"),
				MaxRecvMsgSizeMiB: NewDefaultMaxMessageSizeSettings(),
				Interceptors:      NewDefaultInterceptorSettings(),
			},
		},
		{
			name: "invalid max recv msg size",
			cfg: ServerConfig{
				Endpoint:          Endpoint("localhost:4317"),
				MaxRecvMsgSizeMiB: MaxMessageSizeSettings{RecvMsgSize: -1},
				Interceptors:      NewDefaultInterceptorSettings(),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cfg.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestServerConfigToServerOptions(t *testing.T) {
	cfg := NewDefaultServerConfig()
	cfg.Endpoint = Endpoint("localhost:4317")

	opts, err := cfg.ToServerOptions()
	require.NoError(t, err)
	assert.NotEmpty(t, opts)
}

func TestServerConfigToServerOptionsWithKeepalive(t *testing.T) {
	cfg := NewDefaultServerConfig()
	cfg.Endpoint = Endpoint("localhost:4317")
	cfg.Keepalive = &KeepaliveServerParameters{}

	opts, err := cfg.ToServerOptions()
	require.NoError(t, err)
	assert.NotEmpty(t, opts)
}
