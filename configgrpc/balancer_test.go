// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBalancerNameConstants(t *testing.T) {
	assert.Equal(t, BalancerName("round_robin"), BalancerRoundRobin)
	assert.Equal(t, BalancerName("pick_first"), BalancerPickFirst)
}

func TestBalancerNameValidate(t *testing.T) {
	tests := []struct {
		name    string
		balancer BalancerName
		wantErr bool
	}{
		{name: "empty is valid", balancer: "", wantErr: false},
		{name: "round_robin is valid", balancer: BalancerRoundRobin, wantErr: false},
		{name: "pick_first is valid", balancer: BalancerPickFirst, wantErr: false},
		{name: "unknown balancer", balancer: "least_conn", wantErr: true},
		{name: "random string", balancer: "foobar", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.balancer.Validate()
			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), "unknown balancer name")
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestBalancerNameString(t *testing.T) {
	assert.Equal(t, "round_robin", BalancerRoundRobin.String())
	assert.Equal(t, "pick_first", BalancerPickFirst.String())
	assert.Equal(t, "", BalancerName("").String())
}

func TestBalancerNameServiceConfig(t *testing.T) {
	tests := []struct {
		name     string
		balancer BalancerName
		want     string
	}{
		{
			name:     "empty returns empty string",
			balancer: "",
			want:     "",
		},
		{
			name:     "round_robin returns service config",
			balancer: BalancerRoundRobin,
			want:     `{"loadBalancingPolicy": "round_robin"}`,
		},
		{
			name:     "pick_first returns service config",
			balancer: BalancerPickFirst,
			want:     `{"loadBalancingPolicy": "pick_first"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.balancer.ServiceConfig()
			assert.Equal(t, tt.want, got)
		})
	}
}
