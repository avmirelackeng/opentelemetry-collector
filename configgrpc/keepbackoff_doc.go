// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package configgrpc provides gRPC configuration types and helpers.
//
// KeepaliveBackoffConfig controls the exponential backoff strategy used when
// a gRPC client attempts to reconnect after a keepalive failure or connection
// drop. The parameters mirror those of keepalive.BackoffConfig from the
// google.golang.org/grpc/keepalive package.
//
// Example configuration (YAML):
//
//	keepalive_backoff:
//	  base_delay: 1s
//	  multiplier: 1.6
//	  jitter: 0.2
//	  max_delay: 120s
package configgrpc
