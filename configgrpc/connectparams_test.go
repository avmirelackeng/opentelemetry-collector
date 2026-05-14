// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultConnectParams(t *testing.T) {
	cp := NewDefaultConnectParams()
	assert.Equal(t, 20*time.Second, cp.MinConnectTimeout)
	assert.Equal(t, NewDefaultBackoffConfig(), cp.Backoff)
}

func TestConnectParamsValidate(t *testing.T) {
	tests := []struct {
		name    string
		params  ConnectParams
		wantErr bool
	}{
		{
			name:    "valid defaults",
			params:  NewDefaultConnectParams(),
			wantErr: false,
		},
		{
			name: "zero min_connect_timeout is valid",
			params: ConnectParams{
				MinConnectTimeout: 0,
				Backoff:           NewDefaultBackoffConfig(),
			},
			wantErr: false,
		},
		{
			name: "negative min_connect_timeout is invalid",
			params: ConnectParams{
				MinConnectTimeout: -1 * time.Second,
				Backoff:           NewDefaultBackoffConfig(),
			},
			wantErr: true,
		},
		{
			name: "invalid backoff propagates error",
			params: ConnectParams{
				MinConnectTimeout: 5 * time.Second,
				Backoff: BackoffConfig{
					BaseDelay:  -1 * time.Second,
					Multiplier: 1.6,
					Jitter:     0.2,
					MaxDelay:   120 * time.Second,
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.params.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestConnectParamsToDialOption(t *testing.T) {
	cp := NewDefaultConnectParams()
	opt := cp.ToDialOption()
	require.NotNil(t, opt)
}
