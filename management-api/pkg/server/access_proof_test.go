// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
)

//nolint:funlen // its a unittest
func TestGetAccessProof(t *testing.T) {
	now := time.Now()
	ts, _ := types.TimestampProto(now)

	tests := map[string]struct {
		want     *api.AccessProof
		wantCode codes.Code
		setup    func(*mock_database.MockConfigDatabase) context.Context
	}{
		"errors_when_peer_context_is_missing": {
			wantCode: codes.Internal,
			setup: func(db *mock_database.MockConfigDatabase) context.Context {
				return context.Background()
			},
		},

		"returns_error_when_get_latest_access_grant_for_service_errors": {
			wantCode: codes.Internal,
			setup: func(db *mock_database.MockConfigDatabase) context.Context {
				ctx := setProxyMetadata(context.Background())

				db.
					EXPECT().
					GetLatestAccessGrantForService(ctx, "organization-a", "service").
					Return(nil, errors.New("error"))

				return ctx
			},
		},

		"returns_not_found_when_access_grant_could_not_be_found": {
			wantCode: codes.NotFound,
			setup: func(db *mock_database.MockConfigDatabase) context.Context {
				ctx := setProxyMetadata(context.Background())

				db.
					EXPECT().
					GetLatestAccessGrantForService(ctx, "organization-a", "service").
					Return(nil, database.ErrNotFound)

				return ctx
			},
		},

		"returns_error_when_grant_created_at_is_invalid": {
			wantCode: codes.Internal,
			setup: func(db *mock_database.MockConfigDatabase) context.Context {
				ctx := setProxyMetadata(context.Background())

				db.
					EXPECT().
					GetLatestAccessGrantForService(ctx, "organization-a", "service").
					Return(&database.AccessGrant{
						CreatedAt: time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC),
						RevokedAt: time.Now(),
					}, nil)

				return ctx
			},
		},

		"returns_error_when_grant_revoked_at_is_invalid": {
			wantCode: codes.Internal,
			setup: func(db *mock_database.MockConfigDatabase) context.Context {
				ctx := setProxyMetadata(context.Background())

				db.
					EXPECT().
					GetLatestAccessGrantForService(ctx, "organization-a", "service").
					Return(&database.AccessGrant{
						CreatedAt: time.Now(),
						RevokedAt: time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC),
					}, nil)

				return ctx
			},
		},

		"returns_access_proof_for_successful_request": {
			wantCode: codes.OK,
			want: &api.AccessProof{
				Id:               "id",
				CreatedAt:        ts,
				RevokedAt:        ts,
				AccessRequestId:  "request-1",
				OrganizationName: "organization-a",
				ServiceName:      "service",
			},
			setup: func(db *mock_database.MockConfigDatabase) context.Context {
				ctx := setProxyMetadata(context.Background())

				db.
					EXPECT().
					GetLatestAccessGrantForService(ctx, "organization-a", "service").
					Return(&database.AccessGrant{
						CreatedAt:        now,
						RevokedAt:        now,
						AccessRequestID:  "request-1",
						ID:               "id",
						OrganizationName: "organization-a",
						ServiceName:      "service",
					}, nil)

				return ctx
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, db := newService(t)
			ctx := tt.setup(db)

			response, err := service.GetAccessProof(ctx, &external.GetAccessProofRequest{
				ServiceName: "service",
			})

			if tt.wantCode > 0 {
				assert.Error(t, err)

				st, ok := status.FromError(err)

				assert.True(t, ok)
				assert.Equal(t, tt.wantCode, st.Code())
			} else {
				assert.NoError(t, err)

				if assert.NotNil(t, response) {
					assert.Equal(t, tt.want, response)
				}
			}
		})
	}
}
