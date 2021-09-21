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

	organization, err := h.getOrganizationInformationFromRequest(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "determine organization")
	}

	logger.Debug("clearing organization inway", zap.Any("organization", organization))

	err = h.repository.ClearOrganizationInway(ctx, organization.SerialNumber)
	if err != nil {
		if errors.Is(err, adapters.ErrOrganizationNotFound) {
			logger.Info("did not clear organization because the organization could not be found", zap.Any("organization", organization))
			return &emptypb.Empty{}, nil
		}

		logger.Error("failed to clear the organization inway", zap.Error(err))

		return nil, status.New(codes.Internal, "database error").Err()
	}

	return &emptypb.Empty{}, nil
}
