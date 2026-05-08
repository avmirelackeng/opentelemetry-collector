// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package configgrpc provides configuration structures and helpers for gRPC
// connections used by OpenTelemetry Collector components.
//
// InterceptorSettings configures gRPC interceptors for unary and streaming RPCs.
//
// Example configuration:
//
//	interceptor:
//	  type: unary
//	  enabled: true
//
// Valid interceptor types:
//   - unary: applies to unary (request/response) RPCs
//   - stream: applies to streaming RPCs
package configgrpc
