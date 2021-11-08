// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration

package storage_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/common/nlxversion"
	"go.nlx.io/nlx/directory-api/domain/directory/storage"
)

func TestRegisterOutwayVersion(t *testing.T) {
	t.Parallel()

	type args struct {
		version nlxversion.Version
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		wantErr      error
	}{
		"happy_flow": {
			loadFixtures: false,
			args: args{
				version: nlxversion.Version{
					Version: "1.0.0",
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

			err := repo.RegisterOutwayVersion(context.Background(), tt.args.version)
			require.Equal(t, tt.wantErr, err)

			assertVersionInRepository(t, repo, nlxversion.Version{
				Version:   tt.args.version.Version,
				Component: "outway",
			})
		})
	}
}

func assertVersionInRepository(t *testing.T, r storage.Repository, version nlxversion.Version) {
	statistics, err := r.ListVersionStatistics(context.Background())
	require.NoError(t, err)

	var result = false

	for _, statistic := range statistics {
		if statistic.Version() == version.Version && string(statistic.GatewayType()) == version.Component {
			result = true
		}
	}

	require.True(t, result)
}
