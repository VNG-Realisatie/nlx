// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server

import (
	"context"

	"github.com/gogo/protobuf/types"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/directory-registration-api/registrationapi"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
)

// GetSettings returns the settings for the organization
func (s *ManagementService) GetSettings(ctx context.Context, _ *types.Empty) (*api.Settings, error) {
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
func (s *ManagementService) UpdateSettings(ctx context.Context, req *api.UpdateSettingsRequest) (*types.Empty, error) {
	logger := s.logger.With(zap.String("handler", "update-settings"))

	if req.OrganizationInway != "" {
		inway, err := s.configDatabase.GetInway(ctx, req.OrganizationInway)
		if err != nil {
			logger.Error("could not get the inway from the database", zap.Error(err))
			return nil, status.Error(codes.Internal, "database error")
		}

		setOrganizationInwayRequest := &registrationapi.SetOrganizationInwayRequest{
			Address: inway.SelfAddress,
		}

		_, err = s.directoryClient.SetOrganizationInway(ctx, setOrganizationInwayRequest)
		if err != nil {
			logger.Error("could not update the settings in the directory", zap.Error(err))
			return nil, status.Error(codes.Internal, "database error")
		}
	} else {
		_, err := s.directoryClient.ClearOrganizationInway(ctx, &types.Empty{})
		if err != nil {
			logger.Error("could not clear organization inway in the directory", zap.Error(err))
			return nil, status.Error(codes.Internal, "database error")
		}
	}

	settings := database.Settings{
		OrganizationInway: req.OrganizationInway,
	}

	err := s.configDatabase.UpdateSettings(ctx, &settings)
	if err != nil {
		logger.Error("could not update the settings in the database", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	return &types.Empty{}, nil
}

func convertFromDatabaseSettings(model *database.Settings) *api.Settings {
	settings := &api.Settings{
		OrganizationInway: model.OrganizationInway,
	}

	return settings
}
