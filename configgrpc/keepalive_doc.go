// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package configgrpc provides gRPC configuration utilities for the
// OpenTelemetry Collector.
//
// # Keepalive Configuration
//
// KeepaliveClientParameters controls how a gRPC client sends keepalive pings
// to the server. This is useful to detect broken connections and to keep
// connections alive through proxies and firewalls that may close idle connections.
//
// Example client keepalive configuration:
//
//	params := configgrpc.KeepaliveClientParameters{
//		Time:                10 * time.Second,
//		Timeout:             5 * time.Second,
//		PermitWithoutStream: true,
//	}
//
// KeepaliveServerParameters controls how a gRPC server manages keepalive
// behavior for connected clients.
//
// Example server keepalive configuration:
//
//	params := configgrpc.KeepaliveServerParameters{
//		MaxConnectionIdle:     15 * time.Minute,
//		MaxConnectionAge:      30 * time.Minute,
//		MaxConnectionAgeGrace: 5 * time.Second,
//		Time:                  5 * time.Second,
//		Timeout:               1 * time.Second,
//	}
package configgrpc // import "go.opentelemetry.io/collector/configgrpc"
