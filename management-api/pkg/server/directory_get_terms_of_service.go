// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package server

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/management-api/api"
)

func (s DirectoryService) GetTermsOfService(ctx context.Context, _ *api.GetTermsOfServiceRequest) (*api.GetTermsOfServiceResponse, error) {
	s.logger.Info("rpc request GetTermsOfService")

	response, err := s.directoryClient.GetTermsOfService(ctx, &directoryapi.GetTermsOfServiceRequest{})
	if err != nil {
		s.logger.Debug("unable to get terms of service from directory", zap.Error(err))
		return nil, status.Error(codes.Internal, "unable to get terms of service from directory")
	}

	return &api.GetTermsOfServiceResponse{
		Enabled: response.Enabled,
		Url:     response.Url,
	}, nil
}
