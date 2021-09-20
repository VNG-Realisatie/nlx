// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration
// +build integration

package database_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/directory-inspection-api/pkg/database"
)

func TestListOrganizations(t *testing.T) {
	t.Parallel()

	setup(t)

	tests := map[string]struct {
		loadFixtures bool
		want         []*database.Organization
		wantErr      error
	}{
		"when_no_results": {
			loadFixtures: false,
			want:         nil,
			wantErr:      nil,
		},
		"happy_flow": {
			loadFixtures: true,
			want: []*database.Organization{
				{
					Name: "fixture-organization-name",
				},
				{
					Name: "fixture-second-organization-name",
				},
			},
			wantErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			db, close := newDirectoryDatabase(t, t.Name(), tt.loadFixtures)
			defer close()

			got, err := db.ListOrganizations(context.Background())
			require.Equal(t, tt.wantErr, err)

			if tt.wantErr == nil {
				assert.EqualValues(t, tt.want, got)
			}
		})
	}
}

func TestGetOrganizationInwayAddress(t *testing.T) {
	t.Parallel()

	setup(t)

	type args struct {
		organizationName string
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		want         string
		wantErr      error
	}{
		"when_organization_does_not_exist": {
			loadFixtures: true,
			args: args{
				organizationName: "arbitrary organization name",
			},
			want:    "",
			wantErr: database.ErrNoOrganization,
		},
		"happy_flow": {
			loadFixtures: true,
			args: args{
				organizationName: "fixture-organization-name",
			},
			want:    "https://fixture-inway-address.com",
			wantErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			db, close := newDirectoryDatabase(t, t.Name(), tt.loadFixtures)
			defer close()

			got, err := db.GetOrganizationInwayAddress(context.Background(), tt.args.organizationName)
			require.Equal(t, tt.wantErr, err)

			if tt.wantErr == nil {
				assert.EqualValues(t, tt.want, got)
			}
		})
	}
}
