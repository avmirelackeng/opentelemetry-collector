// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package configgrpc provides gRPC configuration utilities.
//
// WriteBufferSizeSettings controls the size of the write buffer used by
// gRPC connections. Tuning the write buffer size can improve throughput
// for workloads that send large messages or many small messages in bursts.
//
// Example usage:
//
//	settings := configgrpc.NewDefaultWriteBufferSizeSettings()
//	settings.WriteBufferSize = 64 * 1024 // 64 KiB
//	if err := settings.Validate(); err != nil {
//		// handle error
//	}
//	dialOpts := settings.ToDialOptions()
//	serverOpts := settings.ToServerOptions()
package configgrpc
