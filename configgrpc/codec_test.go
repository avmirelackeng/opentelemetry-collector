// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCodecTypeConstants(t *testing.T) {
	assert.Equal(t, CodecType("proto"), CodecTypeProto)
	assert.Equal(t, CodecType("json"), CodecTypeJSON)
	assert.Equal(t, CodecType("raw"), CodecTypeRaw)
}

func TestCodecTypeValidate(t *testing.T) {
	tests := []struct {
		name    string
		codec   CodecType
		wantErr bool
	}{
		{name: "proto", codec: CodecTypeProto, wantErr: false},
		{name: "json", codec: CodecTypeJSON, wantErr: false},
		{name: "raw", codec: CodecTypeRaw, wantErr: false},
		{name: "empty", codec: "", wantErr: true},
		{name: "unknown", codec: "msgpack", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.codec.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestCodecTypeString(t *testing.T) {
	assert.Equal(t, "proto", CodecTypeProto.String())
	assert.Equal(t, "json", CodecTypeJSON.String())
	assert.Equal(t, "raw", CodecTypeRaw.String())
}

func TestCodecTypeIsDefault(t *testing.T) {
	assert.True(t, CodecTypeProto.IsDefault())
	assert.True(t, CodecType("").IsDefault())
	assert.False(t, CodecTypeJSON.IsDefault())
	assert.False(t, CodecTypeRaw.IsDefault())
}

func TestNewDefaultCodecSettings(t *testing.T) {
	cs := NewDefaultCodecSettings()
	assert.Equal(t, CodecTypeProto, cs.Type)
	require.NoError(t, cs.Validate())
}

func TestCodecSettingsValidate(t *testing.T) {
	valid := &CodecSettings{Type: CodecTypeJSON}
	require.NoError(t, valid.Validate())

	invalid := &CodecSettings{Type: "invalid"}
	require.Error(t, invalid.Validate())

	var nilSettings *CodecSettings
	require.NoError(t, nilSettings.Validate())
}
