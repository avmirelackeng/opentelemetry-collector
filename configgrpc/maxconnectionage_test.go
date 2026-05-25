// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultMaxConnectionAgeSettings(t *testing.T) {
	s := NewDefaultMaxConnectionAgeSettings()
	assert.Equal(t, time.Duration(0), s.MaxAge)
	assert.Equal(t, time.Duration(0), s.MaxAgeGrace)
	assert.True(t, s.IsDefault())
}

func TestMaxConnectionAgeSettingsValidate(t *testing.T) {
	tests := []struct {
		name    string
		s       MaxConnectionAgeSettings
		wantErr bool
	}{
		{name: "default", s: NewDefaultMaxConnectionAgeSettings(), wantErr: false},
		{name: "valid max_age", s: MaxConnectionAgeSettings{MaxAge: time.Minute}, wantErr: false},
		{name: "valid max_age and grace", s: MaxConnectionAgeSettings{MaxAge: time.Minute, MaxAgeGrace: 5 * time.Second}, wantErr: false},
		{name: "negative max_age", s: MaxConnectionAgeSettings{MaxAge: -time.Second}, wantErr: true},
		{name: "negative max_age_grace", s: MaxConnectionAgeSettings{MaxAge: time.Minute, MaxAgeGrace: -time.Second}, wantErr: true},
		{name: "grace without max_age", s: MaxConnectionAgeSettings{MaxAgeGrace: 5 * time.Second}, wantErr: true},
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

func TestMaxConnectionAgeSettingsIsDefault(t *testing.T) {
	assert.True(t, NewDefaultMaxConnectionAgeSettings().IsDefault())
	assert.False(t, MaxConnectionAgeSettings{MaxAge: time.Minute}.IsDefault())
	assert.False(t, MaxConnectionAgeSettings{MaxAge: time.Minute, MaxAgeGrace: time.Second}.IsDefault())
}

func TestMaxConnectionAgeSettingsToServerOption(t *testing.T) {
	t.Run("default returns nil", func(t *testing.T) {
		s := NewDefaultMaxConnectionAgeSettings()
		assert.Nil(t, s.ToServerOption())
	})

	t.Run("non-default returns option", func(t *testing.T) {
		s := MaxConnectionAgeSettings{MaxAge: 30 * time.Minute, MaxAgeGrace: 5 * time.Second}
		opt := s.ToServerOption()
		assert.NotNil(t, opt)
	})
}
