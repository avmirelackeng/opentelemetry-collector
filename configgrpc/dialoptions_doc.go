// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package configgrpc provides gRPC client and server configuration.
//
// DialOptionsSettings allows callers to specify a set of well-known gRPC dial
// options by name, avoiding the need to import gRPC internals directly.
//
// Supported options:
//
//	"block"     - grpc.WithBlock: makes Dial block until the connection is ready.
//	"insecure"  - grpc.WithInsecure: disables TLS (use only in trusted environments).
//	"no_proxy"  - grpc.WithNoProxy: bypasses any configured HTTP proxy.
//
// Example configuration (YAML):
//
//	dial_options:
//	  options:
//	    - block
//	    - no_proxy
package configgrpc
