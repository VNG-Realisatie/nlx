package server

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/directory-registration-api/registrationapi"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
)

// GetInsight returns the insight configuration of a organization
func (s *ManagementService) GetInsightConfiguration(ctx context.Context, req *api.Empty) (*api.InsightConfiguration, error) {
	s.logger.Info("rpc request GetInsightConfiguration")

	insightConfig, err := s.configDatabase.GetInsightConfiguration(ctx)
	if err != nil {
		s.logger.Error("error getting insight from DB", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	if insightConfig == nil {
		s.logger.Warn("insight configuration not found")
		return nil, status.Error(codes.NotFound, "insight configuration not found")
	}

	response := &api.InsightConfiguration{
		IrmaServerURL: insightConfig.IrmaServerURL,
		InsightAPIURL: insightConfig.InsightAPIURL,
	}

	return response, nil
}

// PutInsight sets the insight configuration of a organization
func (s *ManagementService) PutInsightConfiguration(ctx context.Context, req *api.InsightConfiguration) (*api.InsightConfiguration, error) {
	s.logger.Info("rpc request PutInsight")

	_, err := s.directoryRegistrationClient.SetInsightConfiguration(ctx, &registrationapi.SetInsightConfigurationRequest{
		InsightAPIURL: req.InsightAPIURL,
		IrmaServerURL: req.IrmaServerURL,
	})

	if err != nil {
		return nil, err
	}

	config := &database.InsightConfiguration{
		IrmaServerURL: req.IrmaServerURL,
		InsightAPIURL: req.InsightAPIURL,
	}

	err = s.configDatabase.PutInsightConfiguration(ctx, config)
	if err != nil {
		s.logger.Error("error updating inway insight config in DB", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	return req, nil
}
