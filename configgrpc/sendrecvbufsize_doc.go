// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package configgrpc provides gRPC configuration types and utilities.
//
// # Send/Receive Buffer Size Settings
//
// SendRecvBufferSizeSettings controls the TCP-level read and write buffer
// sizes used by gRPC connections. Both client and server sides can be
// configured independently.
//
// Example configuration:
//
//	send_recv_buffer_size:
//	  read_buffer_size: 32768   # 32 KiB
//	  write_buffer_size: 32768  # 32 KiB
//
// A value of 0 (the default) means the system default will be used.
// Tuning these values can improve throughput on high-latency or
// high-bandwidth network paths.
package configgrpc
