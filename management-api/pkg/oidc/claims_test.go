// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package oidc

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClaimVerify(t *testing.T) {
	tests := map[string]struct {
		claims  *Claims
		wantErr error
	}{
		"expired": {
			claims: &Claims{
				Subject:   "subject",
				Name:      "name",
				Email:     "test@nlx.io",
				Picture:   "mock-picture",
				Issuer:    "nlx",
				Audience:  audience{"audience"},
				ExpiresAt: time.Now().Unix(),
			},
			wantErr: fmt.Errorf("JWT is expired"),
		},
		"used_before_issued": {
			claims: &Claims{
				Subject:   "subject",
				Name:      "name",
				Email:     "test@nlx.io",
				Picture:   "mock-picture",
				Issuer:    "nlx",
				Audience:  audience{"audience"},
				ExpiresAt: time.Now().Add(10 * time.Second).Unix(),
				IssuedAt:  time.Now().Add(10 * time.Second).Unix(),
			},
			wantErr: fmt.Errorf("cannot use JWT before it has been issued"),
		},
		"used_before_valid": {
			claims: &Claims{
				Subject:   "subject",
				Name:      "name",
				Email:     "test@nlx.io",
				Picture:   "mock-picture",
				Issuer:    "nlx",
				Audience:  audience{"audience"},
				ExpiresAt: time.Now().Add(10 * time.Second).Unix(),
				IssuedAt:  time.Now().Unix(),
				NotBefore: time.Now().Add(10 * time.Second).Unix(),
			},
			wantErr: fmt.Errorf("cannot use JWT before it's valid"),
		},
		"happy_flow": {
			claims: &Claims{
				Subject:   "subject",
				Name:      "name",
				Email:     "test@nlx.io",
				Picture:   "mock-picture",
				Issuer:    "nlx",
				Audience:  audience{"audience"},
				ExpiresAt: time.Now().Add(10 * time.Second).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
			wantErr: nil,
		},
	}

	for name, test := range tests {
		tc := test

		t.Run(name, func(t *testing.T) {
			err := tc.claims.Verify()
			assert.Equal(t, tc.wantErr, err)
		})
	}
}
