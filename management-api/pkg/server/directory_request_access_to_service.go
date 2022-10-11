// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package server

import (
	"context"

	"go.uber.org/zap"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
)

func (s DirectoryService) RequestAccessToService(ctx context.Context, request *api.RequestAccessToServiceRequest) (*api.RequestAccessToServiceResponse, error) {
	logger := s.logger.With(zap.String("organizationSerialNumber", request.OrganizationSerialNumber), zap.String("serviceName", request.ServiceName))
	logger.Info("rpc request RequestAccessToService")

	ar := &database.OutgoingAccessRequest{
		Organization: database.Organization{
			SerialNumber: request.OrganizationSerialNumber,
		},
		ServiceName: request.ServiceName,
	}

	accessRequest, err := s.configDatabase.CreateOutgoingAccessRequest(ctx, ar)
	if err != nil {
		return nil, err
	}

	response := convertOutgoingAccessRequest(accessRequest)

	service, err := s.getService(ctx, logger, request.OrganizationSerialNumber, request.ServiceName)
	if err != nil {
		return nil, err
	}

	logger = logger.With(zap.Any("service", service))

	logger.Debug("send access request to inway")

	return &api.RequestAccessToServiceResponse{
		OutgoingAccessRequest: response,
	}, nil
}
