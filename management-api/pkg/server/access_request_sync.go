// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package server

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/permissions"
	"go.nlx.io/nlx/management-api/pkg/syncer"
)

func (s *ManagementService) SynchronizeOutgoingAccessRequests(ctx context.Context, req *api.SynchronizeOutgoingAccessRequestsRequest) (*emptypb.Empty, error) {
	err := s.authorize(ctx, permissions.SyncOutgoingAccessRequests)
	if err != nil {
		return nil, err
	}

	logger := s.logger.With(zap.String("organizationSerialNumber", req.OrganizationSerialNumber), zap.String("serviceName", req.ServiceName))
	logger.Info("rpc request SynchronizeLatestOutgoingAccessRequest")

	outgoingAccessRequests, err := s.configDatabase.ListLatestOutgoingAccessRequests(ctx, req.OrganizationSerialNumber, req.ServiceName)
	if err != nil {
		s.logger.Error("error getting latest access request states", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	if len(outgoingAccessRequests) < 1 {
		return &emptypb.Empty{}, nil
	}

	organizationInwayProxyAddress, err := s.directoryClient.GetOrganizationInwayProxyAddress(ctx, req.OrganizationSerialNumber)
	if err != nil {
		s.logger.Error("cannot get organization inway proxy address", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	client, err := s.createManagementClientFunc(ctx, organizationInwayProxyAddress, s.orgCert)
	if err != nil {
		s.logger.Error("cannot setup management client", zap.Error(err))

		return nil, status.Errorf(codes.Internal, "internal error")
	}

	defer client.Close()

	err = syncer.SyncOutgoingAccessRequests(&syncer.SyncArgs{
		Ctx:      ctx,
		Logger:   s.logger,
		Clock:    s.clock,
		DB:       s.configDatabase,
		Client:   client,
		Requests: outgoingAccessRequests,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error occurred while syncing at least one Outgoing Access Request")
	}

	return &emptypb.Empty{}, nil
}
