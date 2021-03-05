// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package registrationservice

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"go.nlx.io/nlx/directory-registration-api/pkg/database"
	"go.nlx.io/nlx/directory-registration-api/registrationapi"
)

// nolint:gocritic // these are valid regex patterns
var organizationNameRegex = regexp.MustCompile(`^[a-zA-Z0-9-. _\s]{1,100}$`)

func IsValidOrganizationName(name string) bool {
	return organizationNameRegex.MatchString(name)
}

func (h *DirectoryRegistrationService) SetInsightConfiguration(ctx context.Context, req *registrationapi.SetInsightConfigurationRequest) (*emptypb.Empty, error) {
	logger := h.logger.With(zap.String("handler", "set-insight-configuration"))

	logger.Info("rpc request SetInsightConfiguration", zap.String("insight api url", req.InsightAPIURL), zap.String("irma server url", req.IrmaServerURL))

	organizationName, err := h.getOrganisationNameFromRequest(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get organization name from request: %v", err)
	}

	if !IsValidOrganizationName(organizationName) {
		logger.Error("invalid organization name", zap.String("organization name", organizationName))
		return nil, status.New(codes.InvalidArgument, "Invalid organization name").Err()
	}

	err = h.db.SetInsightConfiguration(ctx, organizationName, req.InsightAPIURL, req.IrmaServerURL)
	if err != nil {
		logger.Error("failed to execute SetInsightConfiguration", zap.Error(err))
		return nil, status.New(codes.Internal, "database error").Err()
	}

	return &emptypb.Empty{}, nil
}

func (h *DirectoryRegistrationService) SetOrganizationInway(ctx context.Context, req *registrationapi.SetOrganizationInwayRequest) (*emptypb.Empty, error) {
	logger := h.logger.With(zap.String("handler", "SetOrganizationInway"))

	organizationName, err := h.getOrganisationNameFromRequest(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "determine organization name")
	}

	if req.Address == "" {
		return nil, status.New(codes.InvalidArgument, "address is empty").Err()
	}

	err = h.db.SetOrganizationInway(ctx, organizationName, req.Address)
	if err != nil {
		if errors.Is(err, database.ErrNoInwayWithAddress) {
			return nil, status.New(codes.NotFound, "inway with address not found").Err()
		}

		logger.Error("failed to set inway for organiation", zap.Error(err))

		return nil, status.New(codes.Internal, "database error").Err()
	}

	return &emptypb.Empty{}, nil
}

func (h *DirectoryRegistrationService) ClearOrganizationInway(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	logger := h.logger.With(zap.String("handler", "ClearOrganizationInway"))

	organizationName, err := h.getOrganisationNameFromRequest(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "determine organization name")
	}

	err = h.db.ClearOrganizationInway(ctx, organizationName)
	if err != nil {
		if errors.Is(err, database.ErrNoOrganization) {
			return nil, status.New(codes.NotFound, "organization not found").Err()
		}

		logger.Error("failed to clear inway for organiation", zap.Error(err))

		return nil, status.New(codes.Internal, "database error").Err()
	}

	return &emptypb.Empty{}, nil
}
