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

func TestNewDefaultIdleTimeoutSettings(t *testing.T) {
	s := NewDefaultIdleTimeoutSettings()
	assert.Equal(t, time.Duration(0), s.Timeout)
	assert.False(t, s.IsEnabled())
}

func TestIdleTimeoutSettingsValidate(t *testing.T) {
	tests := []struct {
		name    string
		s       IdleTimeoutSettings
		wantErr bool
	}{
		{name: "zero is valid", s: IdleTimeoutSettings{Timeout: 0}, wantErr: false},
		{name: "positive is valid", s: IdleTimeoutSettings{Timeout: 30 * time.Second}, wantErr: false},
		{name: "negative is invalid", s: IdleTimeoutSettings{Timeout: -1 * time.Second}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.s.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestIdleTimeoutSettingsIsEnabled(t *testing.T) {
	assert.False(t, (&IdleTimeoutSettings{Timeout: 0}).IsEnabled())
	assert.True(t, (&IdleTimeoutSettings{Timeout: time.Second}).IsEnabled())
}

func TestIdleTimeoutSettingsToDialOption(t *testing.T) {
	t.Run("disabled returns empty option", func(t *testing.T) {
		s := NewDefaultIdleTimeoutSettings()
		opt := s.ToDialOption()
		_, ok := opt.(grpc.EmptyDialOption)
		assert.True(t, ok)
	})
	t.Run("enabled returns keepalive option", func(t *testing.T) {
		s := IdleTimeoutSettings{Timeout: 10 * time.Second}
		opt := s.ToDialOption()
		_, ok := opt.(grpc.EmptyDialOption)
		assert.False(t, ok)
		assert.NotNil(t, opt)
	})
}

func TestIdleTimeoutSettingsToServerOption(t *testing.T) {
	t.Run("disabled returns empty option", func(t *testing.T) {
		s := NewDefaultIdleTimeoutSettings()
		opt := s.ToServerOption()
		_, ok := opt.(grpc.EmptyServerOption)
		assert.True(t, ok)
	})
	t.Run("enabled returns keepalive option", func(t *testing.T) {
		s := IdleTimeoutSettings{Timeout: 10 * time.Second}
		opt := s.ToServerOption()
		_, ok := opt.(grpc.EmptyServerOption)
		assert.False(t, ok)
		assert.NotNil(t, opt)
	})
}
