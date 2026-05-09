// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package configgrpc provides gRPC configuration types for the
// OpenTelemetry Collector.
//
// # UserAgent
//
// UserAgent represents the value sent as the "user-agent" metadata header in
// gRPC requests. It follows the format described in RFC 7231:
//
//	<product> [ "/" <product-version> ] *( " " <product> [ "/" <product-version> ] )
//
// Example values:
//
//	"otelcol/0.90.0"
//	"myapp/1.2.3 otelcol/0.90.0"
//
// Use [UserAgent.WithProduct] to compose a user-agent string from multiple
// product tokens.
package configgrpc
