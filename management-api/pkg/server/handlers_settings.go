// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
)

// GetSettings returns the settings for the organization
func (s *ManagementService) GetSettings(ctx context.Context, req *api.GetSettingsRequest) (*api.Settings, error) {
	logger := s.logger.With(zap.String("handler", "get-settings"))

	settings, err := s.configDatabase.GetSettings(ctx)
	if err != nil {
		logger.Error("could not get the settings from the database", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	result := convertFromDatabaseSettings(settings)

	return result, nil
}

// UpdateSettings updates the settings for the organization
func (s *ManagementService) UpdateSettings(ctx context.Context, req *api.UpdateSettingsRequest) (*api.Empty, error) {
	logger := s.logger.With(zap.String("handler", "update-settings"))

	settings := database.Settings{
		InwayNameForManagementAPITraffic: req.InwayNameForManagementApiTraffic,
	}

	err := s.configDatabase.UpdateSettings(ctx, &settings)
	if err != nil {
		logger.Error("could not update the settings in the database", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	return &api.Empty{}, nil
}

func convertFromDatabaseSettings(model *database.Settings) *api.Settings {
	settings := &api.Settings{
		InwayNameForManagementApiTraffic: model.InwayNameForManagementAPITraffic,
	}

	return settings
}
