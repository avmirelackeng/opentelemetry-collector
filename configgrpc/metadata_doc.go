// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package configgrpc provides configuration types for gRPC connections.
//
// # Metadata
//
// The Metadata type represents gRPC metadata (also known as headers) that can
// be attached to outgoing or incoming gRPC calls.
//
// Metadata keys are case-insensitive strings consisting of printable ASCII
// characters. Keys ending in "-bin" are treated as binary values and will be
// base64-encoded by the gRPC library.
//
// Example configuration:
//
//	metadata:
//	  authorization:
//	    - "Bearer my-token"
//	  x-custom-header:
//	    - "value1"
//	    - "value2"
package configgrpc
