// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc // import "go.opentelemetry.io/collector/configgrpc"

import (
	"fmt"
)

// BalancerName represents a gRPC client-side load balancing policy name.
type BalancerName string

const (
	// BalancerRoundRobin is the round-robin load balancing policy.
	BalancerRoundRobin BalancerName = "round_robin"

	// BalancerPickFirst is the pick-first load balancing policy (default gRPC behavior).
	BalancerPickFirst BalancerName = "pick_first"
)

// knownBalancers contains the set of valid balancer names.
var knownBalancers = map[BalancerName]struct{}{
	BalancerRoundRobin: {},
	BalancerPickFirst: {},
}

// Validate checks whether the BalancerName is a known/supported value.
// An empty BalancerName is considered valid (uses gRPC default).
func (b BalancerName) Validate() error {
	if b == "" {
		return nil
	}
	if _, ok := knownBalancers[b]; !ok {
		return fmt.Errorf("unknown balancer name %q: must be one of [round_robin, pick_first]", b)
	}
	return nil
}

// String returns the string representation of the BalancerName.
func (b BalancerName) String() string {
	return string(b)
}

// ServiceConfig returns a gRPC service config JSON snippet that sets this
// balancer as the load balancing policy. Returns an empty string when the
// BalancerName is empty (no override).
func (b BalancerName) ServiceConfig() string {
	if b == "" {
		return ""
	}
	return fmt.Sprintf(`{"loadBalancingPolicy": %q}`, string(b))
}
