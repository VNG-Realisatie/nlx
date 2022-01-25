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

	"go.nlx.io/nlx/management-api/domain"
	"go.nlx.io/nlx/management-api/pkg/database"
)

func TestGetTermsOfServiceStatus(t *testing.T) {
	t.Parallel()

	setup(t)

	fixtureTime := getFixtureTime(t)

	type args struct {
		name string
	}

	tests := map[string]struct {
		loadFixtures bool
		want         *domain.NewTermsOfServiceStatusArgs
		wantErr      error
	}{
		"when_not_found": {
			loadFixtures: false,
			want:         nil,
			wantErr:      database.ErrNotFound,
		},
		"happy_flow": {
			loadFixtures: true,
			want: &domain.NewTermsOfServiceStatusArgs{
				Username:  "fixture-username",
				CreatedAt: fixtureTime,
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

			got, err := configDb.GetTermsOfServiceStatus(context.Background())
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				want, err := domain.NewTermsOfServiceStatus(tt.want)
				require.NoError(t, err)
				assert.Equal(t, want, got)
			}
		})
	}
}

func TestAcceptTermsOfService(t *testing.T) {
	t.Parallel()

	setup(t)

	fixtureTime := getCustomFixtureTime(t, "2021-01-04T01:02:03Z")

	type args struct {
		username  string
		createdAt time.Time
	}

	tests := map[string]struct {
		loadFixtures bool
		args         *args
		wantErr      error
	}{
		"when_created_at_in_future": {
			loadFixtures: false,
			args: &args{
				username:  "test-username",
				createdAt: time.Now().Add(time.Hour),
			},
			wantErr: database.ErrInvalidDate,
		},
		"happy_flow": {
			loadFixtures: false,
			args: &args{
				username:  "test-username",
				createdAt: fixtureTime,
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

			err := configDb.AcceptTermsOfService(context.Background(), tt.args.username, tt.args.createdAt)
			require.ErrorIs(t, err, tt.wantErr)

			model, err := domain.NewTermsOfServiceStatus(&domain.NewTermsOfServiceStatusArgs{
				Username:  tt.args.username,
				CreatedAt: tt.args.createdAt,
			})

			if tt.wantErr == nil {
				assertTermsOfServiceInRepository(t, configDb, model)
			}
		})
	}
}

func assertTermsOfServiceInRepository(t *testing.T, repo database.ConfigDatabase, want *domain.TermsOfServiceStatus) {
	got, err := repo.GetTermsOfServiceStatus(context.Background())
	require.NoError(t, err)

	assert.Equal(t, want.Username(), got.Username())
	assert.Equal(t, want.CreatedAt(), got.CreatedAt())
}
