// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration

package storage_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClearIfSetAsOrganizationInway(t *testing.T) {
	t.Parallel()

	var (
		testOrgSerialNumber = "01234567890123456789"
		testInwayAddrNotOrg = "fixture-inway-address-two.com:443"
		testInwayAddrOrg    = "fixture-inway-address-one.com:443"
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

			storage, close := new(t, true)
			defer close()

			err := storage.ClearIfSetAsOrganizationInway(context.Background(), testOrgSerialNumber, tt.inwayAddress)
			require.NoError(t, err)

			orgInwayAddress, err := storage.GetOrganizationInwayAddress(context.Background(), testOrgSerialNumber)
			require.NoError(t, err)
			require.Equal(t, tt.expectedOrgInway, orgInwayAddress)
		})
	}
}
