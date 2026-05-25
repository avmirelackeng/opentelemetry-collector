// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultMaxRecvMsgSizeSettings(t *testing.T) {
	s := NewDefaultMaxRecvMsgSizeSettings()
	assert.Equal(t, 0, s.MaxRecvMsgSizeBytes)
	assert.True(t, s.IsDefault())
}

func TestMaxRecvMsgSizeSettingsValidate(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantErr bool
	}{
		{name: "zero is valid", size: 0, wantErr: false},
		{name: "positive is valid", size: 1024 * 1024, wantErr: false},
		{name: "large value is valid", size: 64 * 1024 * 1024, wantErr: false},
		{name: "negative is invalid", size: -1, wantErr: true},
		{name: "large negative is invalid", size: -1024, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := MaxRecvMsgSizeSettings{MaxRecvMsgSizeBytes: tt.size}
			err := s.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMaxRecvMsgSizeSettingsIsDefault(t *testing.T) {
	assert.True(t, MaxRecvMsgSizeSettings{MaxRecvMsgSizeBytes: 0}.IsDefault())
	assert.False(t, MaxRecvMsgSizeSettings{MaxRecvMsgSizeBytes: 1}.IsDefault())
}

func TestMaxRecvMsgSizeSettingsToDialOption(t *testing.T) {
	s := NewDefaultMaxRecvMsgSizeSettings()
	opt := s.ToDialOption()
	assert.Nil(t, opt, "default settings should return nil dial option")

	s.MaxRecvMsgSizeBytes = 4 * 1024 * 1024
	opt = s.ToDialOption()
	assert.NotNil(t, opt, "non-default settings should return a dial option")
}

func TestMaxRecvMsgSizeSettingsToServerOption(t *testing.T) {
	s := NewDefaultMaxRecvMsgSizeSettings()
	opt := s.ToServerOption()
	assert.Nil(t, opt, "default settings should return nil server option")

	s.MaxRecvMsgSizeBytes = 8 * 1024 * 1024
	opt = s.ToServerOption()
	assert.NotNil(t, opt, "non-default settings should return a server option")
}
