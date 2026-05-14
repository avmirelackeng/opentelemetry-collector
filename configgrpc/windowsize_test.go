// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultWindowSizeSettings(t *testing.T) {
	s := NewDefaultWindowSizeSettings()
	assert.Equal(t, int32(0), s.InitialWindowSize)
	assert.Equal(t, int32(0), s.InitialConnWindowSize)
}

func TestWindowSizeSettingsValidate(t *testing.T) {
	tests := []struct {
		name    string
		s       WindowSizeSettings
		wantErr bool
	}{
		{"defaults", NewDefaultWindowSizeSettings(), false},
		{"positive values", WindowSizeSettings{InitialWindowSize: 65536, InitialConnWindowSize: 131072}, false},
		{"negative window size", WindowSizeSettings{InitialWindowSize: -1}, true},
		{"negative conn window size", WindowSizeSettings{InitialConnWindowSize: -1}, true},
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

func TestWindowSizeSettingsToDialOptions(t *testing.T) {
	t.Run("defaults produce no options", func(t *testing.T) {
		s := NewDefaultWindowSizeSettings()
		opts := s.ToDialOptions()
		assert.Empty(t, opts)
	})

	t.Run("positive values produce options", func(t *testing.T) {
		s := WindowSizeSettings{InitialWindowSize: 65536, InitialConnWindowSize: 131072}
		opts := s.ToDialOptions()
		assert.Len(t, opts, 2)
	})

	t.Run("only window size set", func(t *testing.T) {
		s := WindowSizeSettings{InitialWindowSize: 65536}
		opts := s.ToDialOptions()
		assert.Len(t, opts, 1)
	})
}

func TestWindowSizeSettingsToServerOptions(t *testing.T) {
	t.Run("defaults produce no options", func(t *testing.T) {
		s := NewDefaultWindowSizeSettings()
		opts := s.ToServerOptions()
		assert.Empty(t, opts)
	})

	t.Run("positive values produce options", func(t *testing.T) {
		s := WindowSizeSettings{InitialWindowSize: 65536, InitialConnWindowSize: 131072}
		opts := s.ToServerOptions()
		assert.Len(t, opts, 2)
	})

	t.Run("only conn window size set", func(t *testing.T) {
		s := WindowSizeSettings{InitialConnWindowSize: 131072}
		opts := s.ToServerOptions()
		assert.Len(t, opts, 1)
	})
}
