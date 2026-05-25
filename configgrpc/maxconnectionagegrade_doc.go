// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package configgrpc provides gRPC configuration utilities.
//
// MaxConnectionAgeGraceSettings controls the grace period after MaxConnectionAge
// before the server forcibly closes connections. This allows in-flight RPCs to
// complete before the connection is terminated.
//
// Example usage:
//
//	settings := configgrpc.NewDefaultMaxConnectionAgeGraceSettings()
//	settings.Enabled = true
//	settings.Grace = 5 * time.Second
//	if err := settings.Validate(); err != nil {
//	    // handle error
//	}
//	opt := settings.ToServerOption()
package configgrpc
