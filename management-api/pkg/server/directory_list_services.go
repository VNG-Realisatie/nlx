// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/management-api/api"
)

func (s DirectoryService) ListServices(ctx context.Context, _ *emptypb.Empty) (*api.DirectoryListServicesResponse, error) {
	s.logger.Info("rpc request ListServices")

	resp, err := s.directoryClient.ListServices(ctx, &directoryapi.ListServicesRequest{})
	if err != nil {
		s.logger.Error("error getting services list from directory", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "directory error")
	}

	services := make([]*api.DirectoryService, len(resp.Services))

	for i, service := range resp.Services {
		convertedService := convertDirectoryService(service)

		accessRequestStates, err := getLatestAccessRequestStates(ctx, s.directoryClient, s.configDatabase, convertedService.Organization.SerialNumber, convertedService.ServiceName)
		if err != nil {
			s.logger.Error("error getting latest access request states", zap.Error(err))
			return nil, status.Errorf(codes.Internal, "database error")
		}

		convertedService.AccessStates = accessRequestStates

		services[i] = convertedService
	}

	return &api.DirectoryListServicesResponse{Services: services}, nil
}
