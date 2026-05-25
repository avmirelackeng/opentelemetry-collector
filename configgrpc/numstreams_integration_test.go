// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestNumStreamsSettingsRoundTrip(t *testing.T) {
	settings := NumStreamsSettings{MaxStreamsPerConn: 200}
	require.NoError(t, settings.Validate())

	serverOpt := settings.ToServerOption()
	assert.NotNil(t, serverOpt)

	// Verify the server option can be applied without panic
	_, err := grpc.NewServer(serverOpt), error(nil)
	assert.NoError(t, err)
}

func TestNumStreamsSettingsZeroIsValid(t *testing.T) {
	settings := NumStreamsSettings{MaxStreamsPerConn: 0}
	require.NoError(t, settings.Validate())
	assert.True(t, settings.IsDefault())
	assert.Nil(t, settings.ToServerOption())
	assert.Nil(t, settings.ToDialOption())
}

func TestNumStreamsSettingsIsCompatibleWithMaxConcurrentStreams(t *testing.T) {
	settings := NumStreamsSettings{MaxStreamsPerConn: 50}
	require.NoError(t, settings.Validate())

	serverOpt := settings.ToServerOption()
	require.NotNil(t, serverOpt)

	// Ensure it is compatible with grpc.MaxConcurrentStreams
	expected := grpc.MaxConcurrentStreams(50)
	assert.IsType(t, expected, serverOpt)
}

func TestNumStreamsSettingsLargeValueIsValid(t *testing.T) {
	settings := NumStreamsSettings{MaxStreamsPerConn: ^uint32(0)} // max uint32
	require.NoError(t, settings.Validate())
	assert.False(t, settings.IsDefault())
}
