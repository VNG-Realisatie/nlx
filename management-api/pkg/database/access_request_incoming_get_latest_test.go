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

func TestGetLatestIncomingAccessRequest(t *testing.T) {
	t.Parallel()

	setup(t)

	fixtureTime := getFixtureTime(t)

	fixtureCertBundle, err := newFixtureCertificateBundle()
	require.NoError(t, err)

	fixturePublicKeyPEM, err := fixtureCertBundle.PublicKeyPEM()
	require.NoError(t, err)

	type args struct {
		organizationName string
		serviceName      string
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		want         *database.IncomingAccessRequest
		wantErr      error
	}{
		"when_service_is_not_present": {
			loadFixtures: true,
			args: args{
				organizationName: "fixture-organization-name",
				serviceName:      "non-existing-service-name",
			},
			want:    nil,
			wantErr: database.ErrNotFound,
		},
		"when_organization_is_not_present": {
			loadFixtures: true,
			args: args{
				organizationName: "non-existing-organization-name",
				serviceName:      "fixture-service-name",
			},
			want:    nil,
			wantErr: database.ErrNotFound,
		},
		"happy_flow": {
			loadFixtures: true,
			args: args{
				organizationName: "fixture-organization-name",
				serviceName:      "fixture-service-name",
			},
			want: &database.IncomingAccessRequest{
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
				OrganizationName:     "fixture-organization-name",
				State:                database.IncomingAccessRequestReceived,
				CreatedAt:            fixtureTime,
				UpdatedAt:            fixtureTime,
				PublicKeyPEM:         fixturePublicKeyPEM,
				PublicKeyFingerprint: fixtureCertBundle.PublicKeyFingerprint(),
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

			got, err := configDb.GetLatestIncomingAccessRequest(context.Background(), tt.args.organizationName, tt.args.serviceName)
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				require.Equal(t, tt.want, got)
			}
		})
	}
}