// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package configgrpc provides configuration for gRPC clients and servers.
//
// MaxConnectionCountSettings controls how many simultaneous connections
// a gRPC server will accept at once. When the limit is reached, new incoming
// connections are rejected until existing ones are closed.
//
// A value of 0 (the default) disables the limit, allowing an unlimited number
// of concurrent connections.
//
// Example usage:
//
//	settings := configgrpc.NewDefaultMaxConnectionCountSettings()
//	settings.MaxConnectionCount = 1000
//	opt, err := settings.ToServerOption()
package configgrpc
