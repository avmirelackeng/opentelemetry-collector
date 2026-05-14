// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package configgrpc provides gRPC configuration types and helpers.
//
// WindowSizeSettings controls the flow-control window sizes used by gRPC
// for streams and connections. Tuning these values can improve throughput
// for high-volume telemetry pipelines.
//
// Example configuration (YAML):
//
//	window_size:
//	  initial_window_size: 65536       # 64 KiB per stream
//	  initial_conn_window_size: 131072  # 128 KiB per connection
//
// A value of 0 (the default) leaves the gRPC library default in place.
package configgrpc
