// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package directory

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/nlxversion"
	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/directory-api/domain"
)

func (h *DirectoryService) RegisterOutway(ctx context.Context, req *directoryapi.RegisterOutwayRequest) (*directoryapi.RegisterOutwayResponse, error) {
	resp := &directoryapi.RegisterOutwayResponse{}

	organizationInformation, err := h.getOrganizationInformationFromRequest(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get organization information from request: %v", err)
	}

	nlxVersion := nlxversion.NewFromGRPCContext(ctx).Version

	// Created at and Updated at time are the same times when registering new outway
	now := time.Now()

	organizationModel, err := domain.NewOrganization(organizationInformation.Name, organizationInformation.SerialNumber)
	if err != nil {
		msg := fmt.Sprintf("validation failed: %s", err.Error())
		return nil, status.New(codes.InvalidArgument, msg).Err()
	}

	outwayModel, err := domain.NewOutway(&domain.NewOutwayArgs{
		Name:         req.Name,
		Organization: organizationModel,
		NlxVersion:   nlxVersion,
		CreatedAt:    now,
		UpdatedAt:    now,
	},
	)
	if err != nil {
		msg := fmt.Sprintf("validation failed: %s", err.Error())
		return nil, status.New(codes.InvalidArgument, msg).Err()
	}

	err = h.repository.RegisterOutway(outwayModel)
	if err != nil {
		h.logger.Error("register outway", zap.String("outway", outwayModel.ToString()), zap.Error(err))
		return nil, status.New(codes.Internal, "failed to register outway").Err()
	}

	return resp, nil
}
