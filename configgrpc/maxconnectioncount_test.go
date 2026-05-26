// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultMaxConnectionCountSettings(t *testing.T) {
	s := NewDefaultMaxConnectionCountSettings()
	assert.Equal(t, uint32(0), s.MaxConnectionCount)
}

func TestMaxConnectionCountSettingsValidate(t *testing.T) {
	tests := []struct {
		name    string
		count   uint32
		wantErr bool
	}{
		{name: "zero (unlimited)", count: 0, wantErr: false},
		{name: "small positive", count: 10, wantErr: false},
		{name: "large positive", count: 100000, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := MaxConnectionCountSettings{MaxConnectionCount: tt.count}
			err := s.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMaxConnectionCountSettingsIsDefault(t *testing.T) {
	defaultSettings := NewDefaultMaxConnectionCountSettings()
	assert.True(t, defaultSettings.IsDefault())

	nonDefault := MaxConnectionCountSettings{MaxConnectionCount: 100}
	assert.False(t, nonDefault.IsDefault())
}

func TestMaxConnectionCountSettingsToServerOption(t *testing.T) {
	t.Run("default returns nil option", func(t *testing.T) {
		s := NewDefaultMaxConnectionCountSettings()
		opt, err := s.ToServerOption()
		require.NoError(t, err)
		assert.Nil(t, opt)
	})

	t.Run("non-default returns server option", func(t *testing.T) {
		s := MaxConnectionCountSettings{MaxConnectionCount: 500}
		opt, err := s.ToServerOption()
		require.NoError(t, err)
		assert.NotNil(t, opt)
	})
}
