// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc // import "go.opentelemetry.io/collector/configgrpc"

import (
	"fmt"
	"strings"
)

// ResolverScheme represents the gRPC resolver scheme to use for name resolution.
type ResolverScheme string

const (
	// ResolverSchemeDNS uses DNS-based name resolution.
	ResolverSchemeDNS ResolverScheme = "dns"
	// ResolverSchemePassthrough bypasses name resolution.
	ResolverSchemePassthrough ResolverScheme = "passthrough"
	// ResolverSchemeIPv4 resolves using IPv4 addresses.
	ResolverSchemeIPv4 ResolverScheme = "ipv4"
	// ResolverSchemeIPv6 resolves using IPv6 addresses.
	ResolverSchemeIPv6 ResolverScheme = "ipv6"
)

var validResolverSchemes = []ResolverScheme{
	ResolverSchemeDNS,
	ResolverSchemePassthrough,
	ResolverSchemeIPv4,
	ResolverSchemeIPv6,
}

// Validate checks that the ResolverScheme is one of the supported values.
func (r ResolverScheme) Validate() error {
	if r == "" {
		return nil
	}
	for _, v := range validResolverSchemes {
		if r == v {
			return nil
		}
	}
	schemes := make([]string, len(validResolverSchemes))
	for i, s := range validResolverSchemes {
		schemes[i] = string(s)
	}
	return fmt.Errorf("unsupported resolver scheme %q, supported schemes: [%s]", r, strings.Join(schemes, ", "))
}

// ApplyToEndpoint prepends the resolver scheme to the endpoint if a scheme is set
// and the endpoint does not already contain a scheme prefix.
func (r ResolverScheme) ApplyToEndpoint(endpoint string) string {
	if r == "" {
		return endpoint
	}
	if strings.Contains(endpoint, "://") {
		return endpoint
	}
	return fmt.Sprintf("%s:///%s", r, endpoint)
}
