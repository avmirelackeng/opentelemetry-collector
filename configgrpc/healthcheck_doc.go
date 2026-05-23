// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package configgrpc provides gRPC configuration types and helpers.
//
// HealthCheckSettings controls the gRPC health checking protocol
// (https://github.com/grpc/grpc/blob/master/doc/health-checking.md).
//
// Example usage:
//
//	hc := configgrpc.NewDefaultHealthCheckSettings()
//	hc.Enabled = true
//	hc.ServiceName = "my.service"
//	hc.Interval = 5 * time.Second
//	if err := hc.Validate(); err != nil {
//		// handle error
//	}
//	req := hc.ToHealthCheckRequest()
package configgrpc
