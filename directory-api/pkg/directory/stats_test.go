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

	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/directory-api/domain"
)

func TestListInOutwayStatistics(t *testing.T) {
	tests := map[string]struct {
		setup            func(serviceMocks)
		expectedResponse *directoryapi.ListInOutwayStatisticsResponse
		expectedError    error
	}{
		"failed_to_get_statistics_from_the_db": {
			setup: func(s serviceMocks) {
				s.repository.EXPECT().ListVersionStatistics(gomock.Any()).Return(nil, errors.New("arbitrary error"))
			},
			expectedResponse: nil,
			expectedError:    status.New(codes.Internal, "Storage error.").Err(),
		},
		"happy flow": {
			setup: func(s serviceMocks) {
				versionStatisticsOutway, _ := domain.NewVersionStatistics(&domain.NewVersionStatisticsArgs{
					GatewayType: domain.TypeOutway,
					Amount:      5,
					Version:     "0.0.1",
				})

				versionStatisticsInway, _ := domain.NewVersionStatistics(&domain.NewVersionStatisticsArgs{
					GatewayType: domain.TypeInway,
					Amount:      20,
					Version:     "0.0.3",
				})

				s.repository.EXPECT().ListVersionStatistics(gomock.Any()).Return([]*domain.VersionStatistics{
					versionStatisticsOutway,
					versionStatisticsInway,
				}, nil)
			},
			expectedResponse: &directoryapi.ListInOutwayStatisticsResponse{
				Versions: []*directoryapi.ListInOutwayStatisticsResponse_Statistics{
					{
						Type:    directoryapi.ListInOutwayStatisticsResponse_Statistics_TYPE_OUTWAY,
						Amount:  5,
						Version: "0.0.1",
					},
					{
						Type:    directoryapi.ListInOutwayStatisticsResponse_Statistics_TYPE_INWAY,
						Amount:  20,
						Version: "0.0.3",
					},
				},
			},
			expectedError: nil,
		},
	}
	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, mocks := newService(t, testNlxVersion128, "", &testClock{
				timeToReturn: time.Now(),
			})

			tt.setup(mocks)

			got, err := service.ListInOutwayStatistics(context.Background(), &directoryapi.ListInOutwayStatisticsRequest{})

			assert.Equal(t, tt.expectedResponse, got)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
