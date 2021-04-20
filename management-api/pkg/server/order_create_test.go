// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/auditlog"
	"go.nlx.io/nlx/management-api/pkg/database"
)

//nolint:funlen // this is a test method
func TestCreateOrder(t *testing.T) {
	validFrom := time.Now().UTC()
	validUntil := time.Now().Add(1 * time.Hour).UTC()
	validCreateOrderRequest := func() api.CreateOrderRequest {
		return api.CreateOrderRequest{
			Reference:    "a-reference",
			Description:  "a-description",
			Delegatee:    "a-delegatee",
			PublicKeyPEM: testPublicKeyPEM,
			ValidFrom:    timestamppb.New(validFrom),
			ValidUntil:   timestamppb.New(validUntil),
			Services: []*api.CreateOrderRequest_Service{
				{
					Organization: "a-organization",
					Service:      "a-service",
				},
			},
		}
	}

	tests := map[string]struct {
		request *api.CreateOrderRequest
		setup   func(serviceMocks)
		wantErr error
	}{
		"when_providing_an_empty_reference": {
			request: func() *api.CreateOrderRequest {
				request := validCreateOrderRequest()
				request.Reference = ""
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
				request.Services = []*api.CreateOrderRequest_Service{}
				return &request
			}(),
			wantErr: status.Error(codes.InvalidArgument, "invalid order: Services: cannot be blank."),
		},
		"when_providing_an_invalid_service_name": {
			request: func() *api.CreateOrderRequest {
				request := validCreateOrderRequest()
				request.Services = []*api.CreateOrderRequest_Service{
					{
						Service: "invalid / service name",
					},
				}
				return &request
			}(),
			wantErr: status.Error(codes.InvalidArgument, "invalid order: Services: (0: (Service: service must be in a valid format.).)."),
		},
		"when_providing_an_invalid_public_key": {
			request: func() *api.CreateOrderRequest {
				request := validCreateOrderRequest()
				request.PublicKeyPEM = "invalid"
				return &request
			}(),
			wantErr: status.Error(codes.InvalidArgument, "invalid order: PublicKeyPEM: expect public key as pem."),
		},
		"when_creating_the_order_fails": {
			wantErr: status.Error(codes.Internal, "failed to create order"),
			setup: func(mocks serviceMocks) {
				mocks.al.
					EXPECT().
					OrderCreate(gomock.Any(), "Jane Doe", "nlxctl", "a-delegatee", []auditlog.RecordService{
						{
							Organization: "a-organization",
							Service:      "a-service",
						},
					})

				mocks.db.
					EXPECT().
					CreateOrder(gomock.Any(), &database.Order{
						Reference:    "a-reference",
						Description:  "a-description",
						PublicKeyPEM: testPublicKeyPEM,
						Delegatee:    "a-delegatee",
						ValidFrom:    validFrom,
						ValidUntil:   validUntil,
						Services: []database.OrderService{
							{
								Organization: "a-organization",
								Service:      "a-service",
							},
						},
					}).
					Return(errors.New("arbitrary error"))
			},
			request: func() *api.CreateOrderRequest {
				request := validCreateOrderRequest()
				return &request
			}(),
		},
		"when_creating_audit_log_Fails": {
			wantErr: status.Error(codes.Internal, "failed to write to auditlog"),
			setup: func(mocks serviceMocks) {
				mocks.al.
					EXPECT().
					OrderCreate(gomock.Any(), "Jane Doe", "nlxctl", "a-delegatee", []auditlog.RecordService{
						{
							Organization: "a-organization",
							Service:      "a-service",
						},
					}).
					Return(errors.New("arbitrary error"))
			},
			request: func() *api.CreateOrderRequest {
				request := validCreateOrderRequest()
				return &request
			}(),
		},
		"happy_path": {
			setup: func(mocks serviceMocks) {
				mocks.db.
					EXPECT().
					CreateOrder(gomock.Any(), &database.Order{
						Reference:    "a-reference",
						Description:  "a-description",
						PublicKeyPEM: testPublicKeyPEM,
						Delegatee:    "a-delegatee",
						ValidFrom:    validFrom,
						ValidUntil:   validUntil,
						Services: []database.OrderService{
							{
								Organization: "a-organization",
								Service:      "a-service",
							},
						},
					}).
					Return(nil)

				mocks.al.
					EXPECT().
					OrderCreate(gomock.Any(), "Jane Doe", "nlxctl", "a-delegatee", []auditlog.RecordService{
						{
							Organization: "a-organization",
							Service:      "a-service",
						},
					})
			},
			request: func() *api.CreateOrderRequest {
				request := validCreateOrderRequest()
				return &request
			}(),
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, _, mocks := newService(t)

			if tt.setup != nil {
				tt.setup(mocks)
			}

			ctx := metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
				"username":               "Jane Doe",
				"grpcgateway-user-agent": "nlxctl",
			}))

			response, err := service.CreateOrder(ctx, tt.request)

			if tt.wantErr == nil {
				assert.IsType(t, &emptypb.Empty{}, response)
			} else {
				assert.Error(t, err)
				assert.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, response)
			}
		})
	}
}
