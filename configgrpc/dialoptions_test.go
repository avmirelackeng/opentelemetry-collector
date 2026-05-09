// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDialOptionConstants(t *testing.T) {
	assert.Equal(t, DialOption("block"), DialOptionBlock)
	assert.Equal(t, DialOption("insecure"), DialOptionInsecure)
	assert.Equal(t, DialOption("no_proxy"), DialOptionNoProxy)
}

func TestNewDefaultDialOptionsSettings(t *testing.T) {
	d := NewDefaultDialOptionsSettings()
	assert.Empty(t, d.Options)
	require.NoError(t, d.Validate())
}

func TestDialOptionsSettingsValidate(t *testing.T) {
	tests := []struct {
		name    string
		opts    []DialOption
		wantErr bool
	}{
		{name: "empty", opts: []DialOption{}, wantErr: false},
		{name: "block", opts: []DialOption{DialOptionBlock}, wantErr: false},
		{name: "insecure", opts: []DialOption{DialOptionInsecure}, wantErr: false},
		{name: "no_proxy", opts: []DialOption{DialOptionNoProxy}, wantErr: false},
		{name: "multiple valid", opts: []DialOption{DialOptionBlock, DialOptionNoProxy}, wantErr: false},
		{name: "unknown", opts: []DialOption{"unknown_option"}, wantErr: true},
		{name: "mixed valid and unknown", opts: []DialOption{DialOptionBlock, "bad"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := DialOptionsSettings{Options: tt.opts}
			err := d.Validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestDialOptionsSettingsToGRPCDialOptions(t *testing.T) {
	t.Run("empty options", func(t *testing.T) {
		d := NewDefaultDialOptionsSettings()
		opts, err := d.ToGRPCDialOptions()
		require.NoError(t, err)
		assert.Empty(t, opts)
	})

	t.Run("valid options produce dial options", func(t *testing.T) {
		d := DialOptionsSettings{Options: []DialOption{DialOptionBlock, DialOptionNoProxy}}
		opts, err := d.ToGRPCDialOptions()
		require.NoError(t, err)
		assert.Len(t, opts, 2)
	})

	t.Run("invalid option returns error", func(t *testing.T) {
		d := DialOptionsSettings{Options: []DialOption{"not_valid"}}
		opts, err := d.ToGRPCDialOptions()
		require.Error(t, err)
		assert.Nil(t, opts)
	})
}
