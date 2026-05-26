// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultMaxConnectionAgeGraceMillisSettings(t *testing.T) {
	s := NewDefaultMaxConnectionAgeGraceMillisSettings()
	assert.Equal(t, uint32(0), s.GraceMillis)
	assert.True(t, s.IsDefault())
}

func TestMaxConnectionAgeGraceMillisSettingsValidate(t *testing.T) {
	tests := []struct {
		name    string
		ms      uint32
		wantErr bool
	}{
		{name: "zero is valid", ms: 0, wantErr: false},
		{name: "small value is valid", ms: 500, wantErr: false},
		{name: "one hour is valid", ms: 3_600_000, wantErr: false},
		{name: "over one hour is invalid", ms: 3_600_001, wantErr: true},
		{name: "large value is invalid", ms: 10_000_000, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := MaxConnectionAgeGraceMillisSettings{GraceMillis: tt.ms}
			err := s.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMaxConnectionAgeGraceMillisSettingsIsDefault(t *testing.T) {
	assert.True(t, MaxConnectionAgeGraceMillisSettings{GraceMillis: 0}.IsDefault())
	assert.False(t, MaxConnectionAgeGraceMillisSettings{GraceMillis: 1}.IsDefault())
	assert.False(t, MaxConnectionAgeGraceMillisSettings{GraceMillis: 5000}.IsDefault())
}

func TestMaxConnectionAgeGraceMillisSettingsToServerOption(t *testing.T) {
	t.Run("default returns nil", func(t *testing.T) {
		s := NewDefaultMaxConnectionAgeGraceMillisSettings()
		opt := s.ToServerOption()
		assert.Nil(t, opt)
	})

	t.Run("non-default returns server option", func(t *testing.T) {
		s := MaxConnectionAgeGraceMillisSettings{GraceMillis: 2000}
		opt := s.ToServerOption()
		assert.NotNil(t, opt)
	})
}
