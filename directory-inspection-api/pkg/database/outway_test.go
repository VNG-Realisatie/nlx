// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration
// +build integration

package database_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/common/nlxversion"
	"go.nlx.io/nlx/directory-inspection-api/pkg/database"
)

func TestRegisterOutwayVersion(t *testing.T) {
	t.Parallel()

	setup(t)

	type args struct {
		version nlxversion.Version
	}

	tests := map[string]struct {
		args    args
		wantErr error
	}{
		"happy_flow": {
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

			db, close := newDirectoryDatabase(t, t.Name())
			defer close()

			err := db.RegisterOutwayVersion(context.Background(), tt.args.version)
			require.Equal(t, tt.wantErr, err)

			assertVersionInRepository(t, db, nlxversion.Version{
				Version:   tt.args.version.Version,
				Component: "outway",
			})
		})
	}
}

func assertVersionInRepository(t *testing.T, db database.DirectoryDatabase, version nlxversion.Version) {
	statistics, err := db.ListVersionStatistics(context.Background())
	require.NoError(t, err)

	var result = false

	for _, statistic := range statistics {
		if statistic.Version == version.Version && string(statistic.Type) == version.Component {
			result = true
		}
	}

	require.True(t, result)
}
