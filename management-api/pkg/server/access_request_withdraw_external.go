// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package server

import (
	"context"
	"errors"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"

	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/grpcerrors"
)

func (s *ManagementService) WithdrawAccessRequest(ctx context.Context, req *external.WithdrawAccessRequestRequest) (*external.WithdrawAccessRequestResponse, error) {
	md, err := s.parseProxyMetadata(ctx)
	if err != nil {
		return nil, err
	}

	incomingAccessRequest, err := s.configDatabase.GetLatestIncomingAccessRequest(ctx, md.OrganizationSerialNumber, req.ServiceName, req.PublicKeyFingerprint)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			return nil, grpcerrors.New(codes.NotFound, external.ErrorReason_ERROR_REASON_ACCESS_REQUEST_NOT_FOUND, "access request could not be found", nil)
		}

		s.logger.Error("could not retrieve incoming access request from database", zap.Error(err))

		return nil, grpcerrors.NewInternal("internal", nil)
	}

	if incomingAccessRequest.State != database.IncomingAccessRequestReceived {
		return nil, grpcerrors.New(codes.FailedPrecondition, external.ErrorReason_ERROR_REASON_ACCESS_REQUEST_INVALID_STATE, fmt.Sprintf("expected state: %s, actual state: %s", database.IncomingAccessRequestReceived, incomingAccessRequest.State), nil)
	}

	err = s.configDatabase.UpdateIncomingAccessRequestState(ctx, incomingAccessRequest.ID, database.IncomingAccessRequestWithdrawn)
	if err != nil {
		s.logger.Error("could not cancel incoming access request", zap.Error(err))
		return nil, grpcerrors.NewInternal("internal", nil)
	}

	return &external.WithdrawAccessRequestResponse{}, nil
}
