// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultHealthCheckSettings(t *testing.T) {
	hc := NewDefaultHealthCheckSettings()
	assert.False(t, hc.Enabled)
	assert.Equal(t, HealthCheckServiceDefault, hc.ServiceName)
	assert.Equal(t, 10*time.Second, hc.Interval)
}

func TestHealthCheckSettingsValidate(t *testing.T) {
	tests := []struct {
		name    string
		settings HealthCheckSettings
		wantErr bool
	}{
		{
			name:     "disabled is always valid",
			settings: HealthCheckSettings{Enabled: false, Interval: 0},
			wantErr:  false,
		},
		{
			name:     "enabled with positive interval is valid",
			settings: HealthCheckSettings{Enabled: true, Interval: 5 * time.Second},
			wantErr:  false,
		},
		{
			name:     "enabled with zero interval is invalid",
			settings: HealthCheckSettings{Enabled: true, Interval: 0},
			wantErr:  true,
		},
		{
			name:     "enabled with negative interval is invalid",
			settings: HealthCheckSettings{Enabled: true, Interval: -1 * time.Second},
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.settings.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestHealthCheckSettingsToHealthCheckRequest(t *testing.T) {
	hc := HealthCheckSettings{
		Enabled:     true,
		ServiceName: "my.test.service",
		Interval:    5 * time.Second,
	}
	req := hc.ToHealthCheckRequest()
	require.NotNil(t, req)
	assert.Equal(t, "my.test.service", req.Service)
}

func TestHealthCheckSettingsToHealthCheckRequestDefaultService(t *testing.T) {
	hc := NewDefaultHealthCheckSettings()
	req := hc.ToHealthCheckRequest()
	require.NotNil(t, req)
	assert.Equal(t, "", req.Service)
}
