// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package configgrpc provides gRPC configuration utilities.
//
// CodecType specifies the encoding codec used for gRPC messages.
//
// Supported values:
//
//   - "proto" (default): Protocol Buffers encoding. This is the standard
//     gRPC encoding and is recommended for production use.
//
//   - "json": JSON encoding. Useful for debugging or interoperability
//     with systems that do not support protobuf.
//
//   - "raw": Raw bytes, no encoding transformation applied. Useful
//     when the message payload is already serialized.
//
// Example configuration (YAML):
//
//	codec:
//	  type: proto
package configgrpc
