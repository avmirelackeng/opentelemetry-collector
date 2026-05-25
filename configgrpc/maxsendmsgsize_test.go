// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultMaxSendMsgSizeSettings(t *testing.T) {
	s := NewDefaultMaxSendMsgSizeSettings()
	assert.Equal(t, 0, s.MaxSendMsgSize)
	assert.True(t, s.IsDefault())
}

func TestMaxSendMsgSizeSettingsValidate(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantErr bool
	}{
		{name: "zero is valid", size: 0, wantErr: false},
		{name: "positive is valid", size: 4 * 1024 * 1024, wantErr: false},
		{name: "negative is invalid", size: -1, wantErr: true},
		{name: "large positive is valid", size: 1024 * 1024 * 1024, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := MaxSendMsgSizeSettings{MaxSendMsgSize: tt.size}
			err := s.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMaxSendMsgSizeSettingsIsDefault(t *testing.T) {
	assert.True(t, MaxSendMsgSizeSettings{MaxSendMsgSize: 0}.IsDefault())
	assert.False(t, MaxSendMsgSizeSettings{MaxSendMsgSize: 1}.IsDefault())
}

func TestMaxSendMsgSizeSettingsToDialOption(t *testing.T) {
	t.Run("default returns nil", func(t *testing.T) {
		s := NewDefaultMaxSendMsgSizeSettings()
		opt := s.ToDialOption()
		assert.Nil(t, opt)
	})
	t.Run("non-default returns option", func(t *testing.T) {
		s := MaxSendMsgSizeSettings{MaxSendMsgSize: 4 * 1024 * 1024}
		opt := s.ToDialOption()
		assert.NotNil(t, opt)
	})
}

func TestMaxSendMsgSizeSettingsToServerOption(t *testing.T) {
	t.Run("default returns nil", func(t *testing.T) {
		s := NewDefaultMaxSendMsgSizeSettings()
		opt := s.ToServerOption()
		assert.Nil(t, opt)
	})
	t.Run("non-default returns option", func(t *testing.T) {
		s := MaxSendMsgSizeSettings{MaxSendMsgSize: 8 * 1024 * 1024}
		opt := s.ToServerOption()
		assert.NotNil(t, opt)
	})
}
