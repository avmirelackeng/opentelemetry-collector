// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTimeoutSettingsDefaults(t *testing.T) {
	s := NewDefaultTimeoutSettings()
	assert.Equal(t, time.Duration(0), s.DialTimeout)
	assert.Equal(t, time.Duration(0), s.RequestTimeout)
}

func TestTimeoutSettingsValidate(t *testing.T) {
	tests := []struct {
		name    string
		s       TimeoutSettings
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid defaults",
			s:    NewDefaultTimeoutSettings(),
		},
		{
			name: "valid positive values",
			s:    TimeoutSettings{DialTimeout: 5 * time.Second, RequestTimeout: 10 * time.Second},
		},
		{
			name:    "negative dial timeout",
			s:       TimeoutSettings{DialTimeout: -1 * time.Second},
			wantErr: true,
			errMsg:  "dial_timeout must be non-negative",
		},
		{
			name:    "negative request timeout",
			s:       TimeoutSettings{RequestTimeout: -1 * time.Millisecond},
			wantErr: true,
			errMsg:  "request_timeout must be non-negative",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.s.Validate()
			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.errMsg)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestTimeoutSettingsHasDialTimeout(t *testing.T) {
	assert.False(t, (&TimeoutSettings{DialTimeout: 0}).HasDialTimeout())
	assert.True(t, (&TimeoutSettings{DialTimeout: time.Second}).HasDialTimeout())
}

func TestTimeoutSettingsHasRequestTimeout(t *testing.T) {
	assert.False(t, (&TimeoutSettings{RequestTimeout: 0}).HasRequestTimeout())
	assert.True(t, (&TimeoutSettings{RequestTimeout: time.Millisecond}).HasRequestTimeout())
}
