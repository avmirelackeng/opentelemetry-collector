// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"errors"
	"strings"

	"google.golang.org/grpc/metadata"
)

// MetadataKey represents a gRPC metadata key.
type MetadataKey string

// MetadataValue represents a gRPC metadata value.
type MetadataValue string

// Metadata is a map of gRPC metadata key-value pairs.
type Metadata map[MetadataKey][]MetadataValue

// Validate checks that all metadata keys and values are valid.
func (m Metadata) Validate() error {
	for k := range m {
		if strings.TrimSpace(string(k)) == "" {
			return errors.New("metadata key must not be empty")
		}
		if !isValidMetadataKey(string(k)) {
			return errors.New("metadata key \"" + string(k) + "\" contains invalid characters")
		}
	}
	return nil
}

// ToGRPCMetadata converts Metadata to grpc/metadata.MD.
func (m Metadata) ToGRPCMetadata() metadata.MD {
	md := metadata.MD{}
	for k, vals := range m {
		key := strings.ToLower(string(k))
		for _, v := range vals {
			md[key] = append(md[key], string(v))
		}
	}
	return md
}

// isValidMetadataKey returns true if the key is a valid gRPC metadata key.
// Keys must consist of printable ASCII characters except for colon, space, and comma.
func isValidMetadataKey(key string) bool {
	if len(key) == 0 {
		return false
	}
	for _, c := range key {
		if c < 0x20 || c > 0x7E {
			return false
		}
		if c == ':' || c == ',' {
			return false
		}
	}
	return true
}
