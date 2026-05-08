// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBackoffConfigDefaults(t *testing.T) {
	cfg := NewDefaultBackoffConfig()
	assert.Equal(t, 1*time.Second, cfg.BaseDelay)
	assert.Equal(t, 1.6, cfg.Multiplier)
	assert.Equal(t, 0.2, cfg.Jitter)
	assert.Equal(t, 120*time.Second, cfg.MaxDelay)
}

func TestBackoffConfigValidate(t *testing.T) {
	tests := []struct {
		name    string
		cfg     BackoffConfig
		wantErr string
	}{
		{
			name: "valid default",
			cfg:  NewDefaultBackoffConfig(),
		},
		{
			name:    "negative base_delay",
			cfg:     BackoffConfig{BaseDelay: -1 * time.Second, Multiplier: 1.6, Jitter: 0.2, MaxDelay: 120 * time.Second},
			wantErr: "base_delay must be non-negative",
		},
		{
			name:    "multiplier below 1",
			cfg:     BackoffConfig{BaseDelay: 1 * time.Second, Multiplier: 0.5, Jitter: 0.2, MaxDelay: 120 * time.Second},
			wantErr: "multiplier must be at least 1.0",
		},
		{
			name:    "jitter above 1",
			cfg:     BackoffConfig{BaseDelay: 1 * time.Second, Multiplier: 1.6, Jitter: 1.5, MaxDelay: 120 * time.Second},
			wantErr: "jitter must be between 0 and 1",
		},
		{
			name:    "negative jitter",
			cfg:     BackoffConfig{BaseDelay: 1 * time.Second, Multiplier: 1.6, Jitter: -0.1, MaxDelay: 120 * time.Second},
			wantErr: "jitter must be between 0 and 1",
		},
		{
			name:    "negative max_delay",
			cfg:     BackoffConfig{BaseDelay: 1 * time.Second, Multiplier: 1.6, Jitter: 0.2, MaxDelay: -1 * time.Second},
			wantErr: "max_delay must be non-negative",
		},
		{
			name:    "base_delay exceeds max_delay",
			cfg:     BackoffConfig{BaseDelay: 10 * time.Second, Multiplier: 1.6, Jitter: 0.2, MaxDelay: 5 * time.Second},
			wantErr: "base_delay must not exceed max_delay",
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

func TestBackoffConfigIsDefault(t *testing.T) {
	cfg := NewDefaultBackoffConfig()
	assert.True(t, cfg.IsDefault())

	cfg.BaseDelay = 2 * time.Second
	assert.False(t, cfg.IsDefault())
}
