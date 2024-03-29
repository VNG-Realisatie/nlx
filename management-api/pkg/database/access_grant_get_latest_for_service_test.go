// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration

package database_test

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/management-api/pkg/database"
)

func TestGetLatestAccessGrantForService(t *testing.T) {
	t.Parallel()

	setup(t)

	fixtureTime := getFixtureTime(t)

	fixtureCertBundle, err := newFixtureCertificateBundle()
	require.NoError(t, err)

	fixtureCertBundlePEM, err := fixtureCertBundle.PublicKeyPEM()
	require.NoError(t, err)

	type args struct {
		organizationSerialNumber string
		serviceName              string
		publicKeyFingerprint     string
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		want         *database.AccessGrant
		wantErr      error
	}{
		"when_organization_not_found": {
			loadFixtures: true,
			args: args{
				organizationSerialNumber: "00000000000000000000",
				serviceName:              "fixture-service-name",
			},
			want:    nil,
			wantErr: database.ErrNotFound,
		},
		"when_service_not_found": {
			loadFixtures: true,
			args: args{
				organizationSerialNumber: "00000000000000000001",
				serviceName:              "non-existing-service",
			},
			want:    nil,
			wantErr: database.ErrNotFound,
		},
		"happy_flow": {
			loadFixtures: true,
			args: args{
				organizationSerialNumber: "00000000000000000001",
				serviceName:              "fixture-service-name",
				publicKeyFingerprint:     "g+jpuLAMFzM09tOZpb0Ehslhje4S/IsIxSWsS4E16Yc=",
			},
			want: &database.AccessGrant{
				ID:                      1,
				IncomingAccessRequestID: 1,
				IncomingAccessRequest: &database.IncomingAccessRequest{
					ID:        1,
					ServiceID: 1,
					Service: &database.Service{
						ID:                     1,
						Name:                   "fixture-service-name",
						EndpointURL:            "http://fixture-api:8000",
						DocumentationURL:       "https://fixture-documentation-url.com",
						APISpecificationURL:    "http://fixture-api:8000/schema?format=openapi-json",
						Internal:               false,
						TechSupportContact:     "fixture@tech-support-contact.com",
						PublicSupportContact:   "fixture@public-support-contact.com",
						Inways:                 nil,
						IncomingAccessRequests: nil,
						OneTimeCosts:           1,
						MonthlyCosts:           2,
						RequestCosts:           3,
						CreatedAt:              fixtureTime,
						UpdatedAt:              fixtureTime,
					},
					Organization: database.IncomingAccessRequestOrganization{
						Name:         "fixture-organization-name",
						SerialNumber: "00000000000000000001",
					},
					State:                database.IncomingAccessRequestReceived,
					CreatedAt:            fixtureTime,
					UpdatedAt:            fixtureTime,
					PublicKeyFingerprint: fixtureCertBundle.PublicKeyFingerprint(),
					PublicKeyPEM:         fixtureCertBundlePEM,
				},
				CreatedAt: fixtureTime,
				RevokedAt: sql.NullTime{},
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

			got, err := configDb.GetLatestAccessGrantForService(context.Background(), tt.args.organizationSerialNumber, tt.args.serviceName, tt.args.publicKeyFingerprint)
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				require.EqualValues(t, tt.want, got)
			}
		})
	}
}
