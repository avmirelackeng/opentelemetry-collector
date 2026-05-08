// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMetadataValidate(t *testing.T) {
	tests := []struct {
		name    string
		md      Metadata
		wantErr bool
	}{
		{name: "valid single", md: Metadata{"authorization": {"Bearer token"}}, wantErr: false},
		{name: "valid multiple keys", md: Metadata{"x-custom": {"v1", "v2"}, "x-other": {"v3"}}, wantErr: false},
		{name: "empty metadata", md: Metadata{}, wantErr: false},
		{name: "nil metadata", md: nil, wantErr: false},
		{name: "empty key", md: Metadata{"": {"value"}}, wantErr: true},
		{name: "whitespace key", md: Metadata{"   ": {"value"}}, wantErr: true},
		{name: "key with colon", md: Metadata{"bad:key": {"value"}}, wantErr: true},
		{name: "key with comma", md: Metadata{"bad,key": {"value"}}, wantErr: true},
		{name: "key with non-ascii", md: Metadata{"bàd": {"value"}}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.md.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMetadataToGRPCMetadata(t *testing.T) {
	md := Metadata{
		"Authorization": {"Bearer token"},
		"X-Custom-Header": {"val1", "val2"},
	}
	grpcMD := md.ToGRPCMetadata()

	assert.Equal(t, []string{"Bearer token"}, grpcMD["authorization"])
	assert.Equal(t, []string{"val1", "val2"}, grpcMD["x-custom-header"])
}

func TestMetadataToGRPCMetadataEmpty(t *testing.T) {
	md := Metadata{}
	grpcMD := md.ToGRPCMetadata()
	assert.Empty(t, grpcMD)
}

func TestMetadataToGRPCMetadataKeyLowercased(t *testing.T) {
	md := Metadata{"X-UPPER": {"value"}}
	grpcMD := md.ToGRPCMetadata()
	_, hasUpper := grpcMD["X-UPPER"]
	_, hasLower := grpcMD["x-upper"]
	assert.False(t, hasUpper)
	assert.True(t, hasLower)
}

func TestIsValidMetadataKey(t *testing.T) {
	assert.True(t, isValidMetadataKey("valid-key"))
	assert.True(t, isValidMetadataKey("x-custom-bin"))
	assert.False(t, isValidMetadataKey(""))
	assert.False(t, isValidMetadataKey("bad:key"))
	assert.False(t, isValidMetadataKey("bad,key"))
}
