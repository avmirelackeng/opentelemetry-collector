// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package configgrpc provides gRPC configuration types and helpers.
//
// TimeoutSettings controls how long the gRPC client waits for various
// operations to complete:
//
//	- DialTimeout: maximum time to wait when establishing a new connection.
//	  A value of 0 means no timeout is applied.
//
//	- RequestTimeout: maximum time to wait for a single RPC call to complete.
//	  A value of 0 means no timeout is applied.
//
// Example usage:
//
//	settings := configgrpc.NewDefaultTimeoutSettings()
//	settings.DialTimeout = 5 * time.Second
//	settings.RequestTimeout = 30 * time.Second
//	if err := settings.Validate(); err != nil {
//		// handle error
//	}
package configgrpc
