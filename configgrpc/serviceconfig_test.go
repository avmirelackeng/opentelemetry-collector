// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServiceConfigDefault(t *testing.T) {
	sc := NewDefaultServiceConfig()
	assert.Equal(t, ServiceConfig(""), sc)
	assert.False(t, sc.IsSet())
}

func TestServiceConfigValidate(t *testing.T) {
	tests := []struct {
		name    string
		config  ServiceConfig
		wantErr bool
	}{
		{name: "empty", config: "", wantErr: false},
		{name: "valid json", config: `{"loadBalancingPolicy":"round_robin"}`, wantErr: false},
		{name: "invalid json", config: `not-json`, wantErr: true},
		{name: "partial json", config: `{"key":`, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestServiceConfigIsSet(t *testing.T) {
	assert.False(t, ServiceConfig("").IsSet())
	assert.True(t, ServiceConfig(`{}`).IsSet())
}

func TestServiceConfigString(t *testing.T) {
	raw := `{"loadBalancingPolicy":"round_robin"}`
	sc := ServiceConfig(raw)
	assert.Equal(t, raw, sc.String())
}

func TestWithLoadBalancingPolicy(t *testing.T) {
	sc, err := WithLoadBalancingPolicy(BalancerRoundRobin)
	require.NoError(t, err)
	assert.True(t, sc.IsSet())
	require.NoError(t, sc.Validate())
	assert.Contains(t, sc.String(), "round_robin")
}

func TestWithLoadBalancingPolicyInvalid(t *testing.T) {
	_, err := WithLoadBalancingPolicy(BalancerName("invalid-balancer"))
	require.Error(t, err)
}
