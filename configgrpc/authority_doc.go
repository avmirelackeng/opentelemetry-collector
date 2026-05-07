// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package configgrpc provides configuration types for gRPC connections used
// within the OpenTelemetry Collector.
//
// Authority
//
// The Authority type represents the value of the HTTP/2 :authority pseudo-header
// (equivalent to the Host header in HTTP/1.1) that is sent with every gRPC
// request.  By default gRPC derives the authority from the target address;
// setting an explicit Authority overrides that behaviour.
//
// Example YAML configuration:
//
//	# Use a custom authority header
//	authority: "my-service.internal"
//
// An empty string (the default) means no override is applied and gRPC will
// use its built-in logic to determine the authority.
package configgrpc
