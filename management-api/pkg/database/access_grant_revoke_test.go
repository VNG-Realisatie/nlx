// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration

package database_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/management-api/pkg/database"
)

func TestRevokeAccessGrant(t *testing.T) {
	t.Parallel()

	setup(t)

	now := time.Now()

	fixtureTimeNew := getCustomFixtureTime(t, "2021-01-04T01:02:03Z")

	fixtureTime := getFixtureTime(t)

	fixtureCertBundle, err := newFixtureCertificateBundle()
	require.NoError(t, err)

	fixturePEM, err := fixtureCertBundle.PublicKeyPEM()
	require.NoError(t, err)

	type args struct {
		accessGrantID uint
		revokedAt     time.Time
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		want         *database.AccessGrant
		wantErr      error
	}{
		"when_access_grant_does_not_exist": {
			loadFixtures: true,
			args: args{
				accessGrantID: 42,
				revokedAt:     now,
			},
			want:    nil,
			wantErr: database.ErrNotFound,
		},
		"when_access_grant_is_already_revoked": {
			loadFixtures: true,
			args: args{
				accessGrantID: 2,
				revokedAt:     now,
			},
			want:    nil,
			wantErr: database.ErrAccessGrantAlreadyRevoked,
		},
		"happy_flow": {
			loadFixtures: true,
			args: args{
				accessGrantID: 1,
				revokedAt:     fixtureTimeNew,
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
					State:                database.IncomingAccessRequestRevoked,
					CreatedAt:            fixtureTime,
					PublicKeyFingerprint: fixtureCertBundle.PublicKeyFingerprint(),
					PublicKeyPEM:         fixturePEM,
				},
				CreatedAt: fixtureTime,
				RevokedAt: sql.NullTime{Time: fixtureTimeNew, Valid: true},
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

			got, err := configDb.RevokeAccessGrant(context.Background(), tt.args.accessGrantID, tt.args.revokedAt)
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				// This check makes sure that updated at has been set.
				require.WithinDuration(t, now, got.IncomingAccessRequest.UpdatedAt, 1*time.Minute)
				// Set to empty time because we are unable to determine the updated at.
				got.IncomingAccessRequest.UpdatedAt = time.Time{}
				require.Equal(t, tt.want, got)
			}
		})
	}
}
