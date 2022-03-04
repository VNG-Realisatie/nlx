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

func TestListIncomingAccessRequests(t *testing.T) {
	t.Parallel()

	setup(t)

	fixtureTime := getFixtureTime(t)

	fixtureCertBundle, err := newFixtureCertificateBundle()
	require.NoError(t, err)

	fixturePEM, err := fixtureCertBundle.PublicKeyPEM()
	require.NoError(t, err)

	tests := map[string]struct {
		loadFixtures   bool
		argServiceName string
		want           []*database.IncomingAccessRequest
		wantErr        error
	}{
		"happy_flow": {
			loadFixtures:   true,
			argServiceName: "fixture-service-name",
			want: []*database.IncomingAccessRequest{
				{
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
					PublicKeyPEM:         fixturePEM,
				},
				{
					ID:        2,
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
						Name:         "fixture-organization-name-two",
						SerialNumber: "00000000000000000002",
					},
					State:                database.IncomingAccessRequestReceived,
					CreatedAt:            fixtureTime,
					UpdatedAt:            fixtureTime,
					PublicKeyFingerprint: fixtureCertBundle.PublicKeyFingerprint(),
					PublicKeyPEM:         fixturePEM,
				},
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

			got, err := configDb.ListIncomingAccessRequests(context.Background(), tt.argServiceName)
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				require.Equal(t, tt.want, got)
			}
		})
	}
}
