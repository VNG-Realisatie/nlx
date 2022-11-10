// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package server

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/permissions"
	"go.nlx.io/nlx/management-api/pkg/syncer"
)

type SyncError string

const (
	InternalError                               SyncError = "internal_error"
	ServiceProviderNoOrganizationInwaySpecified SyncError = "service_provider_no_organization_inway_specified"
	ServiceProviderOrganizationInwayUnreachable SyncError = "service_provider_organization_inway_unreachable"
)

func (s *ManagementService) SynchronizeOutgoingAccessRequests(ctx context.Context, req *api.SynchronizeOutgoingAccessRequestsRequest) (*api.SynchronizeOutgoingAccessRequestsResponse, error) {
	err := s.authorize(ctx, permissions.SyncOutgoingAccessRequests)
	if err != nil {
		return nil, err
	}

	logger := s.logger.With(zap.String("organizationSerialNumber", req.OrganizationSerialNumber), zap.String("serviceName", req.ServiceName))
	logger.Info("rpc request SynchronizeLatestOutgoingAccessRequest")

	outgoingAccessRequests, err := s.configDatabase.ListLatestOutgoingAccessRequests(ctx, req.OrganizationSerialNumber, req.ServiceName)
	if err != nil {
		s.logger.Error("error getting latest access request states", zap.Error(err))
		return nil, status.Errorf(codes.Internal, string(InternalError))
	}

	if len(outgoingAccessRequests) < 1 {
		return &api.SynchronizeOutgoingAccessRequestsResponse{}, nil
	}

	organizationInwayProxyAddress, err := s.directoryClient.GetOrganizationInwayProxyAddress(ctx, req.OrganizationSerialNumber)
	if err != nil {
		s.logger.Error("cannot get organization inway proxy address", zap.Error(err))
		return nil, status.Errorf(codes.Internal, string(InternalError))
	}

	if organizationInwayProxyAddress == "" {
		return nil, status.Errorf(codes.Internal, string(ServiceProviderNoOrganizationInwaySpecified))
	}

	client, err := s.createManagementClientFunc(ctx, organizationInwayProxyAddress, s.orgCert)
	if err != nil {
		s.logger.Error("cannot setup management client", zap.Error(err))

		return nil, status.Errorf(codes.Internal, string(ServiceProviderOrganizationInwayUnreachable))
	}

	defer client.Close()

	err = syncer.SyncOutgoingAccessRequests(&syncer.SyncArgs{
		Ctx:      ctx,
		Logger:   s.logger,
		DB:       s.configDatabase,
		Client:   client,
		Requests: outgoingAccessRequests,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, string(InternalError))
	}

	return &api.SynchronizeOutgoingAccessRequestsResponse{}, nil
}
