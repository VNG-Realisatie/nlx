// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package server_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/domain"
	"go.nlx.io/nlx/management-api/pkg/database"
)

//nolint:funlen // this is a test function
func TestAcceptTermsOfServiceStatus(t *testing.T) {
	tests := map[string]struct {
		setup   func(context.Context, serviceMocks)
		ctx     context.Context
		want    *emptypb.Empty
		wantErr error
	}{
		"when_accepting_terms_of_service_fails": {
			setup: func(ctx context.Context, mocks serviceMocks) {
				mocks.db.
					EXPECT().
					AcceptTermsOfService(ctx, "Jane Doe", gomock.Any()).
					Return(false, errors.New("arbitrary error"))
			},
			ctx: metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
				"username":               "Jane Doe",
				"grpcgateway-user-agent": "nlxctl",
			})),
			want:    nil,
			wantErr: status.Errorf(codes.Internal, "database error"),
		},
		"when_writing_audit_logs_fails": {
			setup: func(ctx context.Context, mocks serviceMocks) {
				mocks.db.
					EXPECT().
					AcceptTermsOfService(ctx, "Jane Doe", gomock.Any()).
					Return(false, nil)

				mocks.al.
					EXPECT().
					AcceptTermsOfService(
						gomock.Any(),
						"Jane Doe",
						"nlxctl",
					).
					Return(errors.New("arbitrary error"))
			},
			ctx: metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
				"username":               "Jane Doe",
				"grpcgateway-user-agent": "nlxctl",
			})),
			want:    nil,
			wantErr: status.Error(codes.Internal, "could not create audit log"),
		},
		"happy_flow": {
			setup: func(ctx context.Context, mocks serviceMocks) {
				mocks.al.
					EXPECT().
					AcceptTermsOfService(
						gomock.Any(),
						"Jane Doe",
						"nlxctl",
					)

				mocks.db.
					EXPECT().
					AcceptTermsOfService(ctx, "Jane Doe", gomock.Any()).
					Return(false, nil)
			},
			ctx: metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
				"username":               "Jane Doe",
				"grpcgateway-user-agent": "nlxctl",
			})),
			want:    &emptypb.Empty{},
			wantErr: nil,
		},
		"happy_flow_already_accepted": {
			setup: func(ctx context.Context, mocks serviceMocks) {
				mocks.db.
					EXPECT().
					AcceptTermsOfService(ctx, "Jane Doe", gomock.Any()).
					Return(true, nil)
			},
			ctx: metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
				"username":               "Jane Doe",
				"grpcgateway-user-agent": "nlxctl",
			})),
			want:    &emptypb.Empty{},
			wantErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, _, mocks := newService(t)
			tt.setup(tt.ctx, mocks)

			got, err := service.AcceptTermsOfService(tt.ctx, &emptypb.Empty{})

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestGetTermsOfServiceStatus(t *testing.T) {
	tests := map[string]struct {
		setup   func(context.Context, serviceMocks)
		want    *api.GetTermsOfServiceStatusResponse
		wantErr error
	}{
		"when_getting_terms_of_service_status_fails": {
			setup: func(ctx context.Context, mocks serviceMocks) {
				mocks.db.
					EXPECT().
					GetTermsOfServiceStatus(ctx).
					Return(nil, errors.New("arbitrary error"))
			},
			want:    nil,
			wantErr: status.Errorf(codes.Internal, "database error"),
		},
		"happy_flow_not_accepted": {
			setup: func(ctx context.Context, mocks serviceMocks) {
				mocks.db.
					EXPECT().
					GetTermsOfServiceStatus(ctx).
					Return(nil, database.ErrNotFound)
			},
			want: &api.GetTermsOfServiceStatusResponse{
				Accepted: false,
			},
			wantErr: nil,
		},
		"happy_flow_accepted": {
			setup: func(ctx context.Context, mocks serviceMocks) {
				model, err := domain.NewTermsOfServiceStatus(&domain.NewTermsOfServiceStatusArgs{
					Username:  "Jane Doe",
					CreatedAt: time.Now(),
				})
				require.NoError(t, err)

				mocks.db.
					EXPECT().
					GetTermsOfServiceStatus(ctx).
					Return(model, nil)
			},
			want: &api.GetTermsOfServiceStatusResponse{
				Accepted: true,
			},
			wantErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, _, mocks := newService(t)
			tt.setup(context.Background(), mocks)

			got, err := service.GetTermsOfServiceStatus(context.Background(), &emptypb.Empty{})

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
