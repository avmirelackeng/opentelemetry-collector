// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package configgrpc provides gRPC client and server configuration.
//
// MaxConnectionAgeSettings controls the server-side maximum connection age
// via gRPC keepalive server parameters.
//
// Example usage:
//
//	settings := configgrpc.NewDefaultMaxConnectionAgeSettings()
//	settings.MaxAge = 30 * time.Minute
//	settings.MaxAgeGrace = 5 * time.Second
//	if err := settings.Validate(); err != nil {
//		// handle error
//	}
//	opt := settings.ToServerOption()
//	if opt != nil {
//		server := grpc.NewServer(opt)
//	}
package configgrpc
