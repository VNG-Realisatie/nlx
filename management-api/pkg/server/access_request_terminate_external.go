// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package server

import (
	"context"
	"errors"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/grpcerrors"
)

func (s *ManagementService) TerminateAccess(ctx context.Context, req *external.TerminateAccessRequest) (*external.TerminateAccessResponse, error) {
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

	if incomingAccessRequest.State != database.IncomingAccessRequestApproved {
		return nil, grpcerrors.New(codes.FailedPrecondition, external.ErrorReason_ERROR_REASON_ACCESS_REQUEST_INVALID_STATE, fmt.Sprintf("expected state: %s, actual state: %s", database.IncomingAccessRequestApproved, incomingAccessRequest.State), nil)
	}

	accessGrantID, err := s.configDatabase.GetAccessGrantIDForIncomingAccessRequest(ctx, incomingAccessRequest.ID)
	if err != nil {
		s.logger.Error("could not get access grant for incoming access request", zap.Error(err))
		return nil, grpcerrors.NewInternal("internal", nil)
	}

	terminatedAt := s.clock.Now()

	err = s.configDatabase.TerminateAccessGrant(ctx, accessGrantID, terminatedAt)
	if err != nil {
		if errors.Is(err, database.ErrAccessGrantAlreadyTerminated) {
			return nil, grpcerrors.New(codes.FailedPrecondition, external.ErrorReason_ERROR_REASON_ACCESS_GRANT_ALREADY_TERMINATED, "access grant already terminated", nil)
		}

		s.logger.Error("could not terminate access grant for incoming access request", zap.Error(err))

		return nil, grpcerrors.NewInternal("internal", nil)
	}

	return &external.TerminateAccessResponse{
		TerminatedAt: timestamppb.New(terminatedAt),
	}, nil
}
