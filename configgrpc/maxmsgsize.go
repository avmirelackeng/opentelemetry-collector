// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"fmt"

	"google.golang.org/grpc"
)

// MaxMessageSizeBytes represents a gRPC max message size in bytes.
// A value of 0 means the default gRPC max message size will be used.
type MaxMessageSizeBytes int

const (
	// DefaultMaxRecvMsgSizeMiB is the default max receive message size (4 MiB).
	DefaultMaxRecvMsgSizeMiB = 4
	// DefaultMaxSendMsgSizeMiB is the default max send message size (unlimited in gRPC, but we cap it).
	DefaultMaxSendMsgSizeMiB = 0
)

// MaxMessageSizeSettings holds the max send and receive message size settings.
type MaxMessageSizeSettings struct {
	// RecvMsgSizeMiB is the maximum message size in mebibytes the client/server can receive.
	// 0 means use the gRPC default.
	RecvMsgSizeMiB int `mapstructure:"recv_msg_size_mib"`
	// SendMsgSizeMiB is the maximum message size in mebibytes the client/server can send.
	// 0 means use the gRPC default.
	SendMsgSizeMiB int `mapstructure:"send_msg_size_mib"`
}

// NewDefaultMaxMessageSizeSettings returns a MaxMessageSizeSettings with default values.
func NewDefaultMaxMessageSizeSettings() MaxMessageSizeSettings {
	return MaxMessageSizeSettings{
		RecvMsgSizeMiB: DefaultMaxRecvMsgSizeMiB,
		SendMsgSizeMiB: DefaultMaxSendMsgSizeMiB,
	}
}

// Validate checks that the MaxMessageSizeSettings are valid.
func (m MaxMessageSizeSettings) Validate() error {
	if m.RecvMsgSizeMiB < 0 {
		return fmt.Errorf("recv_msg_size_mib must be non-negative, got %d", m.RecvMsgSizeMiB)
	}
	if m.SendMsgSizeMiB < 0 {
		return fmt.Errorf("send_msg_size_mib must be non-negative, got %d", m.SendMsgSizeMiB)
	}
	return nil
}

// ToDialOptions converts the MaxMessageSizeSettings to gRPC dial options.
func (m MaxMessageSizeSettings) ToDialOptions() []grpc.DialOption {
	var opts []grpc.DialOption
	if m.RecvMsgSizeMiB > 0 {
		opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(m.RecvMsgSizeMiB*1024*1024)))
	}
	if m.SendMsgSizeMiB > 0 {
		opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(m.SendMsgSizeMiB*1024*1024)))
	}
	return opts
}

// ToServerOptions converts the MaxMessageSizeSettings to gRPC server options.
func (m MaxMessageSizeSettings) ToServerOptions() []grpc.ServerOption {
	var opts []grpc.ServerOption
	if m.RecvMsgSizeMiB > 0 {
		opts = append(opts, grpc.MaxRecvMsgSize(m.RecvMsgSizeMiB*1024*1024))
	}
	if m.SendMsgSizeMiB > 0 {
		opts = append(opts, grpc.MaxSendMsgSize(m.SendMsgSizeMiB*1024*1024))
	}
	return opts
}
