// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration

package database_test

import (
	"context"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/management-api/pkg/database"
)

func TestDeleteService(t *testing.T) {
	t.Parallel()

	setup(t)

	type args struct {
		serviceName              string
		organizationSerialNumber string
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		wantErr      error
	}{
		"happy_flow": {
			loadFixtures: true,
			args: args{
				serviceName:              "fixture-service-name",
				organizationSerialNumber: "00000000000000000001",
			},
			wantErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			configDb, close := newConfigDatabase(t, t.Name(), tt.loadFixtures)
			defer close()

			err := configDb.DeleteService(context.Background(), tt.args.serviceName, tt.args.organizationSerialNumber)
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				assertServiceDeleted(t, configDb, tt.args.serviceName, tt.args.organizationSerialNumber)
			}
		})
	}
}

func assertServiceDeleted(t *testing.T, repo database.ConfigDatabase, serviceName, organizationSerialNumber string) {
	_, err := repo.GetService(context.Background(), serviceName)
	require.Equal(t, database.ErrNotFound, err)

	orders, err := repo.ListOutgoingOrders(context.Background())
	require.NoError(t, err)

	for _, order := range orders {
		for _, outgoingOrderAccessProof := range order.OutgoingOrderAccessProofs {
			if outgoingOrderAccessProof.AccessProof.OutgoingAccessRequest.ServiceName == serviceName && outgoingOrderAccessProof.AccessProof.OutgoingAccessRequest.Organization.SerialNumber == organizationSerialNumber {
				t.Error("outgoing order for service is not deleted")
			}
		}
	}

}
