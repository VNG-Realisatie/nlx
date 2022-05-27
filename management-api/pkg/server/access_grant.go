// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/permissions"
)

func (s *ManagementService) ListAccessGrantsForService(ctx context.Context, req *api.ListAccessGrantsForServiceRequest) (*api.ListAccessGrantsForServiceResponse, error) {
	err := s.authorize(ctx, permissions.ReadAccessGrants)
	if err != nil {
		return nil, err
	}

	_, err = s.configDatabase.GetService(ctx, req.ServiceName)
	if err != nil {
		if errIsNotFound(err) {
			return nil, status.Error(codes.NotFound, "service not found")
		}

		s.logger.Error("error fetching service", zap.String("name", req.ServiceName), zap.Error(err))

		return nil, status.Error(codes.Internal, "database error")
	}

	accessRequests, err := s.configDatabase.ListAccessGrantsForService(ctx, req.ServiceName)
	if err != nil {
		s.logger.Error("fetching access grants", zap.String("service name", req.ServiceName), zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	response := &api.ListAccessGrantsForServiceResponse{
		AccessGrants: make([]*api.AccessGrant, len(accessRequests)),
	}

	for i, accessGrant := range accessRequests {
		responseAccessGrant := convertAccessGrant(accessGrant)
		response.AccessGrants[i] = responseAccessGrant
	}

	return response, nil
}

func (s *ManagementService) RevokeAccessGrant(ctx context.Context, req *api.RevokeAccessGrantRequest) (*api.AccessGrant, error) {
	err := s.authorize(ctx, permissions.RevokeAccessGrant)
	if err != nil {
		return nil, err
	}

	userInfo, err := retrieveUserFromContext(ctx)
	if err != nil {
		s.logger.Error("could not retrieve user info for audit log from grpc context", zap.Error(err))
		return nil, status.Error(codes.Internal, "could not retrieve user info to create audit log")
	}

	accessGrant, err := s.configDatabase.GetAccessGrant(ctx, uint(req.AccessGrantID))
	if err != nil {
		if err == database.ErrNotFound {
			return nil, status.Error(codes.NotFound, fmt.Sprintf("access grant with id:%d does not exist", req.AccessGrantID))
		}

		s.logger.Error("cannot get access grant from database", zap.Uint64("access grant id", req.AccessGrantID), zap.Error(err))

		return nil, status.Error(codes.Internal, "internal error")
	}

	err = s.auditLogger.AccessGrantRevoke(ctx, userInfo.Email, userInfo.UserAgent, accessGrant.IncomingAccessRequest.Organization.SerialNumber, accessGrant.IncomingAccessRequest.Organization.Name, accessGrant.IncomingAccessRequest.Service.Name)
	if err != nil {
		return nil, status.Error(codes.Internal, "could not create audit log")
	}

	accessGrant, err = s.configDatabase.RevokeAccessGrant(ctx, uint(req.AccessGrantID), time.Now())
	if err != nil {
		if errors.Is(err, database.ErrAccessGrantAlreadyRevoked) {
			s.logger.Warn("access grant is already revoked")
			return nil, status.Error(codes.AlreadyExists, "access grant is already revoked")
		}

		if errIsNotFound(err) {
			s.logger.Warn("access grant not found")
			return nil, status.Error(codes.NotFound, "access grant not found")
		}

		s.logger.Error("revoking access grant", zap.Error(err))

		return nil, status.Error(codes.Internal, "database error")
	}

	apiAccessGrant := convertAccessGrant(accessGrant)

	return apiAccessGrant, nil
}

func convertAccessGrant(accessGrant *database.AccessGrant) *api.AccessGrant {
	createdAt := timestamppb.New(accessGrant.CreatedAt)

	var revokedAt *timestamppb.Timestamp

	if accessGrant.RevokedAt.Valid {
		revokedAt = timestamppb.New(accessGrant.RevokedAt.Time)
	}

	return &api.AccessGrant{
		Id: uint64(accessGrant.ID),
		Organization: &api.Organization{
			Name:         accessGrant.IncomingAccessRequest.Organization.Name,
			SerialNumber: accessGrant.IncomingAccessRequest.Organization.SerialNumber,
		},
		ServiceName:          accessGrant.IncomingAccessRequest.Service.Name,
		PublicKeyFingerprint: accessGrant.IncomingAccessRequest.PublicKeyFingerprint,
		CreatedAt:            createdAt,
		RevokedAt:            revokedAt,
	}
}
