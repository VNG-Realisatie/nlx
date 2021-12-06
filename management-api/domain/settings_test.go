// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package domain_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/management-api/domain"
)

func Test_NewSettings(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		organizationInwayName    string
		organizationEmailAddress string
		expectedErr              error
	}{
		"invalid_inway_name": {
			organizationInwayName:    "430941//,.§",
			organizationEmailAddress: "mock@email.com",
			expectedErr:              errors.New("inway name: must be in a valid format"),
		},
		"invalid_email_address": {
			organizationInwayName:    "mock-inway-name",
			organizationEmailAddress: "@invalidemail.com",
			expectedErr:              errors.New("organization email address: must be a valid email address"),
		},
		"happy_flow": {
			organizationInwayName:    "mock-inway-name",
			organizationEmailAddress: "mock@email.com",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			result, err := domain.NewSettings(tt.organizationInwayName, tt.organizationEmailAddress)

			require.Equal(t, tt.expectedErr, err)
			if tt.expectedErr == nil {
				require.Equal(t, tt.organizationInwayName, result.OrganizationInwayName())
				require.Equal(t, tt.organizationEmailAddress, result.OrganizationEmailAddress())
			}
		})
	}
}
