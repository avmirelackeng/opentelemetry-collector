// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestKeepaliveClientParametersValidate(t *testing.T) {
	tests := []struct {
		name    string
		params  KeepaliveClientParameters
		wantErr string
	}{
		{
			name:   "zero values are valid",
			params: KeepaliveClientParameters{},
		},
		{
			name: "positive values are valid",
			params: KeepaliveClientParameters{
				Time:                10 * time.Second,
				Timeout:             5 * time.Second,
				PermitWithoutStream: true,
			},
		},
		{
			name:    "negative time is invalid",
			params:  KeepaliveClientParameters{Time: -1 * time.Second},
			wantErr: "keepalive time must be non-negative",
		},
		{
			name:    "negative timeout is invalid",
			params:  KeepaliveClientParameters{Timeout: -1 * time.Second},
			wantErr: "keepalive timeout must be non-negative",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.params.Validate()
			if tt.wantErr != "" {
				require.EqualError(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestKeepaliveServerParametersValidate(t *testing.T) {
	tests := []struct {
		name    string
		params  KeepaliveServerParameters
		wantErr string
	}{
		{
			name:   "zero values are valid",
			params: KeepaliveServerParameters{},
		},
		{
			name: "positive values are valid",
			params: KeepaliveServerParameters{
				MaxConnectionIdle:     15 * time.Minute,
				MaxConnectionAge:      30 * time.Minute,
				MaxConnectionAgeGrace: 5 * time.Second,
				Time:                  5 * time.Second,
				Timeout:               1 * time.Second,
			},
		},
		{
			name:    "negative max_connection_idle is invalid",
			params:  KeepaliveServerParameters{MaxConnectionIdle: -1 * time.Second},
			wantErr: "keepalive max_connection_idle must be non-negative",
		},
		{
			name:    "negative max_connection_age is invalid",
			params:  KeepaliveServerParameters{MaxConnectionAge: -1 * time.Second},
			wantErr: "keepalive max_connection_age must be non-negative",
		},
		{
			name:    "negative max_connection_age_grace is invalid",
			params:  KeepaliveServerParameters{MaxConnectionAgeGrace: -1 * time.Second},
			wantErr: "keepalive max_connection_age_grace must be non-negative",
		},
		{
			name:    "negative time is invalid",
			params:  KeepaliveServerParameters{Time: -1 * time.Second},
			wantErr: "keepalive time must be non-negative",
		},
		{
			name:    "negative timeout is invalid",
			params:  KeepaliveServerParameters{Timeout: -1 * time.Second},
			wantErr: "keepalive timeout must be non-negative",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.params.Validate()
			if tt.wantErr != "" {
				require.EqualError(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestKeepaliveClientParametersDefaults(t *testing.T) {
	params := KeepaliveClientParameters{}
	assert.Equal(t, time.Duration(0), params.Time)
	assert.Equal(t, time.Duration(0), params.Timeout)
	assert.False(t, params.PermitWithoutStream)
}

func TestKeepaliveServerParametersDefaults(t *testing.T) {
	params := KeepaliveServerParameters{}
	assert.Equal(t, time.Duration(0), params.MaxConnectionIdle)
	assert.Equal(t, time.Duration(0), params.MaxConnectionAge)
	assert.Equal(t, time.Duration(0), params.MaxConnectionAgeGrace)
	assert.Equal(t, time.Duration(0), params.Time)
	assert.Equal(t, time.Duration(0), params.Timeout)
}
