// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultMaxConnectionAgeGraceSettings(t *testing.T) {
	s := NewDefaultMaxConnectionAgeGraceSettings()
	assert.False(t, s.Enabled)
	assert.Equal(t, time.Duration(0), s.Grace)
}

func TestMaxConnectionAgeGraceSettingsValidate(t *testing.T) {
	tests := []struct {
		name    string
		s       MaxConnectionAgeGraceSettings
		wantErr bool
	}{
		{
			name:    "default",
			s:       NewDefaultMaxConnectionAgeGraceSettings(),
			wantErr: false,
		},
		{
			name:    "enabled with positive grace",
			s:       MaxConnectionAgeGraceSettings{Enabled: true, Grace: 5 * time.Second},
			wantErr: false,
		},
		{
			name:    "enabled with zero grace",
			s:       MaxConnectionAgeGraceSettings{Enabled: true, Grace: 0},
			wantErr: false,
		},
		{
			name:    "enabled with negative grace",
			s:       MaxConnectionAgeGraceSettings{Enabled: true, Grace: -1 * time.Second},
			wantErr: true,
		},
		{
			name:    "disabled with negative grace",
			s:       MaxConnectionAgeGraceSettings{Enabled: false, Grace: -1 * time.Second},
			wantErr: false,
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

func TestMaxConnectionAgeGraceSettingsIsDefault(t *testing.T) {
	s := NewDefaultMaxConnectionAgeGraceSettings()
	assert.True(t, s.IsDefault())

	s.Enabled = true
	assert.False(t, s.IsDefault())
}

func TestMaxConnectionAgeGraceSettingsToServerOption(t *testing.T) {
	s := NewDefaultMaxConnectionAgeGraceSettings()
	assert.Nil(t, s.ToServerOption())

	s.Enabled = true
	s.Grace = 10 * time.Second
	opt := s.ToServerOption()
	assert.NotNil(t, opt)
}
