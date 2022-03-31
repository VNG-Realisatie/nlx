// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package directory_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/directory-api/domain"
)

func TestListParticipants(t *testing.T) {
	now := time.Now()

	tests := map[string]struct {
		setup            func(serviceMocks)
		expectedResponse *directoryapi.ListParticipantsResponse
		expectedError    error
	}{
		"database_error": {
			setup: func(mocks serviceMocks) {
				mocks.repository.
					EXPECT().
					ListParticipants(gomock.Any()).
					Return(nil, errors.New("arbitrary error"))
			},
			expectedResponse: nil,
			expectedError:    status.New(codes.Internal, "Database error.").Err(),
		},
		"happy_flow": {
			setup: func(mocks serviceMocks) {
				organizationA, _ := domain.NewOrganization("org-a", "00000000000000000001")
				organizationB, _ := domain.NewOrganization("org-b", "00000000000000000002")

				models := []*domain.NewParticipantArgs{
					{
						Organization: organizationA,
						Statistics: &domain.NewParticipantStatisticsArgs{
							Inways:   1,
							Outways:  2,
							Services: 3,
						},
						CreatedAt: now,
					},
					{
						Organization: organizationB,
						Statistics: &domain.NewParticipantStatisticsArgs{
							Inways:   33,
							Outways:  22,
							Services: 11,
						},
						CreatedAt: now,
					},
				}

				mocks.repository.
					EXPECT().
					ListParticipants(gomock.Any()).
					Return(createParticipants(models), nil)
			},
			expectedResponse: &directoryapi.ListParticipantsResponse{
				Participants: []*directoryapi.ListParticipantsResponse_Participant{
					{
						Organization: &directoryapi.Organization{
							SerialNumber: "00000000000000000001",
							Name:         "org-a",
						},
						Statistics: &directoryapi.ListParticipantsResponse_Participant_Statistics{
							Inways:   1,
							Outways:  2,
							Services: 3,
						},
						CreatedAt: timestamppb.New(now),
					},
					{
						Organization: &directoryapi.Organization{
							SerialNumber: "00000000000000000002",
							Name:         "org-b",
						},
						Statistics: &directoryapi.ListParticipantsResponse_Participant_Statistics{
							Inways:   33,
							Outways:  22,
							Services: 11,
						},
						CreatedAt: timestamppb.New(now),
					},
				},
			},
			expectedError: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, mocks := newService(t, "", &testClock{
				timeToReturn: time.Now(),
			})

			if tt.setup != nil {
				tt.setup(mocks)
			}

			got, err := service.ListParticipants(context.Background(), &emptypb.Empty{})

			assert.Equal(t, tt.expectedResponse, got)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func createParticipants(models []*domain.NewParticipantArgs) []*domain.Participant {
	participants := make([]*domain.Participant, len(models))

	for i, p := range models {
		participants[i], _ = domain.NewParticipant(p)
	}

	return participants
}
