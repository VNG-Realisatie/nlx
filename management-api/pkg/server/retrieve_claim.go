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

func (s *ManagementService) RetrieveClaim(ctx context.Context, req *api.RetrieveClaimForOrderRequest) (*api.RetrieveClaimForOrderResponse, error) {
	_, err := s.parseProxyMetadata(ctx)
	if err != nil {
		s.logger.Error("failed to parse proxy metadata", zap.Error(err))
		return nil, err
	}

	if len(req.OrderReference) < 1 {
		return nil, status.Error(codes.InvalidArgument, "an order reference must be provided")
	}

	if len(req.OrderOrganizationName) < 1 {
		return nil, status.Error(codes.InvalidArgument, "an organization name of the order must be provided")
	}

	inwayProxyAddress, err := s.directoryClient.GetOrganizationInwayProxyAddress(ctx, req.OrderOrganizationName)
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
		OrderReference: req.OrderReference,
	})

	if err != nil {
		s.logger.Error("could not request claim", zap.Error(err))
		return nil, status.Error(codes.Internal, "unable to retrieve claim")
	}

	return &api.RetrieveClaimForOrderResponse{
		Claim: response.Claim,
	}, nil
}
