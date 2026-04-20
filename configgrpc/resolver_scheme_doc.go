// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package configgrpc provides utilities for configuring gRPC connections
// in OpenTelemetry Collector components.
//
// # Resolver Scheme
//
// The ResolverScheme type controls which gRPC name resolver is used when
// establishing a gRPC connection. gRPC supports multiple resolver schemes
// that determine how the endpoint address is resolved to one or more
// backend addresses.
//
// Supported schemes:
//
//   - "dns" (default): Uses DNS resolution. The endpoint is resolved via
//     standard DNS lookups. Suitable for most production use cases.
//
//   - "passthrough": Passes the address directly to the dialer without
//     any name resolution. Useful when connecting to a single known
//     IP address or when using a proxy.
//
//   - "ipv4": Resolves the address as an IPv4 address directly.
//
//   - "ipv6": Resolves the address as an IPv6 address directly.
//
// Example configuration (YAML):
//
//	 exporters:
//	   otlp:
//	     endpoint: my-collector:4317
//	     balancer_name: round_robin
//	     resolver:
//	       scheme: dns
//
// When using the "dns" scheme with a load balancer, gRPC will resolve the
// DNS name to multiple IP addresses and distribute requests across them.
// This is the recommended approach for connecting to a pool of collectors.
package configgrpc
