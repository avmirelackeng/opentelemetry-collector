// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultSendRecvBufferSizeSettings(t *testing.T) {
	s := NewDefaultSendRecvBufferSizeSettings()
	assert.Equal(t, 0, s.ReadBufferSize)
	assert.Equal(t, 0, s.WriteBufferSize)
}

func TestSendRecvBufferSizeSettingsValidate(t *testing.T) {
	tests := []struct {
		name    string
		s       SendRecvBufferSizeSettings
		wantErr bool
	}{
		{name: "defaults", s: NewDefaultSendRecvBufferSizeSettings(), wantErr: false},
		{name: "positive values", s: SendRecvBufferSizeSettings{ReadBufferSize: 1024, WriteBufferSize: 2048}, wantErr: false},
		{name: "negative read", s: SendRecvBufferSizeSettings{ReadBufferSize: -1}, wantErr: true},
		{name: "negative write", s: SendRecvBufferSizeSettings{WriteBufferSize: -1}, wantErr: true},
		{name: "both negative", s: SendRecvBufferSizeSettings{ReadBufferSize: -10, WriteBufferSize: -10}, wantErr: true},
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

func TestSendRecvBufferSizeSettingsToDialOptions(t *testing.T) {
	t.Run("defaults produce no options", func(t *testing.T) {
		s := NewDefaultSendRecvBufferSizeSettings()
		opts := s.ToDialOptions()
		assert.Empty(t, opts)
	})

	t.Run("positive values produce options", func(t *testing.T) {
		s := SendRecvBufferSizeSettings{ReadBufferSize: 4096, WriteBufferSize: 8192}
		opts := s.ToDialOptions()
		assert.Len(t, opts, 2)
	})

	t.Run("only read buffer set", func(t *testing.T) {
		s := SendRecvBufferSizeSettings{ReadBufferSize: 4096}
		opts := s.ToDialOptions()
		assert.Len(t, opts, 1)
	})
}

func TestSendRecvBufferSizeSettingsToServerOptions(t *testing.T) {
	t.Run("defaults produce no options", func(t *testing.T) {
		s := NewDefaultSendRecvBufferSizeSettings()
		opts := s.ToServerOptions()
		assert.Empty(t, opts)
	})

	t.Run("positive values produce options", func(t *testing.T) {
		s := SendRecvBufferSizeSettings{ReadBufferSize: 4096, WriteBufferSize: 8192}
		opts := s.ToServerOptions()
		assert.Len(t, opts, 2)
	})

	t.Run("only write buffer set", func(t *testing.T) {
		s := SendRecvBufferSizeSettings{WriteBufferSize: 8192}
		opts := s.ToServerOptions()
		assert.Len(t, opts, 1)
	})
}
