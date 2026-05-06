// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package configgrpc provides gRPC configuration utilities for
// OpenTelemetry Collector components.
//
// # Headers
//
// The Headers type represents gRPC metadata headers that can be attached
// to outgoing gRPC requests. Header keys must be non-empty strings
// containing only lowercase ASCII letters, digits, hyphens, underscores,
// and dots.
//
// Example configuration:
//
//	headers:
//	  x-tenant-id: "my-tenant"
//	  authorization: "Bearer token123"
//	  x-custom-header: "value"
//
// Header keys are case-insensitive per the HTTP/2 and gRPC specifications;
// they will be normalized to lowercase when transmitted.
package configgrpc
