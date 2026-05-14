// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMaxMessageSizeSettingsDefaults(t *testing.T) {
	s := NewDefaultMaxMessageSizeSettings()
	assert.Equal(t, DefaultMaxRecvMsgSizeMiB, s.RecvMsgSizeMiB)
	assert.Equal(t, DefaultMaxSendMsgSizeMiB, s.SendMsgSizeMiB)
}

func TestMaxMessageSizeSettingsValidate(t *testing.T) {
	tests := []struct {
		name    string
		s       MaxMessageSizeSettings
		wantErr bool
	}{
		{
			name:    "valid defaults",
			s:       NewDefaultMaxMessageSizeSettings(),
			wantErr: false,
		},
		{
			name:    "valid zeros",
			s:       MaxMessageSizeSettings{RecvMsgSizeMiB: 0, SendMsgSizeMiB: 0},
			wantErr: false,
		},
		{
			name:    "valid large values",
			s:       MaxMessageSizeSettings{RecvMsgSizeMiB: 128, SendMsgSizeMiB: 64},
			wantErr: false,
		},
		{
			name:    "negative recv",
			s:       MaxMessageSizeSettings{RecvMsgSizeMiB: -1, SendMsgSizeMiB: 0},
			wantErr: true,
		},
		{
			name:    "negative send",
			s:       MaxMessageSizeSettings{RecvMsgSizeMiB: 0, SendMsgSizeMiB: -1},
			wantErr: true,
		},
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

func TestMaxMessageSizeSettingsToDialOptions(t *testing.T) {
	t.Run("zeros produce no options", func(t *testing.T) {
		s := MaxMessageSizeSettings{RecvMsgSizeMiB: 0, SendMsgSizeMiB: 0}
		opts := s.ToDialOptions()
		assert.Empty(t, opts)
	})

	t.Run("non-zero values produce options", func(t *testing.T) {
		s := MaxMessageSizeSettings{RecvMsgSizeMiB: 4, SendMsgSizeMiB: 8}
		opts := s.ToDialOptions()
		assert.Len(t, opts, 2)
	})

	t.Run("only recv set", func(t *testing.T) {
		s := MaxMessageSizeSettings{RecvMsgSizeMiB: 4, SendMsgSizeMiB: 0}
		opts := s.ToDialOptions()
		assert.Len(t, opts, 1)
	})
}

func TestMaxMessageSizeSettingsToServerOptions(t *testing.T) {
	t.Run("zeros produce no options", func(t *testing.T) {
		s := MaxMessageSizeSettings{RecvMsgSizeMiB: 0, SendMsgSizeMiB: 0}
		opts := s.ToServerOptions()
		assert.Empty(t, opts)
	})

	t.Run("non-zero values produce options", func(t *testing.T) {
		s := MaxMessageSizeSettings{RecvMsgSizeMiB: 4, SendMsgSizeMiB: 8}
		opts := s.ToServerOptions()
		assert.Len(t, opts, 2)
	})

	t.Run("only send set", func(t *testing.T) {
		s := MaxMessageSizeSettings{RecvMsgSizeMiB: 0, SendMsgSizeMiB: 8}
		opts := s.ToServerOptions()
		assert.Len(t, opts, 1)
	})
}
