package server

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/api"
)

func (s DirectoryService) GetOrganizationService(ctx context.Context, request *api.GetOrganizationServiceRequest) (*api.DirectoryService, error) {
	logger := s.logger.With(zap.String("organizationSerialNumber", request.OrganizationSerialNumber), zap.String("serviceName", request.ServiceName))
	logger.Info("rpc request GetOrganizationService")

	service, err := s.getService(ctx, logger, request.OrganizationSerialNumber, request.ServiceName)
	if err != nil {
		return nil, err
	}

	directoryService := convertDirectoryService(service)

	accessRequestStates, err := getLatestAccessRequestStates(ctx, s.configDatabase, directoryService.Organization.SerialNumber, directoryService.ServiceName)
	if err != nil {
		s.logger.Error("error getting latest access request states", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "database error")
	}

	directoryService.AccessStates = accessRequestStates

	return directoryService, nil
}
