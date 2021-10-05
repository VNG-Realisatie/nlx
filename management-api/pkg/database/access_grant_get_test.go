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

func TestGetAccessGrant(t *testing.T) {
	t.Parallel()

	setup(t)

	fixtureTime := getFixtureTime(t)

	fixtureCertBundle, err := newFixtureCertificateBundle()
	require.NoError(t, err)

	fixturePublicKeyPEM, err := fixtureCertBundle.PublicKeyPEM()
	require.NoError(t, err)

	type args struct {
		id uint
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		want         *database.AccessGrant
		wantErr      error
	}{
		"when_the_access_grant_does_not_exist": {
			loadFixtures: false,
			args: args{
				id: 1,
			},
			want:    nil,
			wantErr: database.ErrNotFound,
		},
		"happy_flow": {
			loadFixtures: true,
			args: args{
				id: 1,
			},
			want: &database.AccessGrant{
				ID:                      1,
				CreatedAt:               fixtureTime,
				IncomingAccessRequestID: 1,
				IncomingAccessRequest: &database.IncomingAccessRequest{
					ID:        1,
					ServiceID: 1,
					Service: &database.Service{
						ID:                   1,
						Name:                 "fixture-service-name",
						EndpointURL:          "http://fixture-api:8000",
						DocumentationURL:     "https://fixture-documentation-url.com",
						APISpecificationURL:  "http://fixture-api:8000/schema?format=openapi-json",
						Internal:             false,
						TechSupportContact:   "fixture@tech-support-contact.com",
						PublicSupportContact: "fixture@public-support-contact.com",
						OneTimeCosts:         1,
						MonthlyCosts:         2,
						RequestCosts:         3,
						CreatedAt:            fixtureTime,
						UpdatedAt:            fixtureTime,
					},
					Organization: database.IncomingAccessRequestOrganization{
						Name:         "fixture-organization-name",
						SerialNumber: "00000000000000000001",
					},
					State:                database.IncomingAccessRequestReceived,
					CreatedAt:            fixtureTime,
					UpdatedAt:            fixtureTime,
					PublicKeyPEM:         fixturePublicKeyPEM,
					PublicKeyFingerprint: fixtureCertBundle.PublicKeyFingerprint()},
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

			got, err := configDb.GetAccessGrant(context.Background(), tt.args.id)
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				require.Equal(t, tt.want, got)
			}
		})
	}
}
