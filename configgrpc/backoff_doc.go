// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package configgrpc provides gRPC client and server configuration.
//
// BackoffConfig controls the exponential backoff strategy used when
// re-establishing gRPC connections after a failure.
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
//	base_delay  - Initial wait before first retry (default: 1s).
//	multiplier  - Growth factor applied to delay after each attempt (default: 1.6).
//	jitter      - Random fraction added to delay to avoid thundering herd (default: 0.2).
//	max_delay   - Maximum delay cap between retries (default: 120s).
package configgrpc
