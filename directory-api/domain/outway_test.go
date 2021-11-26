// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package domain_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/directory-api/domain"
)

//nolint:funlen // this is a test
func TestNewOutway(t *testing.T) {
	now := time.Now()

	tests := map[string]struct {
		outwayArgs  *domain.NewOutwayArgs
		expectedErr string
	}{
		"invalid_name": {
			outwayArgs: &domain.NewOutwayArgs{
				Name:         "#*%",
				Organization: createNewOrganization(t),
				NlxVersion:   "0.0.0",
				CreatedAt:    now,
				UpdatedAt:    now,
			},
			expectedErr: "Name: must be in a valid format.",
		},
		"invalid_version": {
			outwayArgs: &domain.NewOutwayArgs{
				Name:         "name",
				Organization: createNewOrganization(t),
				NlxVersion:   "invalid",
				CreatedAt:    now,
				UpdatedAt:    now,
			},
			expectedErr: "NlxVersion: must be a valid semantic version.",
		},
		"without_organization": {
			outwayArgs: &domain.NewOutwayArgs{
				Name:         "name",
				Organization: nil,
				NlxVersion:   "0.0.0",
				CreatedAt:    now,
				UpdatedAt:    now,
			},
			expectedErr: "Organization: is required.",
		},
		"empty_name": {
			outwayArgs: &domain.NewOutwayArgs{
				Name:         "",
				Organization: createNewOrganization(t),
				NlxVersion:   "0.0.0",
				CreatedAt:    now,
				UpdatedAt:    now,
			},
			expectedErr: "",
		},
		"created_at_in_future": {
			outwayArgs: &domain.NewOutwayArgs{
				Name:         "name",
				Organization: createNewOrganization(t),
				NlxVersion:   "0.0.0",
				CreatedAt:    now.Add(1 * time.Hour),
				UpdatedAt:    now,
			},
			expectedErr: "CreatedAt: must not be in the future.",
		},
		"updated_at_in_future": {
			outwayArgs: &domain.NewOutwayArgs{
				Name:         "name",
				Organization: createNewOrganization(t),
				NlxVersion:   "0.0.0",
				CreatedAt:    now,
				UpdatedAt:    now.Add(1 * time.Hour),
			},
			expectedErr: "UpdatedAt: must not be in the future.",
		},
		"happy_flow": {
			outwayArgs: &domain.NewOutwayArgs{
				Name:         "name",
				Organization: createNewOrganization(t),
				NlxVersion:   "0.0.0",
				CreatedAt:    now,
				UpdatedAt:    now,
			},
			expectedErr: "",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			result, err := domain.NewOutway(tt.outwayArgs)

			if tt.expectedErr != "" {
				assert.Nil(t, result)
				assert.EqualError(t, err, tt.expectedErr)
			} else {
				assert.NotNil(t, result)
				assert.Nil(t, err)

				assert.Equal(t, tt.outwayArgs.Name, result.Name())
				assert.Equal(t, tt.outwayArgs.Organization.Name(), result.Organization().Name())
				assert.Equal(t, tt.outwayArgs.NlxVersion, result.NlxVersion())
			}
		})
	}
}
