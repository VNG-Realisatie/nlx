// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package server_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/management-api/api"
)

func TestGetTermsOfService(t *testing.T) {
	ctx := context.Background()

	tests := map[string]struct {
		setup      func(context.Context, directoryServiceMocks)
		wantResult *api.GetTermsOfServiceResponse
		wantErr    error
	}{
		"failed_to_fetch_from_directory_client": {
			setup: func(ctx context.Context, mocks directoryServiceMocks) {
				mocks.d.
					EXPECT().
					GetTermsOfService(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("arbitrary error"))
			},
			wantResult: nil,
			wantErr:    status.Error(codes.Internal, "unable to get terms of service from directory"),
		},
		"happy_flow": {
			setup: func(ctx context.Context, mocks directoryServiceMocks) {
				mocks.d.
					EXPECT().
					GetTermsOfService(gomock.Any(), gomock.Any()).
					Return(&directoryapi.GetTermsOfServiceResponse{
						Enabled: true,
						Url:     "https://example.com",
					}, nil)
			},
			wantResult: &api.GetTermsOfServiceResponse{
				Enabled: true,
				Url:     "https://example.com",
			},
			wantErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, mocks := newDirectoryService(t)

			tt.setup(ctx, mocks)

			result, err := service.GetTermsOfService(ctx, &api.GetTermsOfServiceRequest{})

			assert.Equal(t, tt.wantResult, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
