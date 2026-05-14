// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ClientConfig holds the configuration for a gRPC client connection.
// It aggregates all the individual settings needed to establish and
// configure a gRPC client.
type ClientConfig struct {
	// Endpoint is the target endpoint for the gRPC connection.
	Endpoint Endpoint `mapstructure:"endpoint"`

	// TLS contains TLS configuration for the connection.
	TLS TLSClientSettings `mapstructure:"tls"`

	// Keepalive contains keepalive settings for the connection.
	Keepalive *KeepaliveClientConfig `mapstructure:"keepalive"`

	// ReadBufferSize sets the size of the read buffer for the connection.
	ReadBufferSize int `mapstructure:"read_buffer_size"`

	// WriteBufferSize sets the size of the write buffer for the connection.
	WriteBufferSize int `mapstructure:"write_buffer_size"`

	// WaitForReady controls whether the client waits for the connection
	// to be ready before sending RPCs.
	WaitForReady WaitForReady `mapstructure:"wait_for_ready"`

	// Headers are additional metadata to send with each RPC.
	Headers Headers `mapstructure:"headers"`

	// BalancerName specifies the balancer to use for the connection.
	BalancerName BalancerName `mapstructure:"balancer_name"`

	// Authority overrides the authority used for TLS handshake.
	Authority Authority `mapstructure:"authority"`

	// Auth configures client-side authentication.
	Auth *AuthConfig `mapstructure:"auth"`
}

// AuthConfig holds authentication configuration for a gRPC client.
type AuthConfig struct {
	// AuthenticatorID is the identifier of the authenticator to use.
	AuthenticatorID string `mapstructure:"authenticator"`
}

// NewDefaultClientConfig returns a ClientConfig with sensible defaults.
func NewDefaultClientConfig() ClientConfig {
	return ClientConfig{
		WaitForReady: WaitForReadyDisabled,
	}
}

// Validate checks that the ClientConfig fields are valid.
func (c *ClientConfig) Validate() error {
	if err := c.Endpoint.Validate(); err != nil {
		return fmt.Errorf("invalid endpoint: %w", err)
	}
	if err := c.WaitForReady.Validate(); err != nil {
		return fmt.Errorf("invalid wait_for_ready: %w", err)
	}
	if err := c.Headers.Validate(); err != nil {
		return fmt.Errorf("invalid headers: %w", err)
	}
	if err := c.BalancerName.Validate(); err != nil {
		return fmt.Errorf("invalid balancer_name: %w", err)
	}
	if err := c.Authority.Validate(); err != nil {
		return fmt.Errorf("invalid authority: %w", err)
	}
	if c.Keepalive != nil {
		if err := c.Keepalive.Validate(); err != nil {
			return fmt.Errorf("invalid keepalive: %w", err)
		}
	}
	return nil
}

// ToDialOptions converts the ClientConfig into a slice of grpc.DialOption
// that can be used to establish a gRPC connection.
func (c *ClientConfig) ToDialOptions(ctx context.Context) ([]grpc.DialOption, error) {
	var opts []grpc.DialOption

	// Use insecure credentials by default; TLS config can override.
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if c.ReadBufferSize > 0 {
		opts = append(opts, grpc.WithReadBufferSize(c.ReadBufferSize))
	}
	if c.WriteBufferSize > 0 {
		opts = append(opts, grpc.WithWriteBufferSize(c.WriteBufferSize))
	}

	if c.Keepalive != nil {
		kpOpt, err := c.Keepalive.ToDialOption()
		if err != nil {
			return nil, fmt.Errorf("failed to build keepalive dial option: %w", err)
		}
		opts = append(opts, kpOpt)
	}

	if c.BalancerName.IsSet() {
		opts = append(opts, grpc.WithDefaultServiceConfig(
			fmt.Sprintf(`{"loadBalancingPolicy":%q}`, c.BalancerName.String()),
		))
	}

	if c.Authority.IsSet() {
		opts = append(opts, grpc.WithAuthority(c.Authority.String()))
	}

	return opts, nil
}
