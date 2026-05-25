// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultMaxConnectionIdleSettings(t *testing.T) {
	s := NewDefaultMaxConnectionIdleSettings()
	assert.Equal(t, time.Duration(0), s.MaxConnectionIdle)
	assert.True(t, s.IsDefault())
}

func TestMaxConnectionIdleSettingsValidate(t *testing.T) {
	tests := []struct {
		name    string
		idle    time.Duration
		wantErr bool
	}{
		{name: "zero is valid", idle: 0, wantErr: false},
		{name: "positive is valid", idle: 30 * time.Minute, wantErr: false},
		{name: "one nanosecond is valid", idle: 1, wantErr: false},
		{name: "negative is invalid", idle: -1 * time.Second, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := MaxConnectionIdleSettings{MaxConnectionIdle: tt.idle}
			err := s.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMaxConnectionIdleSettingsIsDefault(t *testing.T) {
	assert.True(t, MaxConnectionIdleSettings{MaxConnectionIdle: 0}.IsDefault())
	assert.False(t, MaxConnectionIdleSettings{MaxConnectionIdle: time.Minute}.IsDefault())
}

func TestMaxConnectionIdleSettingsToServerOption(t *testing.T) {
	idle := 15 * time.Minute
	s := MaxConnectionIdleSettings{MaxConnectionIdle: idle}
	params := s.ToServerOption()
	assert.Equal(t, idle, params.MaxConnectionIdle)
}

func TestMaxConnectionIdleSettingsDefaultToServerOption(t *testing.T) {
	s := NewDefaultMaxConnectionIdleSettings()
	params := s.ToServerOption()
	assert.Equal(t, time.Duration(0), params.MaxConnectionIdle)
}
