// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

// +build integration

package database_test

import (
	"context"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/management-api/pkg/database"
)

func TestPutOrganizationInway(t *testing.T) {
	t.Parallel()

	setup(t)

	tests := map[string]struct {
		inwayID     *uint
		expected    *database.Settings
		expectedErr error
	}{
		"non_existing_inway": {
			inwayID:     newUint(9999999),
			expected:    nil,
			expectedErr: database.ErrInwayNotFound,
		},
		"happy_flow": {
			inwayID: newUint(1),
			expected: &database.Settings{
				InwayID: newUint(1),
			},
			expectedErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			configDb, close := newConfigDatabase(t, t.Name())
			defer close()

			actual, err := configDb.PutOrganizationInway(context.Background(), tt.inwayID)
			require.ErrorIs(t, err, tt.expectedErr)

			if tt.expectedErr == nil {
				require.Equal(t, tt.expected.InwayID, actual.InwayID)
			}
		})
	}
}

func newUint(x uint) *uint {
	return &x
}
