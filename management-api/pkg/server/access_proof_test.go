// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	"go.nlx.io/nlx/management-api/pkg/server"
)

//nolint:funlen // its a unittest
func TestGetAccessProof(t *testing.T) {
	now := time.Now()
	ts, _ := ptypes.TimestampProto(now)

	tests := map[string]struct {
		want    *api.AccessProof
		wantErr error
		setup   func(*mock_database.MockConfigDatabase) context.Context
	}{
		"errors_when_peer_context_is_missing": {
			wantErr: status.Error(codes.Internal, "missing metadata from the management proxy"),
			setup: func(db *mock_database.MockConfigDatabase) context.Context {
				return context.Background()
			},
		},

		"returns_error_when_get_latest_access_grant_for_service_errors": {
			wantErr: status.Error(codes.Internal, "database error"),
			setup: func(db *mock_database.MockConfigDatabase) context.Context {
				ctx := setProxyMetadata(context.Background())

				db.
					EXPECT().
					GetService(ctx, "service").
					Return(&database.Service{Name: "service"}, nil)
				db.
					EXPECT().
					GetLatestAccessGrantForService(ctx, "organization-a", "service").
					Return(nil, errors.New("error"))

				return ctx
			},
		},

		"returns_not_found_when_access_grant_could_not_be_found": {
			wantErr: status.Error(codes.NotFound, "access proof not found"),
			setup: func(db *mock_database.MockConfigDatabase) context.Context {
				ctx := setProxyMetadata(context.Background())

				db.
					EXPECT().
					GetService(ctx, "service").
					Return(&database.Service{Name: "service"}, nil)

				db.
					EXPECT().
					GetLatestAccessGrantForService(ctx, "organization-a", "service").
					Return(nil, database.ErrNotFound)

				return ctx
			},
		},

		"returns_not_found_when_service_no_long_exists": {
			wantErr: server.ErrServiceDoesNotExist,
			setup: func(db *mock_database.MockConfigDatabase) context.Context {
				ctx := setProxyMetadata(context.Background())

				db.
					EXPECT().
					GetService(ctx, "service").
					Return(nil, database.ErrNotFound)

				return ctx
			},
		},

		"returns_access_proof_for_successful_request": {
			want: &api.AccessProof{
				Id:               1,
				CreatedAt:        ts,
				RevokedAt:        ts,
				AccessRequestId:  1,
				OrganizationName: "organization-a",
				ServiceName:      "service",
			},
			setup: func(db *mock_database.MockConfigDatabase) context.Context {
				ctx := setProxyMetadata(context.Background())

				db.
					EXPECT().
					GetService(ctx, "service").
					Return(&database.Service{Name: "service"}, nil)

				db.
					EXPECT().
					GetLatestAccessGrantForService(ctx, "organization-a", "service").
					Return(&database.AccessGrant{
						CreatedAt:               now,
						RevokedAt:               sql.NullTime{Time: now},
						ID:                      1,
						IncomingAccessRequestID: 1,
						IncomingAccessRequest: &database.IncomingAccessRequest{
							ID:               1,
							OrganizationName: "organization-a",
							ServiceID:        1,
							Service: &database.Service{
								Name: "service",
							},
						},
					}, nil)

				return ctx
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, _, mocks := newService(t)
			ctx := tt.setup(mocks.db)

			response, err := service.GetAccessProof(ctx, &external.GetAccessProofRequest{
				ServiceName: "service",
			})

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, response)
		})
	}
}
