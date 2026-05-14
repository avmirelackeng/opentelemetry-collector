// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

// ServerConfig defines the configuration for a gRPC server.
type ServerConfig struct {
	// Endpoint is the address to listen on.
	Endpoint Endpoint `mapstructure:"endpoint"`

	// TLSSetting defines TLS settings for the server.
	TLSSetting *TLSVersion `mapstructure:"tls,omitempty"`

	// MaxRecvMsgSizeMiB is the max message size in MiB the server can receive.
	MaxRecvMsgSizeMiB MaxMessageSizeSettings `mapstructure:"max_recv_msg_size_mib"`

	// Keepalive defines the keepalive parameters for the server.
	Keepalive *KeepaliveServerParameters `mapstructure:"keepalive,omitempty"`

	// Interceptors defines interceptor settings.
	Interceptors InterceptorSettings `mapstructure:"interceptors"`
}

// NewDefaultServerConfig returns a ServerConfig with default values.
func NewDefaultServerConfig() ServerConfig {
	return ServerConfig{
		MaxRecvMsgSizeMiB: NewDefaultMaxMessageSizeSettings(),
		Interceptors:      NewDefaultInterceptorSettings(),
	}
}

// Validate checks that the ServerConfig is valid.
func (s *ServerConfig) Validate() error {
	if err := s.Endpoint.Validate(); err != nil {
		return errors.Join(errors.New("invalid endpoint"), err)
	}
	if err := s.MaxRecvMsgSizeMiB.Validate(); err != nil {
		return errors.Join(errors.New("invalid max_recv_msg_size_mib"), err)
	}
	if err := s.Interceptors.Validate(); err != nil {
		return errors.Join(errors.New("invalid interceptors"), err)
	}
	return nil
}

// ToServerOptions converts the ServerConfig to a slice of grpc.ServerOption.
func (s *ServerConfig) ToServerOptions() ([]grpc.ServerOption, error) {
	var opts []grpc.ServerOption

	recvOpts := s.MaxRecvMsgSizeMiB.ToServerOptions()
	opts = append(opts, recvOpts...)

	var creds credentials.TransportCredentials
	if s.TLSSetting != nil {
		creds = insecure.NewCredentials()
	} else {
		creds = insecure.NewCredentials()
	}
	opts = append(opts, grpc.Creds(creds))

	if s.Keepalive != nil {
		kOpts := s.Keepalive.ToServerOptions()
		opts = append(opts, kOpts...)
	}

	return opts, nil
}
