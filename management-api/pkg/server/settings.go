// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server

import (
	"context"
	"errors"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/domain"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/permissions"
)

// GetSettings returns the settings for the organization
func (s *ManagementService) GetSettings(ctx context.Context, _ *emptypb.Empty) (*api.Settings, error) {
	err := s.authorize(ctx, permissions.ReadOrganizationSettings)
	if err != nil {
		return nil, err
	}

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
	err := s.authorize(ctx, permissions.UpdateOrganizationSettings)
	if err != nil {
		return nil, err
	}

	logger := s.logger.With(zap.String("handler", "update-settings"))

	userInfo, err := retrieveUserFromContext(ctx)
	if err != nil {
		logger.Error("could not retrieve user info for audit log from grpc context", zap.Error(err))
		return nil, status.Error(codes.Internal, "could not retrieve user info to create audit log")
	}

	settings, err := domain.NewSettings(req.OrganizationInway, req.OrganizationEmailAddress)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = s.auditLogger.OrganizationSettingsUpdate(ctx, userInfo.Email, userInfo.UserAgent)
	if err != nil {
		logger.Error("could not create audit log", zap.Error(err))
		return nil, status.Error(codes.Internal, "could not create audit log")
	}

	if settings.OrganizationInwayName() == "" {
		_, err = s.directoryClient.ClearOrganizationInway(ctx, &emptypb.Empty{})
		if err != nil {
			logger.Error("could not clear the organization inway in the directory", zap.Error(err))
			return nil, status.Error(codes.Internal, "nlx directory unreachable")
		}
	}

	_, err = s.directoryClient.SetOrganizationContactDetails(ctx, &directoryapi.SetOrganizationContactDetailsRequest{
		EmailAddress: settings.OrganizationEmailAddress(),
	})
	if err != nil {
		logger.Error("could not update the organization email address in the directory", zap.Error(err))
		return nil, status.Error(codes.Internal, "nlx directory unreachable")
	}

	err = s.configDatabase.UpdateSettings(ctx, settings)
	if err != nil {
		if errors.Is(err, database.ErrInwayNotFound) {
			logger.Warn("inway not found", zap.String("inway", req.OrganizationInway))
			return nil, status.Error(codes.InvalidArgument, "inway not found")
		}

		logger.Error("could not update the settings in the database", zap.Error(err))

		return nil, status.Error(codes.Internal, "database error")
	}

	return &emptypb.Empty{}, nil
}

func convertFromDatabaseSettings(model *domain.Settings) *api.Settings {
	return &api.Settings{
		OrganizationInway:        model.OrganizationInwayName(),
		OrganizationEmailAddress: model.OrganizationEmailAddress(),
	}
}
