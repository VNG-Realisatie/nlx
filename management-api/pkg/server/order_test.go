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

	"go.nlx.io/nlx/management-api/api"
)

func validCreateOrderRequest() api.CreateOrderRequest {
	return api.CreateOrderRequest{
		Reference:    "a-reference",
		Description:  "a-description",
		Delegatee:    "a-delegatee",
		PublicKeyPEM: testPublicKeyPEM,
		Services: []string{
			"a-service",
		},
	}
}

//nolint:funlen // this is a test method
func TestCreateOrder(t *testing.T) {
	tests := map[string]struct {
		request *api.CreateOrderRequest
		wantErr error
	}{
		"when_providing_an_empty_reference": {
			request: func() *api.CreateOrderRequest {
				request := validCreateOrderRequest()
				request.Reference = ""
				request.Services = []string{"jo"}
				return &request
			}(),
			wantErr: status.Error(codes.InvalidArgument, "invalid order: Reference: cannot be blank."),
		},
		"when_providing_a_reference_which_is_too_long": {
			request: func() *api.CreateOrderRequest {
				request := validCreateOrderRequest()
				request.Reference = "reference-with-length-of-101-chars-abcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdea"
				return &request
			}(),
			wantErr: status.Error(codes.InvalidArgument, "invalid order: Reference: the length must be between 1 and 100."),
		},
		"when_providing_an_empty_description": {
			request: func() *api.CreateOrderRequest {
				request := validCreateOrderRequest()
				request.Description = ""
				return &request
			}(),
			wantErr: status.Error(codes.InvalidArgument, "invalid order: Description: cannot be blank."),
		},
		"when_providing_a_description_which_is_too_long": {
			request: func() *api.CreateOrderRequest {
				request := validCreateOrderRequest()
				request.Description = "description-with-length-of-101-chars-abcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcd"
				return &request
			}(),
			wantErr: status.Error(codes.InvalidArgument, "invalid order: Description: the length must be between 1 and 100."),
		},
		"when_providing_an_invalid_delegatee": {
			request: func() *api.CreateOrderRequest {
				request := validCreateOrderRequest()
				request.Delegatee = "invalid / delegatee"
				return &request
			}(),
			wantErr: status.Error(codes.InvalidArgument, "invalid order: Delegatee: must be in a valid format."),
		},
		"when_providing_an_end_date_which_is_before_the_start_date": {
			request: func() *api.CreateOrderRequest {
				request := validCreateOrderRequest()
				request.ValidFrom = timestamppb.New(time.Now().Add(time.Millisecond))
				request.ValidUntil = timestamppb.New(time.Now())
				return &request
			}(),
			wantErr: status.Error(codes.InvalidArgument, "invalid order: ValidUntil: order can not expire before the start date."),
		},
		"when_providing_an_empty_list_of_services": {
			request: func() *api.CreateOrderRequest {
				request := validCreateOrderRequest()
				request.Services = []string{}
				return &request
			}(),
			wantErr: status.Error(codes.InvalidArgument, "invalid order: Services: cannot be blank."),
		},
		"when_providing_an_invalid_service_name": {
			request: func() *api.CreateOrderRequest {
				request := validCreateOrderRequest()
				request.Services = []string{
					"invalid / service name",
				}
				return &request
			}(),
			wantErr: status.Error(codes.InvalidArgument, "invalid order: Services: (0: service must be in a valid format.)."),
		},
		"when_providing_an_invalid_public_key": {
			request: func() *api.CreateOrderRequest {
				request := validCreateOrderRequest()
				request.PublicKeyPEM = "invalid"
				return &request
			}(),
			wantErr: status.Error(codes.InvalidArgument, "invalid order: PublicKeyPEM: expect public key as pem."),
		},
		"happy_path": {
			request: func() *api.CreateOrderRequest {
				request := validCreateOrderRequest()
				return &request
			}(),
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, _, _ := newService(t)

			response, err := service.CreateOrder(context.Background(), tt.request)

			if tt.wantErr == nil {
				assert.Equal(t, &emptypb.Empty{}, response)
			} else {
				assert.Error(t, err)
				assert.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, response)
			}
		})
	}
}
