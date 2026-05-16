// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

// ChannelzSettings exposes gRPC channelz configuration.
//
// Channelz is a debugging tool built into gRPC that provides detailed
// information about the internal state of gRPC channels, subchannels,
// servers, and sockets. When enabled, a channelz service is registered
// and can be queried to inspect live connection states, call statistics,
// and channel events.
//
// Example configuration (YAML):
//
//	channelz:
//	  enabled: true
//	  address: ":9090"
