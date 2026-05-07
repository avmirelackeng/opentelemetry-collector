// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package configgrpc provides utilities for configuring gRPC connections.
//
// RetryPolicy configures the retry behavior for gRPC client calls.
// It follows the gRPC service config retry policy specification:
// https://github.com/grpc/grpc/blob/master/doc/service_config.md
//
// Example usage in a service config:
//
//	retretryPolicy := &RetryPolicy{
//		MaxAttempts:       5,
//		InitialBackoff:    0.5,
//		MaxBackoff:        30.0,
//		BackoffMultiplier: 2.0,
//		RetryableStatusCodes: []string{"UNAVAILABLE", "RESOURCE_EXHAUSTED"},
//	}
//
// Fields:
//   - MaxAttempts: Total attempts including the initial one. Minimum value is 2.
//   - InitialBackoff: Starting backoff in seconds. Must be positive.
//   - MaxBackoff: Upper bound on backoff in seconds. Must be >= InitialBackoff.
//   - BackoffMultiplier: Factor by which backoff increases each retry. Must be positive.
//   - RetryableStatusCodes: gRPC status codes (e.g. "UNAVAILABLE") that trigger a retry.
package configgrpc
