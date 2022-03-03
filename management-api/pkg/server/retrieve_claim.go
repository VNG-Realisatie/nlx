// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
)

func (s *ManagementService) RetrieveClaimForOrder(ctx context.Context, req *api.RetrieveClaimForOrderRequest) (*api.RetrieveClaimForOrderResponse, error) {
	if len(req.OrderReference) < 1 {
		return nil, status.Error(codes.InvalidArgument, "an order reference must be provided")
	}

	if len(req.OrderOrganizationSerialNumber) < 1 {
		return nil, status.Error(codes.InvalidArgument, "an organization serial number of the order must be provided")
	}

	if req.ServiceName == "" {
		return nil, status.Error(codes.InvalidArgument, "a service name must be provided")
	}

	if req.ServiceOrganizationSerialNumber == "" {
		return nil, status.Error(codes.InvalidArgument, "an organization serial number for the service must be provided")
	}

	inwayProxyAddress, err := s.directoryClient.GetOrganizationInwayProxyAddress(ctx, req.OrderOrganizationSerialNumber)
	if err != nil {
		s.logger.Error("unable to get the inway proxy address of the external organization", zap.Error(err))
		return nil, status.Error(codes.Internal, "unable to retrieve claim")
	}

	externalManagementClient, err := s.createManagementClientFunc(ctx, inwayProxyAddress, s.orgCert)
	if err != nil {
		s.logger.Error("can not setup external management client", zap.Error(err))
		return nil, status.Error(codes.Internal, "unable to retrieve claim")
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
			if st.Message() == errMessageOrderRevoked {
				return nil, status.Errorf(codes.Unauthenticated, errMessageOrderRevoked)
			}
		}

		return nil, status.Error(codes.Internal, "unable to retrieve claim")
	}

	return &api.RetrieveClaimForOrderResponse{
		Claim: response.Claim,
	}, nil
}
