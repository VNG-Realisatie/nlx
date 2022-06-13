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
)

func (s DirectoryService) GetTermsOfService(ctx context.Context, _ *emptypb.Empty) (*api.GetTermsOfServiceResponse, error) {
	s.logger.Info("rpc request GetTermsOfService")

	response, err := s.directoryClient.GetTermsOfService(ctx, &emptypb.Empty{})
	if err != nil {
		s.logger.Debug("unable to get terms of service from directory", zap.Error(err))
		return nil, status.Error(codes.Internal, "unable to get terms of service from directory")
	}

	return &api.GetTermsOfServiceResponse{
		Enabled: response.Enabled,
		Url:     response.Url,
	}, nil
}
