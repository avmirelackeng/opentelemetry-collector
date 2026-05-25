// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package configgrpc provides gRPC configuration types and helpers for
// OpenTelemetry Collector components.
//
// MaxConnectionIdleSettings controls how long a server-side connection may
// remain idle before the server sends a GoAway frame and closes it. Setting
// MaxConnectionIdle to zero (the default) disables the limit.
//
// Example usage:
//
//	settings := configgrpc.NewDefaultMaxConnectionIdleSettings()
//	settings.MaxConnectionIdle = 30 * time.Minute
//	if err := settings.Validate(); err != nil {
//		// handle error
//	}
//	params := settings.ToServerOption()
//	server := grpc.NewServer(grpc.KeepaliveParams(params))
package configgrpc
