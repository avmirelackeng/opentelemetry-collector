// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestNewDefaultReadBufferSizeSettings(t *testing.T) {
	settings := NewDefaultReadBufferSizeSettings()
	assert.NotNil(t, settings)
	// Default should be zero (use gRPC default)
	assert.Equal(t, 0, settings.ReadBufferSize)
}

func TestReadBufferSizeSettingsValidate(t *testing.T) {
	tests := []struct {
		name     string
		settings ReadBufferSizeSettings
		wantErr  bool
		errMsg   string
	}{
		{
			name:     "default zero value is valid",
			settings: ReadBufferSizeSettings{ReadBufferSize: 0},
			wantErr:  false,
		},
		{
			name:     "positive value is valid",
			settings: ReadBufferSizeSettings{ReadBufferSize: 32 * 1024},
			wantErr:  false,
		},
		{
			name:     "large value is valid",
			settings: ReadBufferSizeSettings{ReadBufferSize: 1024 * 1024},
			wantErr:  false,
		},
		{
			name:     "negative value is invalid",
			settings: ReadBufferSizeSettings{ReadBufferSize: -1},
			wantErr:  true,
			errMsg:   "read buffer size must be non-negative",
		},
		{
			name:     "negative large value is invalid",
			settings: ReadBufferSizeSettings{ReadBufferSize: -1024},
			wantErr:  true,
			errMsg:   "read buffer size must be non-negative",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.settings.Validate()
			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.errMsg)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestReadBufferSizeSettingsToDialOptions(t *testing.T) {
	tests := []struct {
		name           string
		settings       ReadBufferSizeSettings
		wantNumOptions int
	}{
		{
			name:           "zero value produces no dial options",
			settings:       ReadBufferSizeSettings{ReadBufferSize: 0},
			wantNumOptions: 0,
		},
		{
			name:           "positive value produces one dial option",
			settings:       ReadBufferSizeSettings{ReadBufferSize: 32 * 1024},
			wantNumOptions: 1,
		},
		{
			name:           "large value produces one dial option",
			settings:       ReadBufferSizeSettings{ReadBufferSize: 1024 * 1024},
			wantNumOptions: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := tt.settings.ToDialOptions()
			assert.Len(t, opts, tt.wantNumOptions)
			for _, opt := range opts {
				assert.Implements(t, (*grpc.DialOption)(nil), opt)
			}
		})
	}
}

func TestReadBufferSizeSettingsToServerOptions(t *testing.T) {
	tests := []struct {
		name           string
		settings       ReadBufferSizeSettings
		wantNumOptions int
	}{
		{
			name:           "zero value produces no server options",
			settings:       ReadBufferSizeSettings{ReadBufferSize: 0},
			wantNumOptions: 0,
		},
		{
			name:           "positive value produces one server option",
			settings:       ReadBufferSizeSettings{ReadBufferSize: 32 * 1024},
			wantNumOptions: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := tt.settings.ToServerOptions()
			assert.Len(t, opts, tt.wantNumOptions)
			for _, opt := range opts {
				assert.Implements(t, (*grpc.ServerOption)(nil), opt)
			}
		})
	}
}
