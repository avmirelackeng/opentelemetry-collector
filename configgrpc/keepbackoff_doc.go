// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package configgrpc provides configuration structures and helpers for
// configuring gRPC clients and servers in OpenTelemetry Collector components.
//
// # Keepalive Backoff Configuration
//
// KeepaliveBackoffConfig controls the backoff strategy used when a keepalive
// ping fails or a connection is lost. It wraps gRPC's backoff.Config and
// allows tuning the exponential backoff parameters used during reconnection.
//
// Example configuration:
//
//	backoff:
//	  base_delay: 1s
//	  multiplier: 1.6
//	  jitter: 0.2
//	  max_delay: 120s
//
// Fields:
//
//   - BaseDelay: The initial delay before the first reconnection attempt.
//     Defaults to 1 second.
//
//   - Multiplier: The factor by which the delay is multiplied after each
//     failed attempt. Defaults to 1.6.
//
//   - Jitter: The fraction of the delay to randomize. A value of 0.2 means
//     the delay is randomized by ±20%. Must be in [0.0, 1.0]. Defaults to 0.2.
//
//   - MaxDelay: The maximum delay between reconnection attempts.
//     Defaults to 120 seconds.
package configgrpc
