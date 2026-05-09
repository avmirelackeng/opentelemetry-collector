// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserAgentValidate(t *testing.T) {
	tests := []struct {
		name    string
		ua      UserAgent
		wantErr bool
	}{
		{name: "valid simple", ua: UserAgent("otelcol/0.90.0"), wantErr: false},
		{name: "valid composite", ua: UserAgent("myapp/1.0 otelcol/0.90.0"), wantErr: false},
		{name: "empty", ua: UserAgent(""), wantErr: true},
		{name: "control char tab", ua: UserAgent("otelcol\t0.90"), wantErr: true},
		{name: "control char newline", ua: UserAgent("otelcol\n0.90"), wantErr: true},
		{name: "delete char", ua: UserAgent("otelcol\x7f0.90"), wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.ua.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestUserAgentIsSet(t *testing.T) {
	assert.False(t, UserAgent("").IsSet())
	assert.True(t, UserAgent("otelcol/0.90.0").IsSet())
}

func TestUserAgentString(t *testing.T) {
	ua := UserAgent("otelcol/0.90.0")
	assert.Equal(t, "otelcol/0.90.0", ua.String())
}

func TestUserAgentWithProduct(t *testing.T) {
	tests := []struct {
		name     string
		base     UserAgent
		product  string
		expected UserAgent
	}{
		{name: "append to existing", base: UserAgent("myapp/1.0"), product: "otelcol/0.90.0", expected: UserAgent("myapp/1.0 otelcol/0.90.0")},
		{name: "empty base", base: UserAgent(""), product: "otelcol/0.90.0", expected: UserAgent("otelcol/0.90.0")},
		{name: "empty product", base: UserAgent("myapp/1.0"), product: "", expected: UserAgent("myapp/1.0")},
		{name: "whitespace product", base: UserAgent("myapp/1.0"), product: "   ", expected: UserAgent("myapp/1.0")},
		{name: "both empty", base: UserAgent(""), product: "", expected: UserAgent("")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.base.WithProduct(tt.product)
			assert.Equal(t, tt.expected, result)
		})
	}
}
