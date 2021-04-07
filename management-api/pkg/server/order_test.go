// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/management-api/api"
)

func validCreateOrderRequest() api.CreateOrderRequest {
	return api.CreateOrderRequest{
		Reference:   "a-reference",
		Description: "a-description",
		Delegatee:   "a-delegatee",
		Services: []string{
			"a-service",
		},
	}
}

func TestCreateOrder(t *testing.T) {
	tests := map[string]struct {
		request *api.CreateOrderRequest
		want    func(*testing.T, *common_tls.CertificateBundle, *emptypb.Empty)
		wantErr error
	}{
		"when_providing_an_empty_reference": {
			request: func() *api.CreateOrderRequest {
				request := validCreateOrderRequest()
				request.Reference = ""
				return &request
			}(),
			wantErr: status.Error(codes.InvalidArgument, "invalid order: the reference must be provided"),
		},
		"when_providing_a_reference_which_is_too_long": {
			request: func() *api.CreateOrderRequest {
				request := validCreateOrderRequest()
				request.Reference = "reference-with-length-of-101-chars-abcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdea"
				return &request
			}(),
			wantErr: status.Error(codes.InvalidArgument, "invalid order: the reference should not exceed 100 characters"),
		},
		"when_providing_an_empty_description": {
			request: func() *api.CreateOrderRequest {
				request := validCreateOrderRequest()
				request.Description = ""
				return &request
			}(),
			wantErr: status.Error(codes.InvalidArgument, "invalid order: the description must be provided"),
		},
		"when_providing_a_description_which_is_too_long": {
			request: func() *api.CreateOrderRequest {
				request := validCreateOrderRequest()
				request.Description = "description-with-length-of-101-chars-abcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcd"
				return &request
			}(),
			wantErr: status.Error(codes.InvalidArgument, "invalid order: the description should not exceed 100 characters"),
		},
		"when_providing_an_invalid_delegatee": {
			request: func() *api.CreateOrderRequest {
				request := validCreateOrderRequest()
				request.Delegatee = "invalid / delegatee"
				return &request
			}(),
			wantErr: status.Error(codes.InvalidArgument, "invalid order: the delegatee should be a valid organization name (alphanumeric, max. 100 chars)"),
		},
		"when_providing_an_end_date_which_is_before_the_start_date": {
			request: func() *api.CreateOrderRequest {
				request := validCreateOrderRequest()
				request.ValidFrom = timestamppb.New(time.Now().Add(time.Millisecond))
				request.ValidUntil = timestamppb.New(time.Now())
				return &request
			}(),
			wantErr: status.Error(codes.InvalidArgument, "invalid order: valid from should be a timestamp before the valid until timestamp"),
		},
		"when_providing_an_empty_list_of_services": {
			request: func() *api.CreateOrderRequest {
				request := validCreateOrderRequest()
				request.Services = []string{}
				return &request
			}(),
			wantErr: status.Error(codes.InvalidArgument, "invalid order: at least one service should be specified"),
		},
		"when_providing_an_invalid_service_name": {
			request: func() *api.CreateOrderRequest {
				request := validCreateOrderRequest()
				request.Services = []string{
					"invalid / service name",
				}
				return &request
			}(),
			wantErr: status.Error(codes.InvalidArgument, "invalid order: service 'invalid / service name' is not a valid service name"),
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, _, _ := newService(t)

			_, err := service.CreateOrder(context.Background(), tt.request)
			assert.Error(t, err)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}
