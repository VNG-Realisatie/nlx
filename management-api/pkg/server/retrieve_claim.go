// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	common_grpcerrors "go.nlx.io/nlx/common/grpcerrors"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/grpcerrors"
)

func (s *ManagementService) RetrieveClaimForOrder(ctx context.Context, req *api.RetrieveClaimForOrderRequest) (*api.RetrieveClaimForOrderResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, grpcerrors.NewFromValidationError(err)
	}

	inwayProxyAddress, err := s.directoryClient.GetOrganizationInwayProxyAddress(ctx, req.OrderOrganizationSerialNumber)
	if err != nil {
		s.logger.Error("unable to get the inway proxy address of the external organization", zap.Error(err))
		return nil, grpcerrors.NewInternal("unable to retrieve claim", nil)
	}

	externalManagementClient, err := s.createManagementClientFunc(ctx, inwayProxyAddress, s.orgCert)
	if err != nil {
		s.logger.Error("can not setup external management client", zap.Error(err))
		return nil, grpcerrors.NewInternal("unable to retrieve claim", nil)
	}

	defer externalManagementClient.Close()

	response, err := externalManagementClient.RequestClaim(context.Background(), &external.RequestClaimRequest{
		OrderReference:                  req.OrderReference,
		ServiceOrganizationSerialNumber: req.ServiceOrganizationSerialNumber,
		ServiceName:                     req.ServiceName,
	})

	if err != nil {
		s.logger.Error("could not request claim", zap.Error(err))

		st, ok := status.FromError(err)
		if ok {
			if common_grpcerrors.Equal(err, external.ErrorReason_ORDER_NOT_FOUND) {
				return nil, grpcerrors.New(codes.NotFound, api.ErrorReason_ORDER_NOT_FOUND, "order not found", nil)
			}

			if common_grpcerrors.Equal(err, external.ErrorReason_ORDER_REVOKED) {
				return nil, grpcerrors.New(codes.Unauthenticated, api.ErrorReason_ORDER_REVOKED, "order is revoked", nil)
			}

			if st.Message() == errMessageOutwayUnableToSignClaim {
				return nil, grpcerrors.NewInternal("outway of delegator is unable to sign claim", nil)
			}
		}

		return nil, grpcerrors.NewInternal("unable to retrieve claim", nil)
	}

	return &api.RetrieveClaimForOrderResponse{
		Claim: response.Claim,
	}, nil
}
