// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

//go:build integration

package pgadapter_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/directory-api/domain/directory/storage"
)

func TestGetOrganizationInwayManagementAPIProxyAddress(t *testing.T) {
	t.Parallel()

	type args struct {
		organizationSerialNumber string
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		want         string
		wantError    error
	}{
		"no_address_registered": {
			loadFixtures: false,
			args: args{
				organizationSerialNumber: testOrganizationSerialNumber,
			},
			want:      "",
			wantError: storage.ErrNotFound,
		},
		"happy_flow": {
			loadFixtures: true,
			args: args{
				organizationSerialNumber: testOrganizationSerialNumber,
			},
			want:      "fixture-inway-proxy-address-one.com:8443",
			wantError: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			storage, close := new(t, tt.loadFixtures)
			defer close()

			got, err := storage.GetOrganizationInwayManagementAPIProxyAddress(context.Background(), tt.args.organizationSerialNumber)
			assert.Equal(t, tt.wantError, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
