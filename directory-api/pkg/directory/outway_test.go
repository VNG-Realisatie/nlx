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
)

//nolint:funlen // This is a test function
func TestRegisterOutway(t *testing.T) {
	tests := map[string]struct {
		setup        func(serviceMocks)
		request      *directoryapi.RegisterOutwayRequest
		wantResponse *directoryapi.RegisterOutwayResponse
		wantErr      error
	}{
		"when_an_unexpected_repository_error_occurs": {
			setup: func(mocks serviceMocks) {
				mocks.repository.
					EXPECT().
					RegisterOutway(gomock.Any()).
					Return(errors.New("arbitrary error"))
			},
			request: &directoryapi.RegisterOutwayRequest{
				Name: "outway-01",
			},
			wantResponse: nil,
			wantErr:      status.New(codes.Internal, "failed to register outway").Err(),
		},
		"happy_flow_empty_name": {
			setup: func(mocks serviceMocks) {
				mocks.repository.
					EXPECT().
					RegisterOutway(gomock.Any()).
					Return(nil).
					AnyTimes()
			},
			request: &directoryapi.RegisterOutwayRequest{
				Name: "",
			},
			wantResponse: &directoryapi.RegisterOutwayResponse{},
			wantErr:      nil,
		},
		"happy_flow": {
			setup: func(mocks serviceMocks) {
				mocks.repository.
					EXPECT().
					RegisterOutway(gomock.Any()).
					Return(nil).
					AnyTimes()
			},
			request: &directoryapi.RegisterOutwayRequest{
				Name: "outway-01",
			},
			wantResponse: &directoryapi.RegisterOutwayResponse{},
			wantErr:      nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, mocks := newService(t, testNlxVersion128, "", &testClock{
				timeToReturn: time.Now(),
			})

			if tt.setup != nil {
				tt.setup(mocks)
			}

			got, err := service.RegisterOutway(context.Background(), tt.request)

			assert.Equal(t, tt.wantResponse, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
