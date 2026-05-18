// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package configgrpc provides configuration structures and helpers for gRPC
// clients and servers.
//
// TracingSettings controls whether OpenTelemetry trace context is propagated
// on outbound and inbound gRPC calls.
//
// Example configuration (YAML):
//
//	tracing:
//	  mode: enabled
//
// Supported modes:
//   - none    (default): no tracing interceptors are applied.
//   - enabled: tracing interceptors should be wired in by the caller.
package configgrpc
