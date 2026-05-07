// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configgrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func validRetryPolicy() *RetryPolicy {
	return &RetryPolicy{
		MaxAttempts:          3,
		InitialBackoff:       0.5,
		MaxBackoff:           30.0,
		BackoffMultiplier:    2.0,
		RetryableStatusCodes: []string{"UNAVAILABLE"},
	}
}

func TestRetryPolicyValidate(t *testing.T) {
	require.NoError(t, validRetryPolicy().Validate())
}

func TestRetryPolicyNilValidate(t *testing.T) {
	var r *RetryPolicy
	require.NoError(t, r.Validate())
}

func TestRetryPolicyMaxAttempts(t *testing.T) {
	r := validRetryPolicy()
	r.MaxAttempts = 1
	assert.ErrorContains(t, r.Validate(), "max_attempts")

	r.MaxAttempts = 0
	assert.ErrorContains(t, r.Validate(), "max_attempts")
}

func TestRetryPolicyInitialBackoff(t *testing.T) {
	r := validRetryPolicy()
	r.InitialBackoff = 0
	assert.ErrorContains(t, r.Validate(), "initial_backoff")

	r.InitialBackoff = -1
	assert.ErrorContains(t, r.Validate(), "initial_backoff")
}

func TestRetryPolicyMaxBackoff(t *testing.T) {
	r := validRetryPolicy()
	r.MaxBackoff = 0
	assert.ErrorContains(t, r.Validate(), "max_backoff")

	r.MaxBackoff = -1
	assert.ErrorContains(t, r.Validate(), "max_backoff")
}

func TestRetryPolicyMaxBackoffLessThanInitial(t *testing.T) {
	r := validRetryPolicy()
	r.MaxBackoff = 0.1
	r.InitialBackoff = 1.0
	assert.ErrorContains(t, r.Validate(), "max_backoff")
}

func TestRetryPolicyBackoffMultiplier(t *testing.T) {
	r := validRetryPolicy()
	r.BackoffMultiplier = 0
	assert.ErrorContains(t, r.Validate(), "backoff_multiplier")

	r.BackoffMultiplier = -0.5
	assert.ErrorContains(t, r.Validate(), "backoff_multiplier")
}

func TestRetryPolicyRetryableStatusCodes(t *testing.T) {
	r := validRetryPolicy()
	r.RetryableStatusCodes = nil
	assert.ErrorContains(t, r.Validate(), "retryable_status_codes")

	r.RetryableStatusCodes = []string{}
	assert.ErrorContains(t, r.Validate(), "retryable_status_codes")
}
