// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package inspectionservice_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"go.nlx.io/nlx/directory-inspection-api/inspectionapi"
	"go.nlx.io/nlx/directory-inspection-api/pkg/database"
	"go.nlx.io/nlx/directory-inspection-api/pkg/database/mock"
	"go.nlx.io/nlx/directory-inspection-api/pkg/inspectionservice"
)

func TestInspectionService_ListVersionStatistics(t *testing.T) {
	tests := []struct {
		name             string
		db               func(ctrl *gomock.Controller) database.DirectoryDatabase
		expectedResponse *inspectionapi.ListInOutwayStatisticsResponse
		expectedError    error
	}{
		{
			name: "failed to get statistics from the db",
			db: func(ctrl *gomock.Controller) database.DirectoryDatabase {
				db := mock.NewMockDirectoryDatabase(ctrl)
				db.EXPECT().ListVersionStatistics(gomock.Any()).Return(nil, errors.New("arbitrary error"))

				return db
			},
			expectedResponse: nil,
			expectedError:    status.New(codes.Internal, "Database error.").Err(),
		},
		{
			name: "for an type which is not supported",
			db: func(ctrl *gomock.Controller) database.DirectoryDatabase {
				db := mock.NewMockDirectoryDatabase(ctrl)
				db.EXPECT().ListVersionStatistics(gomock.Any()).Return([]*database.VersionStatistics{
					{
						Type:    "arbitrary",
						Amount:  5,
						Version: "0.0.1",
					},
				}, nil)

				return db
			},
			expectedResponse: &inspectionapi.ListInOutwayStatisticsResponse{},
			expectedError:    nil,
		},
		{
			name: "happy flow",
			db: func(ctrl *gomock.Controller) database.DirectoryDatabase {
				db := mock.NewMockDirectoryDatabase(ctrl)
				db.EXPECT().ListVersionStatistics(gomock.Any()).Return([]*database.VersionStatistics{
					{
						Type:    "outway",
						Amount:  5,
						Version: "0.0.1",
					},
					{
						Type:    "outway",
						Amount:  10,
						Version: "0.0.2",
					},
					{
						Type:    "inway",
						Amount:  5,
						Version: "0.0.2",
					},
				}, nil)

				return db
			},
			expectedResponse: &inspectionapi.ListInOutwayStatisticsResponse{
				Versions: []*inspectionapi.ListInOutwayStatisticsResponse_Statistics{
					{
						Type:    inspectionapi.ListInOutwayStatisticsResponse_Statistics_OUTWAY,
						Amount:  5,
						Version: "0.0.1",
					},
					{
						Type:    inspectionapi.ListInOutwayStatisticsResponse_Statistics_OUTWAY,
						Amount:  10,
						Version: "0.0.2",
					},
					{
						Type:    inspectionapi.ListInOutwayStatisticsResponse_Statistics_INWAY,
						Amount:  5,
						Version: "0.0.2",
					},
				},
			},
			expectedError: nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			h := inspectionservice.New(zap.NewNop(), tt.db(ctrl), testGetOrganizationNameFromRequest)
			got, err := h.ListVersionStatistics(context.Background(), &emptypb.Empty{})

			assert.Equal(t, tt.expectedResponse, got)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
