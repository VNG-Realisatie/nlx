// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	"go.nlx.io/nlx/management-api/pkg/server"
)

//nolint:funlen // its a unittest
func TestGetAccessProof(t *testing.T) {
	now := time.Now()
	ts := timestamppb.New(now)

	tests := map[string]struct {
		setup   func(*testing.T, *mock_database.MockConfigDatabase, *tls.CertificateBundle) context.Context
		want    *external.AccessProof
		wantErr error
	}{
		"when_the_peer_context_is_missing": {
			setup: func(*testing.T, *mock_database.MockConfigDatabase, *tls.CertificateBundle) context.Context {
				return context.Background()
			},
			wantErr: status.Error(codes.Internal, "missing metadata from the management proxy"),
		},
		"when_getting_the_latest_access_grant_for_the_service_errors": {
			setup: func(t *testing.T, db *mock_database.MockConfigDatabase, certBundle *tls.CertificateBundle) context.Context {
				ctx := setProxyMetadataWithCertBundle(t, context.Background(), certBundle)

				db.
					EXPECT().
					GetService(ctx, "service").
					Return(&database.Service{Name: "service"}, nil)
				db.
					EXPECT().
					GetLatestAccessGrantForService(ctx, certBundle.Certificate().Subject.SerialNumber, "service", "public-key-fingerprint").
					Return(nil, errors.New("error"))

				return ctx
			},
			wantErr: status.Error(codes.Internal, "database error"),
		},
		"when_access_grant_could_not_be_found": {
			setup: func(t *testing.T, db *mock_database.MockConfigDatabase, certBundle *tls.CertificateBundle) context.Context {
				ctx := setProxyMetadataWithCertBundle(t, context.Background(), certBundle)

				db.
					EXPECT().
					GetService(ctx, "service").
					Return(&database.Service{Name: "service"}, nil)

				db.
					EXPECT().
					GetLatestAccessGrantForService(ctx, certBundle.Certificate().Subject.SerialNumber, "service", "public-key-fingerprint").
					Return(nil, database.ErrNotFound)

				return ctx
			},
			wantErr: status.Error(codes.NotFound, "access proof not found"),
		},
		"when_the_service_no_long_exists": {
			wantErr: server.ErrServiceDoesNotExist,
			setup: func(t *testing.T, db *mock_database.MockConfigDatabase, certBundle *tls.CertificateBundle) context.Context {
				ctx := setProxyMetadataWithCertBundle(t, context.Background(), certBundle)

				db.
					EXPECT().
					GetService(ctx, "service").
					Return(nil, database.ErrNotFound)

				return ctx
			},
		},
		"happy_flow": {
			setup: func(t *testing.T, db *mock_database.MockConfigDatabase, certBundle *tls.CertificateBundle) context.Context {
				ctx := setProxyMetadataWithCertBundle(t, context.Background(), certBundle)

				db.
					EXPECT().
					GetService(ctx, "service").
					Return(&database.Service{Name: "service"}, nil)

				db.
					EXPECT().
					GetLatestAccessGrantForService(ctx, certBundle.Certificate().Subject.SerialNumber, "service", "public-key-fingerprint").
					Return(&database.AccessGrant{
						CreatedAt:               now,
						RevokedAt:               sql.NullTime{Time: now},
						ID:                      1,
						IncomingAccessRequestID: 1,
						IncomingAccessRequest: &database.IncomingAccessRequest{
							ID: 1,
							Organization: database.IncomingAccessRequestOrganization{
								SerialNumber: certBundle.Certificate().Subject.SerialNumber,
								Name:         certBundle.Certificate().Subject.Organization[0],
							},
							ServiceID: 1,
							Service: &database.Service{
								Name: "service",
							},
						},
					}, nil)

				return ctx
			},
			want: &external.AccessProof{
				Id:              1,
				CreatedAt:       ts,
				RevokedAt:       ts,
				AccessRequestId: 1,
				Organization: &external.Organization{
					SerialNumber: "00000000000000000001",
					Name:         "nlx-test",
				},
				ServiceName: "service",
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, certBundle, mocks := newService(t)
			ctx := tt.setup(t, mocks.db, certBundle)

			actual, err := service.GetAccessProof(ctx, &external.GetAccessProofRequest{
				ServiceName:          "service",
				PublicKeyFingerprint: "public-key-fingerprint",
			})

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, actual)
		})
	}
}
