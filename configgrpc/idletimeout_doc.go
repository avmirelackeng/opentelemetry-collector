// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package configgrpc provides gRPC configuration types and helpers.
//
// IdleTimeoutSettings controls how long an idle connection is kept open before
// being closed. An idle connection is one that has no active RPCs.
//
// Example usage:
//
//	settings := configgrpc.NewDefaultIdleTimeoutSettings()
//	settings.Timeout = 30 * time.Second
//	if err := settings.Validate(); err != nil {
//		// handle error
//	}
//	// Use as a dial option:
//	opt := settings.ToDialOption()
//	// Use as a server option:
//	srvOpt := settings.ToServerOption()
package configgrpc
