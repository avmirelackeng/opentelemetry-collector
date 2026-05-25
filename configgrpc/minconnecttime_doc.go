// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package configgrpc provides configuration structures and helpers for
// configuring gRPC clients and servers.
//
// MinConnectTimeSettings controls the minimum time a connection attempt
// is given to complete before the connection is considered failed.
//
// Example usage:
//
//	settings := configgrpc.NewDefaultMinConnectTimeSettings()
//	settings.MinConnectTimeout = 30 * time.Second
//	if err := settings.Validate(); err != nil {
//		// handle error
//	}
//	opt := settings.ToDialOption()
package configgrpc
