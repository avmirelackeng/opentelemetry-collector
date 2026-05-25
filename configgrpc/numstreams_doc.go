// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package configgrpc provides gRPC configuration for OpenTelemetry Collector.
//
// NumStreamsSettings controls the maximum number of concurrent streams
// allowed per connection for both clients and servers.
//
// Example usage:
//
//	settings := configgrpc.NewDefaultNumStreamsSettings()
//	settings.MaxStreamsPerConn = 100
//	if err := settings.Validate(); err != nil {
//		// handle error
//	}
//	opt := settings.ToServerOption()
//
package configgrpc
