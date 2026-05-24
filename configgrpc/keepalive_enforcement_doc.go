// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package configgrpc provides configuration structures and utilities for
// configuring gRPC clients and servers in the OpenTelemetry Collector.
//
// # Keepalive Enforcement Policy
//
// KeepaliveEnforcementPolicy controls server-side keepalive enforcement.
// It allows the server to enforce minimum keepalive intervals for clients
// and optionally allow keepalive pings even when there are no active streams.
//
// Example configuration:
//
//	enforcementPolicy:
//	  minTime: 5m
//	  permitWithoutStream: false
//
// Fields:
//
//	- MinTime: The minimum amount of time a client should wait before sending
//	  a keepalive ping. If a client sends pings more frequently, the server
//	  will close the connection. Defaults to 5 minutes.
//
//	- PermitWithoutStream: If true, the server allows keepalive pings even
//	  when there are no active streams (RPCs). If false (the default), the
//	  server will close connections that send pings without active streams.
package configgrpc
