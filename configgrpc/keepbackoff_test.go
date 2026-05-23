// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultKeepaliveBackoffConfig(t *testing.T) {
	cfg := NewDefaultKeepaliveBackoffConfig()
	assert.Equal(t, 1.0*time.Second, cfg.BaseDelay)
	assert.Equal(t, 1.6, cfg.Multiplier)
	assert.Equal(t, 0.2, cfg.Jitter)
	assert.Equal(t, 120*time.Second, cfg.MaxDelay)
}

func TestKeepaliveBackoffConfigValidate(t *testing.T) {
	tests := []struct {
		name    string
		cfg     KeepaliveBackoffConfig
		wantErr string
	}{
		{
			name: "valid default",
			cfg:  NewDefaultKeepaliveBackoffConfig(),
		},
		{
			name:    "negative base_delay",
			cfg:     KeepaliveBackoffConfig{BaseDelay: -1 * time.Second, Multiplier: 1.6, Jitter: 0.2, MaxDelay: 120 * time.Second},
			wantErr: "base_delay must be non-negative",
		},
		{
			name:    "multiplier less than 1",
			cfg:     KeepaliveBackoffConfig{BaseDelay: 1 * time.Second, Multiplier: 0.5, Jitter: 0.2, MaxDelay: 120 * time.Second},
			wantErr: "multiplier must be >= 1.0",
		},
		{
			name:    "jitter out of range",
			cfg:     KeepaliveBackoffConfig{BaseDelay: 1 * time.Second, Multiplier: 1.6, Jitter: 1.5, MaxDelay: 120 * time.Second},
			wantErr: "jitter must be between 0 and 1",
		},
		{
			name:    "negative max_delay",
			cfg:     KeepaliveBackoffConfig{BaseDelay: 1 * time.Second, Multiplier: 1.6, Jitter: 0.2, MaxDelay: -1 * time.Second},
			wantErr: "max_delay must be non-negative",
		},
		{
			name:    "max_delay less than base_delay",
			cfg:     KeepaliveBackoffConfig{BaseDelay: 10 * time.Second, Multiplier: 1.6, Jitter: 0.2, MaxDelay: 5 * time.Second},
			wantErr: "max_delay must be >= base_delay",
		},
		{
			name: "zero max_delay is allowed",
			cfg:  KeepaliveBackoffConfig{BaseDelay: 1 * time.Second, Multiplier: 1.0, Jitter: 0.0, MaxDelay: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cfg.Validate()
			if tt.wantErr != "" {
				require.EqualError(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestKeepaliveBackoffConfigToBackoffConfig(t *testing.T) {
	cfg := NewDefaultKeepaliveBackoffConfig()
	bc := cfg.ToBackoffConfig()
	assert.Equal(t, cfg.BaseDelay, bc.BaseDelay)
	assert.Equal(t, cfg.Multiplier, bc.Multiplier)
	assert.Equal(t, cfg.Jitter, bc.Jitter)
	assert.Equal(t, cfg.MaxDelay, bc.MaxDelay)
}
