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
func TestNewParticipant(t *testing.T) {
	now := time.Now()

	tests := map[string]struct {
		args        *domain.NewParticipantArgs
		expectedErr string
	}{
		"without_organization": {
			args: &domain.NewParticipantArgs{
				Organization: nil,
				Statistics: &domain.NewParticipantStatisticsArgs{
					Inways:   1,
					Outways:  2,
					Services: 3,
				},
				CreatedAt: now,
			},
			expectedErr: "Organization: is required.",
		},
		"without_statistics": {
			args: &domain.NewParticipantArgs{
				Organization: createNewOrganization(t),
				Statistics:   nil,
				CreatedAt:    now,
			},
			expectedErr: "Statistics: is required.",
		},
		"created_at_in_future": {
			args: &domain.NewParticipantArgs{
				Organization: createNewOrganization(t),
				Statistics: &domain.NewParticipantStatisticsArgs{
					Inways:   1,
					Outways:  2,
					Services: 3,
				},
				CreatedAt: now.Add(1 * time.Hour),
			},
			expectedErr: "CreatedAt: must not be in the future.",
		},
		"happy_flow": {
			args: &domain.NewParticipantArgs{
				Organization: createNewOrganization(t),
				Statistics: &domain.NewParticipantStatisticsArgs{
					Inways:   1,
					Outways:  2,
					Services: 3,
				},
				CreatedAt: now,
			},
			expectedErr: "",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			result, err := domain.NewParticipant(tt.args)

			if tt.expectedErr != "" {
				assert.Nil(t, result)
				assert.EqualError(t, err, tt.expectedErr)
			} else {
				assert.NotNil(t, result)
				assert.Nil(t, err)

				assert.Equal(t, tt.args.Organization.Name(), result.Organization().Name())
				assert.Equal(t, tt.args.Statistics.Inways, result.Statistics().Inways())
				assert.Equal(t, tt.args.Statistics.Outways, result.Statistics().Outways())
				assert.Equal(t, tt.args.Statistics.Services, result.Statistics().Services())
				assert.Equal(t, tt.args.CreatedAt, result.CreatedAt())
			}
		})
	}
}
