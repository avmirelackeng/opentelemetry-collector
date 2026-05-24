// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultKeepaliveEnforcementPolicy(t *testing.T) {
	policy := NewDefaultKeepaliveEnforcementPolicy()
	assert.Equal(t, 5*time.Minute, policy.MinTime)
	assert.False(t, policy.PermitWithoutStream)
}

func TestKeepaliveEnforcementPolicyValidate(t *testing.T) {
	tests := []struct {
		name    string
		policy  KeepaliveEnforcementPolicy
		wantErr bool
	}{
		{
			name:    "default is valid",
			policy:  NewDefaultKeepaliveEnforcementPolicy(),
			wantErr: false,
		},
		{
			name: "zero min_time is valid",
			policy: KeepaliveEnforcementPolicy{
				MinTime:             0,
				PermitWithoutStream: false,
			},
			wantErr: false,
		},
		{
			name: "negative min_time is invalid",
			policy: KeepaliveEnforcementPolicy{
				MinTime: -1 * time.Second,
			},
			wantErr: true,
		},
		{
			name: "permit without stream enabled",
			policy: KeepaliveEnforcementPolicy{
				MinTime:             1 * time.Minute,
				PermitWithoutStream: true,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.policy.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestKeepaliveEnforcementPolicyToEnforcementPolicy(t *testing.T) {
	policy := KeepaliveEnforcementPolicy{
		MinTime:             2 * time.Minute,
		PermitWithoutStream: true,
	}

	ep := policy.ToEnforcementPolicy()
	assert.Equal(t, 2*time.Minute, ep.MinTime)
	assert.True(t, ep.PermitWithoutStream)
}

func TestKeepaliveEnforcementPolicyDefaultToEnforcementPolicy(t *testing.T) {
	policy := NewDefaultKeepaliveEnforcementPolicy()
	ep := policy.ToEnforcementPolicy()
	assert.Equal(t, 5*time.Minute, ep.MinTime)
	assert.False(t, ep.PermitWithoutStream)
}
