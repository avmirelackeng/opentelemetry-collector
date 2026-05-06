// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCompressionTypeConstants(t *testing.T) {
	assert.Equal(t, CompressionType("none"), CompressionNone)
	assert.Equal(t, CompressionType("gzip"), CompressionGzip)
	assert.Equal(t, CompressionType("snappy"), CompressionSnappy)
	assert.Equal(t, CompressionType("zstd"), CompressionZstd)
}

func TestCompressionTypeValidate(t *testing.T) {
	validTypes := []CompressionType{
		CompressionNone,
		CompressionGzip,
		CompressionSnappy,
		CompressionZstd,
	}
	for _, ct := range validTypes {
		t.Run(string(ct), func(t *testing.T) {
			require.NoError(t, ct.Validate())
		})
	}

	invalidTypes := []CompressionType{
		"lz4",
		"brotli",
		"",
		"GZIP",
	}
	for _, ct := range invalidTypes {
		t.Run(string(ct), func(t *testing.T) {
			require.Error(t, ct.Validate())
		})
	}
}

func TestCompressionTypeString(t *testing.T) {
	assert.Equal(t, "gzip", CompressionGzip.String())
	assert.Equal(t, "none", CompressionNone.String())
	assert.Equal(t, "snappy", CompressionSnappy.String())
	assert.Equal(t, "zstd", CompressionZstd.String())
}

func TestCompressionTypeIsEnabled(t *testing.T) {
	assert.False(t, CompressionNone.IsEnabled())
	assert.False(t, CompressionType("").IsEnabled())
	assert.True(t, CompressionGzip.IsEnabled())
	assert.True(t, CompressionSnappy.IsEnabled())
	assert.True(t, CompressionZstd.IsEnabled())
}
