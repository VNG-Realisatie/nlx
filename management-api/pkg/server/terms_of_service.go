// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package server

import (
	"context"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/permissions"
)

// GetTermsOfServiceStatus returns a Terms of Service status
func (s *ManagementService) GetTermsOfServiceStatus(ctx context.Context, _ *emptypb.Empty) (*api.GetTermsOfServiceStatusResponse, error) {
	err := s.authorize(ctx, permissions.ReadTermsOfServiceStatus)
	if err != nil {
		return nil, err
	}

	logger := s.logger
	logger.Info("rpc request GetTermsOfServiceStatus")

	_, err = s.configDatabase.GetTermsOfServiceStatus(ctx)
	if err != nil {
		if errIsNotFound(err) {
			return &api.GetTermsOfServiceStatusResponse{
				Accepted: false,
			}, nil
		}

		logger.Error("error getting Terms of Service from DB", zap.Error(err))

		return nil, status.Error(codes.Internal, "database error")
	}

	return &api.GetTermsOfServiceStatusResponse{
		Accepted: true,
	}, nil
}

// AcceptTermsOfService accepts the Terms of Service
func (s *ManagementService) AcceptTermsOfService(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	err := s.authorize(ctx, permissions.AcceptTermsOfService)
	if err != nil {
		return nil, err
	}

	logger := s.logger
	logger.Info("rpc request AcceptTermsOfService")

	userInfo, err := retrieveUserFromContext(ctx)
	if err != nil {
		s.logger.Error("could not retrieve user info for audit log from grpc context", zap.Error(err))
		return nil, status.Error(codes.Internal, "could not retrieve user info to create audit log")
	}

	alreadyAccepted, err := s.configDatabase.AcceptTermsOfService(ctx, userInfo.Email, time.Now())
	if err != nil {
		logger.Error("error accepting Terms of Service from DB", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	if alreadyAccepted {
		return &emptypb.Empty{}, nil
	}

	err = s.auditLogger.AcceptTermsOfService(ctx, userInfo.Email, userInfo.UserAgent)
	if err != nil {
		return nil, status.Error(codes.Internal, "could not create audit log")
	}

	return &emptypb.Empty{}, nil
}
