// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server_test

import (
	context "context"
	"errors"
	"testing"
	"time"

	"github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
)

func TestListAccessGrantsForService(t *testing.T) {
	createTimestamp := func(ti time.Time) *types.Timestamp {
		return &types.Timestamp{
			Seconds: ti.Unix(),
			Nanos:   int32(ti.Nanosecond()),
		}
	}

	tests := []struct {
		name             string
		req              *api.ListAccessGrantsForServiceRequest
		returnReq        []*database.AccessGrant
		returnErr        error
		expectedResponse *api.ListAccessGrantsForServiceResponse
		expectedErr      error
	}{
		{
			"happy_flow",
			&api.ListAccessGrantsForServiceRequest{
				ServiceName: "test-service",
			},
			[]*database.AccessGrant{
				{
					ID:                   "12345abcde",
					OrganizationName:     "test-organization",
					ServiceName:          "test-service",
					PublicKeyFingerprint: "test-finger-print",
					CreatedAt:            time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC),
				},
			},
			nil,
			&api.ListAccessGrantsForServiceResponse{
				AccessGrants: []*api.AccessGrant{
					{
						Id:                   "12345abcde",
						OrganizationName:     "test-organization",
						ServiceName:          "test-service",
						PublicKeyFingerprint: "test-finger-print",
						CreatedAt:            createTimestamp(time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC)),
					},
				},
			},

			nil,
		},
		{
			"convert_access_grant_error",
			&api.ListAccessGrantsForServiceRequest{
				ServiceName: "test-service",
			},
			[]*database.AccessGrant{
				{
					ID:                   "12345abcde",
					OrganizationName:     "test-organization",
					ServiceName:          "test-service",
					PublicKeyFingerprint: "test-finger-print",
					CreatedAt:            time.Date(0, time.January, 0, 0, 0, 0, 0, time.UTC),
				},
			},
			nil,
			nil,
			status.Error(codes.Internal, "error converting access grant"),
		},
		{
			"database_error",
			&api.ListAccessGrantsForServiceRequest{
				ServiceName: "test-service",
			},
			nil,
			errors.New("arbitrary error"),
			nil,
			status.Error(codes.Internal, "database error"),
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			service, ctrl, db := newService(t)
			defer ctrl.Finish()

			ctx := context.Background()

			db.EXPECT().ListAccessGrantsForService(ctx, tt.req.ServiceName).
				Return(tt.returnReq, tt.returnErr)
			actual, err := service.ListAccessGrantsForService(ctx, tt.req)

			assert.Equal(t, tt.expectedResponse, actual)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
