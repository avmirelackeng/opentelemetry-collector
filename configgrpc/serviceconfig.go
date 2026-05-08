// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"encoding/json"
	"errors"
	"fmt"
)

// ServiceConfig represents a gRPC service configuration in JSON format.
// See https://github.com/grpc/grpc/blob/master/doc/service_config.md for details.
type ServiceConfig string

// NewDefaultServiceConfig returns a ServiceConfig with sensible defaults.
func NewDefaultServiceConfig() ServiceConfig {
	return ""
}

// Validate checks that the ServiceConfig is either empty or valid JSON.
func (s ServiceConfig) Validate() error {
	if s == "" {
		return nil
	}
	var raw map[string]any
	if err := json.Unmarshal([]byte(s), &raw); err != nil {
		return fmt.Errorf("service config must be valid JSON: %w", err)
	}
	return nil
}

// IsSet returns true if the ServiceConfig is non-empty.
func (s ServiceConfig) IsSet() bool {
	return s != ""
}

// String returns the string representation of the ServiceConfig.
func (s ServiceConfig) String() string {
	return string(s)
}

// WithLoadBalancingPolicy returns a ServiceConfig JSON string that sets
// the specified load balancing policy name.
func WithLoadBalancingPolicy(policy BalancerName) (ServiceConfig, error) {
	if err := policy.Validate(); err != nil {
		return "", fmt.Errorf("invalid balancer policy: %w", err)
	}
	type lbConfig struct {
		LoadBalancingConfig []map[string]any `json:"loadBalancingConfig"`
	}
	cfg := lbConfig{
		LoadBalancingConfig: []map[string]any{
			{string(policy): map[string]any{}},
		},
	}
	b, err := json.Marshal(cfg)
	if err != nil {
		return "", errors.New("failed to marshal service config")
	}
	return ServiceConfig(b), nil
}
