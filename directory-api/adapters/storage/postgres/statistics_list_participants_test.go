// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

//go:build integration

package pgadapter_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/directory-api/domain"
)

func TestListParticipants(t *testing.T) {
	t.Parallel()

	type wantOrganization struct {
		serialNumber string
		name         string
	}

	tests := map[string]struct {
		loadFixtures bool
		want         []*domain.NewParticipantArgs
		wantErr      error
	}{
		"when_no_participants": {
			loadFixtures: false,
			want:         nil,
			wantErr:      nil,
		},
		"happy_flow": {
			loadFixtures: true,
			want: []*domain.NewParticipantArgs{
				{
					Organization: createNewOrganization(t, "duplicate-org-name", "11111111111111111111"),
					Statistics: &domain.NewParticipantStatisticsArgs{
						Inways:   1,
						Outways:  0,
						Services: 0,
					},
					CreatedAt: time.Now(),
				},
				{
					Organization: createNewOrganization(t, "fixture-organization-name", "01234567890123456789"),
					Statistics: &domain.NewParticipantStatisticsArgs{
						Inways:   2,
						Outways:  1,
						Services: 1,
					},
					CreatedAt: time.Now(),
				},
				{
					Organization: createNewOrganization(t, "fixture-second-organization-name", "01234567890123456781"),
					Statistics: &domain.NewParticipantStatisticsArgs{
						Inways:   1,
						Outways:  3,
						Services: 0,
					},
					CreatedAt: time.Now(),
				},
				{
					Organization: createNewOrganization(t, "duplicate-org-name", "22222222222222222222"),
					Statistics: &domain.NewParticipantStatisticsArgs{
						Inways:   1,
						Outways:  1,
						Services: 0,
					},
					CreatedAt: time.Now(),
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

			want := make([]*domain.Participant, len(tt.want))

			for i, s := range tt.want {
				var err error
				want[i], err = domain.NewParticipant(s)
				require.NoError(t, err)
			}

			gotList, err := repo.ListParticipants(context.Background())
			require.Equal(t, tt.wantErr, err)

			if tt.wantErr == nil {
				for i, got := range gotList {
					assert.WithinDuration(t, want[i].CreatedAt(), got.CreatedAt(), time.Minute*5)

					assert.Equal(t, want[i].Organization().SerialNumber(), got.Organization().SerialNumber())
					assert.Equal(t, want[i].Organization().Name(), got.Organization().Name())
					assert.Equal(t, want[i].Statistics().Inways(), got.Statistics().Inways())
					assert.Equal(t, want[i].Statistics().Outways(), got.Statistics().Outways())
					assert.Equal(t, want[i].Statistics().Services(), got.Statistics().Services())
				}
			}
		})
	}
}
