// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc // import "go.opentelemetry.io/collector/configgrpc"

import (
	"fmt"

	"google.golang.org/grpc"
)

// WaitForReady is a typed boolean that controls gRPC's "wait for ready" semantics.
// When enabled, RPCs will block waiting for the connection to be in the READY
// state instead of failing immediately when the connection is in the
// TRANSIENT_FAILURE state.
//
// See https://github.com/grpc/grpc/blob/master/doc/wait-for-ready.md for details.
type WaitForReady bool

const (
	// WaitForReadyEnabled causes RPCs to wait until the connection is ready.
	WaitForReadyEnabled WaitForReady = true

	// WaitForReadyDisabled causes RPCs to fail fast when the connection is not ready.
	WaitForReadyDisabled WaitForReady = false
)

// Validate checks that the WaitForReady value is valid.
// Since WaitForReady is a boolean type, all values are valid.
func (w WaitForReady) Validate() error {
	return nil
}

// String returns the string representation of the WaitForReady value.
func (w WaitForReady) String() string {
	if w {
		return "true"
	}
	return "false"
}

// AsCallOption returns a grpc.CallOption that applies the WaitForReady setting
// to an individual RPC call.
func (w WaitForReady) AsCallOption() grpc.CallOption {
	return grpc.WaitForReady(bool(w))
}

// AsDialOption returns a grpc.DialOption that applies the WaitForReady setting
// as the default for all RPC calls on the connection.
func (w WaitForReady) AsDialOption() grpc.DialOption {
	if w {
		return grpc.WithDefaultCallOptions(grpc.WaitForReady(true))
	}
	return grpc.WithDefaultCallOptions(grpc.WaitForReady(false))
}

// ParseWaitForReady parses a string into a WaitForReady value.
// Accepts "true", "false", "1", "0", "yes", "no" (case-insensitive).
func ParseWaitForReady(s string) (WaitForReady, error) {
	switch s {
	case "true", "1", "yes", "True", "TRUE", "YES":
		return WaitForReadyEnabled, nil
	case "false", "0", "no", "False", "FALSE", "NO":
		return WaitForReadyDisabled, nil
	default:
		return WaitForReadyDisabled, fmt.Errorf("unknown wait-for-ready value %q: expected true or false", s)
	}
}
