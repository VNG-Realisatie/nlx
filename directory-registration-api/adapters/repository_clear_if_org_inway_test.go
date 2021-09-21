// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration
// +build integration

package adapters_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClearIfSetAsOrganizationInway(t *testing.T) {
	t.Parallel()

	setup(t)

	var (
		testOrgSerialNumber = "01234567890123456789"
		testInwayAddrNotOrg = "https://fixture-inway-address-two.com"
		testInwayAddrOrg    = "https://fixture-inway-address-one.com"
	)

	tests := map[string]struct {
		inwayAddress     string
		expectedOrgInway string
	}{
		"is_not_org_inway": {
			inwayAddress:     testInwayAddrNotOrg,
			expectedOrgInway: testInwayAddrOrg,
		},
		"is_org_inway": {
			inwayAddress:     testInwayAddrOrg,
			expectedOrgInway: "",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			repo, close := newRepo(t, t.Name(), true)
			defer close()

			err := repo.ClearIfSetAsOrganizationInway(context.Background(), testOrgSerialNumber, tt.inwayAddress)
			require.NoError(t, err)

			orgInwayAddress, err := repo.GetOrganizationInwayAddress(context.Background(), testOrgSerialNumber)
			require.NoError(t, err)
			require.Equal(t, tt.expectedOrgInway, orgInwayAddress)
		})
	}
}
