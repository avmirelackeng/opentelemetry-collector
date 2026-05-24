// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultWriteBufferSizeSettings(t *testing.T) {
	s := NewDefaultWriteBufferSizeSettings()
	assert.Equal(t, 0, s.WriteBufferSize)
}

func TestWriteBufferSizeSettingsValidate(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantErr bool
	}{
		{name: "zero is valid", size: 0, wantErr: false},
		{name: "positive is valid", size: 32 * 1024, wantErr: false},
		{name: "negative is invalid", size: -1, wantErr: true},
		{name: "large positive is valid", size: 1024 * 1024, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := WriteBufferSizeSettings{WriteBufferSize: tt.size}
			err := s.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestWriteBufferSizeSettingsToDialOptions(t *testing.T) {
	t.Run("zero returns no options", func(t *testing.T) {
		s := WriteBufferSizeSettings{WriteBufferSize: 0}
		opts := s.ToDialOptions()
		assert.Empty(t, opts)
	})

	t.Run("positive returns dial option", func(t *testing.T) {
		s := WriteBufferSizeSettings{WriteBufferSize: 64 * 1024}
		opts := s.ToDialOptions()
		assert.Len(t, opts, 1)
	})
}

func TestWriteBufferSizeSettingsToServerOptions(t *testing.T) {
	t.Run("zero returns no options", func(t *testing.T) {
		s := WriteBufferSizeSettings{WriteBufferSize: 0}
		opts := s.ToServerOptions()
		assert.Empty(t, opts)
	})

	t.Run("positive returns server option", func(t *testing.T) {
		s := WriteBufferSizeSettings{WriteBufferSize: 64 * 1024}
		opts := s.ToServerOptions()
		assert.Len(t, opts, 1)
	})
}
