// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultInitialWindowSizeSettings(t *testing.T) {
	s := NewDefaultInitialWindowSizeSettings()
	assert.Equal(t, int32(0), s.InitialWindowSize)
	assert.True(t, s.IsDefault())
}

func TestInitialWindowSizeSettingsValidate(t *testing.T) {
	tests := []struct {
		name    string
		size    int32
		wantErr bool
	}{
		{name: "default", size: 0, wantErr: false},
		{name: "valid small", size: 65536, wantErr: false},
		{name: "valid large", size: 1 << 20, wantErr: false},
		{name: "max allowed", size: 1 << 30, wantErr: false},
		{name: "negative", size: -1, wantErr: true},
		{name: "exceeds max", size: (1 << 30) + 1, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := InitialWindowSizeSettings{InitialWindowSize: tt.size}
			err := s.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestInitialWindowSizeSettingsIsDefault(t *testing.T) {
	s := InitialWindowSizeSettings{InitialWindowSize: 0}
	assert.True(t, s.IsDefault())

	s.InitialWindowSize = 65536
	assert.False(t, s.IsDefault())
}

func TestInitialWindowSizeSettingsToDialOption(t *testing.T) {
	s := NewDefaultInitialWindowSizeSettings()
	opt := s.ToDialOption()
	assert.Nil(t, opt, "default settings should return nil dial option")

	s.InitialWindowSize = 1 << 20
	opt = s.ToDialOption()
	assert.NotNil(t, opt, "non-default settings should return a dial option")
}

func TestInitialWindowSizeSettingsToServerOption(t *testing.T) {
	s := NewDefaultInitialWindowSizeSettings()
	opt := s.ToServerOption()
	assert.Nil(t, opt, "default settings should return nil server option")

	s.InitialWindowSize = 1 << 20
	opt = s.ToServerOption()
	assert.NotNil(t, opt, "non-default settings should return a server option")
}
