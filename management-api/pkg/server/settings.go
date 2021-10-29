// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
)

// GetSettings returns the settings for the organization
func (s *ManagementService) GetSettings(ctx context.Context, _ *emptypb.Empty) (*api.Settings, error) {
	logger := s.logger.With(zap.String("handler", "get-settings"))

	settings, err := s.configDatabase.GetSettings(ctx)
	if err != nil {
		if errIsNotFound(err) {
			return &api.Settings{
				OrganizationInway: "",
			}, nil
		}

		logger.Error("could not get the settings from the database", zap.Error(err))

		return nil, status.Error(codes.Internal, "database error")
	}

	result := convertFromDatabaseSettings(settings)

	return result, nil
}

// UpdateSettings updates the settings for the organization
func (s *ManagementService) UpdateSettings(ctx context.Context, req *api.UpdateSettingsRequest) (*emptypb.Empty, error) {
	logger := s.logger.With(zap.String("handler", "update-settings"))

	var inwayID *uint

	if req.OrganizationInway != "" {
		inway, err := s.configDatabase.GetInway(ctx, req.OrganizationInway)
		if err != nil {
			if errIsNotFound(err) {
				logger.Warn("inway not found", zap.String("inway", req.OrganizationInway))
				return nil, status.Error(codes.InvalidArgument, "inway not found")
			}

			logger.Error("could not get the inway from the database", zap.Error(err))

			return nil, status.Error(codes.Internal, "database error")
		}

		inwayID = &inway.ID
	} else {
		_, err := s.directoryClient.ClearOrganizationInway(ctx, &emptypb.Empty{})
		if err != nil {
			logger.Error("could not clear the organization inway in the directory", zap.Error(err))
			return nil, status.Error(codes.Internal, "database error")
		}
	}

	userInfo, err := retrieveUserInfoFromGRPCContext(ctx)
	if err != nil {
		logger.Error("could not retrieve user info for audit log from grpc context", zap.Error(err))
		return nil, status.Error(codes.Internal, "could not retrieve user info to create audit log")
	}

	err = s.auditLogger.OrganizationSettingsUpdate(ctx, userInfo.username, userInfo.userAgent)
	if err != nil {
		return nil, status.Error(codes.Internal, "could not create audit log")
	}

	_, err = s.configDatabase.PutOrganizationInway(ctx, inwayID)
	if err != nil {
		logger.Error("could not update the settings in the database", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	return &emptypb.Empty{}, nil
}

func convertFromDatabaseSettings(model *database.Settings) *api.Settings {
	if model.Inway == nil {
		return &api.Settings{}
	}

	return &api.Settings{
		OrganizationInway: model.Inway.Name,
	}
}
