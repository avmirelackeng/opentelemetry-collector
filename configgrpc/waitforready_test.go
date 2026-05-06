// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWaitForReadyConstants(t *testing.T) {
	// Verify the string representations of WaitForReady constants
	tests := []struct {
		name     string
		value    WaitForReady
		expected string
	}{
		{"WaitForReadyEnabled", WaitForReadyEnabled, "true"},
		{"WaitForReadyDisabled", WaitForReadyDisabled, "false"},
		{"WaitForReadyUnset", WaitForReadyUnset, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.value.String())
		})
	}
}

func TestParseWaitForReady(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    WaitForReady
		expectError bool
	}{
		{"true", "true", WaitForReadyEnabled, false},
		{"false", "false", WaitForReadyDisabled, false},
		{"empty", "", WaitForReadyUnset, false},
		{"TRUE uppercase", "TRUE", WaitForReadyEnabled, false},
		{"FALSE uppercase", "FALSE", WaitForReadyDisabled, false},
		{"True mixed case", "True", WaitForReadyEnabled, false},
		{"invalid", "invalid", WaitForReadyUnset, true},
		{"1", "1", WaitForReadyEnabled, false},
		{"0", "0", WaitForReadyDisabled, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseWaitForReady(tt.input)
			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestWaitForReadyValidate(t *testing.T) {
	tests := []struct {
		name        string
		value       WaitForReady
		expectError bool
	}{
		{"enabled is valid", WaitForReadyEnabled, false},
		{"disabled is valid", WaitForReadyDisabled, false},
		{"unset is valid", WaitForReadyUnset, false},
		{"invalid value", WaitForReady(99), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.value.Validate()
			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestWaitForReadyIsEnabled(t *testing.T) {
	assert.True(t, WaitForReadyEnabled.IsEnabled())
	assert.False(t, WaitForReadyDisabled.IsEnabled())
	assert.False(t, WaitForReadyUnset.IsEnabled())
}

func TestWaitForReadyMarshalText(t *testing.T) {
	tests := []struct {
		name     string
		value    WaitForReady
		expected string
	}{
		{"enabled", WaitForReadyEnabled, "true"},
		{"disabled", WaitForReadyDisabled, "false"},
		{"unset", WaitForReadyUnset, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			text, err := tt.value.MarshalText()
			require.NoError(t, err)
			assert.Equal(t, tt.expected, string(text))
		})
	}
}

func TestWaitForReadyUnmarshalText(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    WaitForReady
		expectError bool
	}{
		{"true", "true", WaitForReadyEnabled, false},
		{"false", "false", WaitForReadyDisabled, false},
		{"empty", "", WaitForReadyUnset, false},
		{"invalid", "maybe", WaitForReadyUnset, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var w WaitForReady
			err := w.UnmarshalText([]byte(tt.input))
			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, w)
			}
		})
	}
}
