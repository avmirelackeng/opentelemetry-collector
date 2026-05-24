// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package configgrpc provides gRPC configuration utilities.
//
// # Rate Limit Settings
//
// RateLimitSettings controls server-side rate limiting for gRPC connections
// and streams. It provides two independent controls:
//
//   - MaxConcurrentStreams: limits the total number of concurrent RPC streams
//     that a single client connection may have open at any given time.
//     Defaults to 0 (unlimited).
//
//   - MaxConnectionsPerSecond: limits the rate at which new client connections
//     are accepted by the server. Defaults to 0 (unlimited).
//
// Example configuration:
//
//	rate_limit:
//	  max_concurrent_streams: 100
//	  max_connections_per_second: 50
package configgrpc
