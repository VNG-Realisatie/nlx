// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package registrationservice

import (
	"context"
	"errors"
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
		if errors.Is(err, database.ErrOrganizationNotFound) {
			return nil, status.New(codes.NotFound, "organization not found").Err()
		}

		logger.Error("failed to clear inway for organiation", zap.Error(err))

		return nil, status.New(codes.Internal, "database error").Err()
	}

	return &emptypb.Empty{}, nil
}
