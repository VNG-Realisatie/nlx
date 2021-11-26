// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration

package storage_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/common/nlxversion"
	"go.nlx.io/nlx/directory-api/domain"
)

func TestRegisterOutway(t *testing.T) {
	t.Parallel()

	type args struct {
		version nlxversion.Version
	}

	tests := map[string]struct {
		loadFixtures bool
		args         *domain.NewOutwayArgs
		wantErr      error
	}{
		"happy_flow": {
			loadFixtures: false,
			args: &domain.NewOutwayArgs{
				Name:         "fixture-outway-name-one",
				Organization: createNewOrganization(t, "org-a", "00000000000000000001"),
				NlxVersion:   "1.0.0",
				CreatedAt:    time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
				UpdatedAt:    time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
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

			outway, _ := domain.NewOutway(tt.args)

			err := repo.RegisterOutway(outway)
			require.Equal(t, tt.wantErr, err)

			assertOutwayInRepository(t, repo, outway)
		})
	}
}
