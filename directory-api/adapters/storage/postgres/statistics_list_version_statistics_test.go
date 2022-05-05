// Copyright © VNG Realisatie 2022
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

func TestListVersionStatistics(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		loadFixtures bool
		want         []*domain.NewVersionStatisticsArgs
		wantErr      error
	}{
		"when_no_versions_registered": {
			loadFixtures: false,
			want:         nil,
			wantErr:      nil,
		},
		"happy_flow": {
			loadFixtures: true,
			want: []*domain.NewVersionStatisticsArgs{
				{
					GatewayType: "inway",
					Version:     "1.0.0",
					Amount:      5,
				},
				{
					GatewayType: "outway",
					Version:     "2.0.0",
					Amount:      1,
				},
				{
					GatewayType: "outway",
					Version:     "1.0.0",
					Amount:      4,
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

			got, err := repo.ListVersionStatistics(context.Background())
			require.Equal(t, tt.wantErr, err)

			wantModels := make([]*domain.VersionStatistics, len(tt.want))
			for i, v := range tt.want {
				wantModels[i], err = domain.NewVersionStatistics(&domain.NewVersionStatisticsArgs{
					GatewayType: v.GatewayType,
					Version:     v.Version,
					Amount:      v.Amount,
				})
			}

			if tt.wantErr == nil {
				assert.EqualValues(t, wantModels, got)
			}
		})
	}
}
