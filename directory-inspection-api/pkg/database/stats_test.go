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

func TestListVersionStatistics(t *testing.T) {
	t.Parallel()

	setup(t)

	tests := map[string]struct {
		loadFixtures bool
		want         []*database.VersionStatistics
		wantErr      error
	}{
		"when_no_versions_registered": {
			loadFixtures: false,
			want:         nil,
			wantErr:      nil,
		},
		"happy_flow": {
			loadFixtures: true,
			want: []*database.VersionStatistics{
				{
					Type:    "inway",
					Version: "1.0.0",
					Amount:  1,
				},
				{
					Type:    "outway",
					Version: "2.0.0",
					Amount:  1,
				},
				{
					Type:    "outway",
					Version: "1.0.0",
					Amount:  2,
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

			got, err := db.ListVersionStatistics(context.Background())
			require.Equal(t, tt.wantErr, err)

			if tt.wantErr == nil {
				assert.EqualValues(t, tt.want, got)
			}
		})
	}
}
