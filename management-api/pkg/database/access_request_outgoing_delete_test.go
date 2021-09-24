// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration
// +build integration

package database_test

import (
	"context"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"go.nlx.io/nlx/management-api/pkg/database"
)

func TestDeleteOutgoingAccessRequests(t *testing.T) {
	t.Parallel()

	setup(t)

	type args struct {
		organizationName string
		serviceName      string
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		wantErr      error
	}{
		"when_there_are_no_access_requests_present": {
			loadFixtures: false,
			args: args{
				organizationName: "arbitrary",
				serviceName:      "arbitrary",
			},
			wantErr: nil,
		},
		"happy_flow": {
			loadFixtures: true,
			args: args{
				organizationName: "fixture-organization-name",
				serviceName:      "fixture-service-name",
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

			err := configDb.DeleteOutgoingAccessRequests(context.Background(), tt.args.organizationName, tt.args.serviceName)
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				assertOutgoingAccessRequestDeleted(t, configDb, tt.args.organizationName, tt.args.serviceName)
			}
		})
	}
}

func assertOutgoingAccessRequestDeleted(t *testing.T, repo database.ConfigDatabase, organizationName, serviceName string) {
	_, err := repo.GetLatestOutgoingAccessRequest(context.Background(), organizationName, serviceName)
	require.Equal(t, err, database.ErrNotFound)
}
