// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package registrationservice

import (
	"context"
	"fmt"

	"github.com/gogo/protobuf/types"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/directory-registration-api/registrationapi"
)

func (h *DirectoryRegistrationService) SetInsightConfiguration(ctx context.Context, req *registrationapi.SetInsightConfigurationRequest) (*types.Empty, error) {
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

	return &types.Empty{}, nil
}
