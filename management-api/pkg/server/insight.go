package server

import (
	"context"

	"github.com/gogo/protobuf/types"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/directory-registration-api/registrationapi"
	"go.nlx.io/nlx/management-api/api"
)

// GetInsight returns the insight configuration of a organization
func (s *ManagementService) GetInsightConfiguration(ctx context.Context, _ *types.Empty) (*api.InsightConfiguration, error) {
	s.logger.Info("rpc request GetInsightConfiguration")

	settings, err := s.configDatabase.GetSettings(ctx)
	if err != nil {
		if errIsNotFound(err) {
			s.logger.Warn("insight configuration not found")

			return nil, status.Error(codes.NotFound, "insight configuration not found")
		}

		s.logger.Error("error getting insight from DB", zap.Error(err))

		return nil, status.Error(codes.Internal, "database error")
	}

	return &api.InsightConfiguration{
		IrmaServerURL: settings.IrmaServerURL,
		InsightAPIURL: settings.InsightAPIURL,
	}, nil
}

// PutInsight sets the insight configuration of a organization
func (s *ManagementService) PutInsightConfiguration(ctx context.Context, req *api.InsightConfiguration) (*api.InsightConfiguration, error) {
	s.logger.Info("rpc request PutInsight")

	_, err := s.directoryClient.SetInsightConfiguration(ctx, &registrationapi.SetInsightConfigurationRequest{
		InsightAPIURL: req.InsightAPIURL,
		IrmaServerURL: req.IrmaServerURL,
	})
	if err != nil {
		return nil, err
	}

	_, err = s.configDatabase.PutInsightConfiguration(ctx, req.IrmaServerURL, req.InsightAPIURL)
	if err != nil {
		s.logger.Error("error updating inway insight config in DB", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	return req, nil
}
