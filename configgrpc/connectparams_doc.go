// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package configgrpc provides gRPC configuration types and utilities.
//
// ConnectParams controls how a gRPC client establishes and re-establishes
// connections to a server. It combines a minimum connection timeout with
// an exponential backoff policy for reconnection attempts.
//
// Example configuration (YAML):
//
//	connect_params:
//	  min_connect_timeout: 20s
//	  backoff:
//	    base_delay: 1s
//	    multiplier: 1.6
//	    jitter: 0.2
//	    max_delay: 120s
package configgrpc
