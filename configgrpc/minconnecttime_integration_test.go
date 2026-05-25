// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMinConnectTimeSettingsRoundTrip(t *testing.T) {
	original := MinConnectTimeSettings{
		MinConnectTimeout: 15 * time.Second,
	}
	require.NoError(t, original.Validate())
	opt := original.ToDialOption()
	assert.NotNil(t, opt)
}

func TestMinConnectTimeSettingsZeroIsValid(t *testing.T) {
	settings := MinConnectTimeSettings{MinConnectTimeout: 0}
	require.NoError(t, settings.Validate())
	opt := settings.ToDialOption()
	assert.NotNil(t, opt)
}

func TestMinConnectTimeSettingsIsCompatibleWithConnectParams(t *testing.T) {
	settings := NewDefaultMinConnectTimeSettings()
	require.NoError(t, settings.Validate())

	// Verify the dial option is non-nil and can be used
	opt := settings.ToDialOption()
	assert.NotNil(t, opt, "ToDialOption should return a valid grpc.DialOption")
}

func TestMinConnectTimeSettingsLargeValueIsValid(t *testing.T) {
	settings := MinConnectTimeSettings{
		MinConnectTimeout: 10 * time.Minute,
	}
	require.NoError(t, settings.Validate())
	assert.False(t, settings.IsDefault())
}
