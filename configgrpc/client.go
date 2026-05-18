// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"errors"

	"google.golang.org/grpc"
)

// ClientConfig holds all configuration for a gRPC client connection.
type ClientConfig struct {
	Endpoint        Endpoint               `mapstructure:"endpoint"`
	TLS             TLSSettings            `mapstructure:"tls"`
	Keepalive       KeepaliveClientParams  `mapstructure:"keepalive"`
	Backoff         BackoffConfig          `mapstructure:"backoff"`
	ConnectParams   ConnectParams          `mapstructure:"connect_params"`
	Headers         Headers                `mapstructure:"headers"`
	Metadata        Metadata               `mapstructure:"metadata"`
	Compression     CompressionType        `mapstructure:"compression"`
	Authority       Authority              `mapstructure:"authority"`
	UserAgent       UserAgent              `mapstructure:"user_agent"`
	MaxMessageSize  MaxMessageSizeSettings `mapstructure:"max_message_size"`
	WindowSize      WindowSizeSettings     `mapstructure:"window_size"`
	SendRecvBufSize SendRecvBufferSizeSettings `mapstructure:"send_recv_buf_size"`
	DialOptions     DialOptionsSettings    `mapstructure:"dial_options"`
	ServiceConfig   ServiceConfig          `mapstructure:"service_config"`
	Codec           CodecSettings          `mapstructure:"codec"`
	Interceptor     InterceptorSettings    `mapstructure:"interceptor"`
	Credentials     CredentialsSettings    `mapstructure:"credentials"`
	Channelz        ChannelzSettings       `mapstructure:"channelz"`
	Tracing         TracingSettings        `mapstructure:"tracing"`
	WaitForReady    WaitForReady           `mapstructure:"wait_for_ready"`
	Timeout         TimeoutSettings        `mapstructure:"timeout"`
	Retry           *RetryPolicy           `mapstructure:"retry"`
	Resolver        ResolverScheme         `mapstructure:"resolver"`
	Balancer        BalancerName           `mapstructure:"balancer"`
}

// NewDefaultClientConfig returns a ClientConfig with sensible defaults.
func NewDefaultClientConfig() ClientConfig {
	return ClientConfig{
		Keepalive:       NewDefaultKeepaliveClientParams(),
		Backoff:         NewDefaultBackoffConfig(),
		ConnectParams:   NewDefaultConnectParams(),
		MaxMessageSize:  NewDefaultMaxMessageSizeSettings(),
		WindowSize:      NewDefaultWindowSizeSettings(),
		SendRecvBufSize: NewDefaultSendRecvBufferSizeSettings(),
		DialOptions:     NewDefaultDialOptionsSettings(),
		ServiceConfig:   NewDefaultServiceConfig(),
		Codec:           NewDefaultCodecSettings(),
		Interceptor:     NewDefaultInterceptorSettings(),
		Credentials:     NewDefaultCredentialsSettings(),
		Channelz:        NewDefaultChannelzSettings(),
		Tracing:         NewDefaultTracingSettings(),
		Timeout:         NewDefaultTimeoutSettings(),
	}
}

// Validate checks that the ClientConfig is valid.
func (c ClientConfig) Validate() error {
	return errors.Join(
		c.Endpoint.Validate(),
		c.Compression.Validate(),
		c.Authority.Validate(),
		c.MaxMessageSize.Validate(),
		c.WindowSize.Validate(),
		c.SendRecvBufSize.Validate(),
		c.DialOptions.Validate(),
		c.ServiceConfig.Validate(),
		c.Codec.Validate(),
		c.InterceptorSettings.Validate(),
		c.Credentials.Validate(),
		c.Channelz.Validate(),
		c.Tracing.Validate(),
		c.Timeout.Validate(),
		c.Resolver.Validate(),
		c.Balancer.Validate(),
	)
}

// ToDialOptions converts the ClientConfig into a slice of grpc.DialOption.
func (c ClientConfig) ToDialOptions() ([]grpc.DialOption, error) {
	var opts []grpc.DialOption

	if dopts := c.MaxMessageSize.ToDialOptions(); len(dopts) > 0 {
		opts = append(opts, dopts...)
	}
	if dopts := c.WindowSize.ToDialOptions(); len(dopts) > 0 {
		opts = append(opts, dopts...)
	}
	if dopts := c.SendRecvBufSize.ToDialOptions(); len(dopts) > 0 {
		opts = append(opts, dopts...)
	}
	if opt := c.Channelz.ToDialOption(); opt != nil {
		opts = append(opts, opt)
	}
	if opt := c.ConnectParams.ToDialOption(); opt != nil {
		opts = append(opts, opt)
	}
	return opts, nil
}
