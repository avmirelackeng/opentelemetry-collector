// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/keepalive"
)

// TestKeepaliveBackoffConfigRoundTrip verifies that converting to keepalive.BackoffConfig
// and back produces consistent values.
func TestKeepaliveBackoffConfigRoundTrip(t *testing.T) {
	orig := KeepaliveBackoffConfig{
		BaseDelay:  500 * time.Millisecond,
		Multiplier: 2.0,
		Jitter:     0.1,
		MaxDelay:   60 * time.Second,
	}
	require.NoError(t, orig.Validate())

	bc := orig.ToBackoffConfig()
	assert.Equal(t, orig.BaseDelay, bc.BaseDelay)
	assert.Equal(t, orig.Multiplier, bc.Multiplier)
	assert.Equal(t, orig.Jitter, bc.Jitter)
	assert.Equal(t, orig.MaxDelay, bc.MaxDelay)
}

// TestKeepaliveBackoffConfigIsCompatibleWithKeepaliveBackoffConfig verifies
// the output type is keepalive.BackoffConfig.
func TestKeepaliveBackoffConfigIsCompatibleWithKeepaliveBackoffConfig(t *testing.T) {
	cfg := NewDefaultKeepaliveBackoffConfig()
	var bc keepalive.BackoffConfig = cfg.ToBackoffConfig()
	// Ensure the struct is usable as keepalive.BackoffConfig.
	assert.NotZero(t, bc.BaseDelay)
	assert.NotZero(t, bc.Multiplier)
}

// TestKeepaliveBackoffConfigZeroJitterIsValid ensures zero jitter is accepted.
func TestKeepaliveBackoffConfigZeroJitterIsValid(t *testing.T) {
	cfg := KeepaliveBackoffConfig{
		BaseDelay:  1 * time.Second,
		Multiplier: 1.0,
		Jitter:     0.0,
		MaxDelay:   10 * time.Second,
	}
	require.NoError(t, cfg.Validate())
}

// TestKeepaliveBackoffConfigMaxJitterIsValid ensures jitter=1.0 is accepted.
func TestKeepaliveBackoffConfigMaxJitterIsValid(t *testing.T) {
	cfg := KeepaliveBackoffConfig{
		BaseDelay:  1 * time.Second,
		Multiplier: 1.0,
		Jitter:     1.0,
		MaxDelay:   10 * time.Second,
	}
	require.NoError(t, cfg.Validate())
}
