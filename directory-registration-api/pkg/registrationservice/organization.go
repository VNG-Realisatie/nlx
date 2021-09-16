// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package registrationservice

import (
	"context"
	"errors"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"go.nlx.io/nlx/directory-registration-api/adapters"
)

func (h *DirectoryRegistrationService) ClearOrganizationInway(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	logger := h.logger.With(zap.String("handler", "ClearOrganizationInway"))

	organizationName, err := h.getOrganisationNameFromRequest(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "determine organization name")
	}

	err = h.repository.ClearOrganizationInway(ctx, organizationName)
	if err != nil {
		if errors.Is(err, adapters.ErrOrganizationNotFound) {
			return &emptypb.Empty{}, nil
		}

		logger.Error("failed to clear the organization inway", zap.Error(err))

		return nil, status.New(codes.Internal, "database error").Err()
	}

	return &emptypb.Empty{}, nil
}
