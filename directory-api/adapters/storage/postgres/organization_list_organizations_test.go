// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration

package pgadapter_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/directory-api/domain"
)

func TestListOrganizations(t *testing.T) {
	t.Parallel()

	type wantOrganization struct {
		serialNumber string
		name         string
	}

	tests := map[string]struct {
		loadFixtures bool
		want         []*wantOrganization
		wantErr      error
	}{
		"when_no_organizations": {
			loadFixtures: false,
			want:         nil,
			wantErr:      nil,
		},
		"happy_flow": {
			loadFixtures: true,
			want: []*wantOrganization{
				{
					serialNumber: "11111111111111111111",
					name:         "duplicate-org-name",
				},
				{
					serialNumber: "22222222222222222222",
					name:         "duplicate-org-name",
				},
				{
					serialNumber: "01234567890123456789",
					name:         "fixture-organization-name",
				},
				{
					serialNumber: "01234567890123456781",
					name:         "fixture-second-organization-name",
				},
			},
			wantErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			repo, close := new(t, tt.loadFixtures)
			defer close()

			want := make([]*domain.Organization, len(tt.want))

			for i, s := range tt.want {
				var err error
				want[i], err = domain.NewOrganization(s.name, s.serialNumber)
				require.NoError(t, err)
			}

			got, err := repo.ListOrganizations(context.Background())
			require.Equal(t, tt.wantErr, err)

			if tt.wantErr == nil {
				assert.EqualValues(t, want, got)
			}
		})
	}
}
