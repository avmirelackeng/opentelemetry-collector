// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultMaxHeaderListSizeSettings(t *testing.T) {
	s := NewDefaultMaxHeaderListSizeSettings()
	assert.Equal(t, uint32(0), s.MaxHeaderListSize)
	assert.True(t, s.IsDefault())
}

func TestMaxHeaderListSizeSettingsValidate(t *testing.T) {
	tests := []struct {
		name    string
		size    uint32
		wantErr bool
	}{
		{name: "default zero", size: 0, wantErr: false},
		{name: "small value", size: 8192, wantErr: false},
		{name: "1MiB", size: 1 << 20, wantErr: false},
		{name: "1GiB", size: 1 << 30, wantErr: false},
		{name: "exceeds 1GiB", size: (1 << 30) + 1, wantErr: true},
		{name: "large value", size: ^uint32(0), wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := MaxHeaderListSizeSettings{MaxHeaderListSize: tt.size}
			err := s.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMaxHeaderListSizeSettingsIsDefault(t *testing.T) {
	assert.True(t, MaxHeaderListSizeSettings{MaxHeaderListSize: 0}.IsDefault())
	assert.False(t, MaxHeaderListSizeSettings{MaxHeaderListSize: 4096}.IsDefault())
}

func TestMaxHeaderListSizeSettingsToDialOption(t *testing.T) {
	t.Run("default returns nil", func(t *testing.T) {
		s := NewDefaultMaxHeaderListSizeSettings()
		opt := s.ToDialOption()
		assert.Nil(t, opt)
	})

	t.Run("non-default returns option", func(t *testing.T) {
		s := MaxHeaderListSizeSettings{MaxHeaderListSize: 16384}
		opt := s.ToDialOption()
		assert.NotNil(t, opt)
	})
}

func TestMaxHeaderListSizeSettingsToServerOption(t *testing.T) {
	t.Run("default returns nil", func(t *testing.T) {
		s := NewDefaultMaxHeaderListSizeSettings()
		opt := s.ToServerOption()
		assert.Nil(t, opt)
	})

	t.Run("non-default returns option", func(t *testing.T) {
		s := MaxHeaderListSizeSettings{MaxHeaderListSize: 32768}
		opt := s.ToServerOption()
		assert.NotNil(t, opt)
	})
}
