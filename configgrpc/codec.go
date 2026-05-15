// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"errors"
	"fmt"
)

// CodecType represents the codec to use for gRPC message encoding/decoding.
type CodecType string

const (
	// CodecTypeProto is the default protobuf codec.
	CodecTypeProto CodecType = "proto"
	// CodecTypeJSON is the JSON codec.
	CodecTypeJSON CodecType = "json"
	// CodecTypeRaw is the raw bytes codec (no encoding).
	CodecTypeRaw CodecType = "raw"
)

// Validate checks that the CodecType is valid.
func (c CodecType) Validate() error {
	switch c {
	case CodecTypeProto, CodecTypeJSON, CodecTypeRaw:
		return nil
	case "":
		return errors.New("codec type must not be empty")
	default:
		return fmt.Errorf("unknown codec type %q; valid values are: proto, json, raw", c)
	}
}

// String returns the string representation of the CodecType.
func (c CodecType) String() string {
	return string(c)
}

// IsDefault returns true if the codec type is the default (proto).
func (c CodecType) IsDefault() bool {
	return c == CodecTypeProto || c == ""
}

// CodecSettings holds configuration for gRPC codec selection.
type CodecSettings struct {
	// Type specifies the codec type to use.
	Type CodecType `mapstructure:"type"`
}

// NewDefaultCodecSettings returns CodecSettings with default values.
func NewDefaultCodecSettings() CodecSettings {
	return CodecSettings{
		Type: CodecTypeProto,
	}
}

// Validate checks that the CodecSettings are valid.
func (cs *CodecSettings) Validate() error {
	if cs == nil {
		return nil
	}
	return cs.Type.Validate()
}
