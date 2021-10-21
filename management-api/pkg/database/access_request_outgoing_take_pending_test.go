// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration

package database_test

import (
	"context"
	"testing"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/management-api/pkg/database"
)

func TestTakePendingOutgoingAccessRequest(t *testing.T) {
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
		want         *database.OutgoingAccessRequest
		wantErr      error
	}{
		"happy_flow_no_pending_requests": {
			loadFixtures: false,
			want:         nil,
			wantErr:      nil,
		},
		"happy_flow": {
			loadFixtures: true,
			want: &database.OutgoingAccessRequest{
				ID: 1,
				Organization: database.Organization{
					SerialNumber: "00000000000000000001",
					Name:         "fixture-organization-name",
				},
				ServiceName:          "fixture-service-name",
				ReferenceID:          1,
				State:                database.OutgoingAccessRequestCreated,
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

			got, err := configDb.TakePendingOutgoingAccessRequest(context.Background())
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil && tt.want != nil {
				assert.Equal(t, tt.want.ID, got.ID)
				assert.Equal(t, tt.want.Organization, got.Organization)
				assert.Equal(t, tt.want.ServiceName, got.ServiceName)
				assert.Equal(t, tt.want.ReferenceID, got.ReferenceID)
				assert.Equal(t, tt.want.State, got.State)
				assert.Equal(t, tt.want.CreatedAt, got.CreatedAt)
				assert.Equal(t, tt.want.UpdatedAt, got.UpdatedAt)
				assert.Equal(t, tt.want.PublicKeyPEM, got.PublicKeyPEM)
				assert.Equal(t, tt.want.PublicKeyFingerprint, got.PublicKeyFingerprint)

				assert.NotNil(t, got.LockID)
				assert.True(t, got.LockExpiresAt.Time.After(time.Now()))
			}
		})
	}
}
