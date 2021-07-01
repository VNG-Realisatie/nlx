// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server

import (
	"context"
	"errors"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
)

func (s *ManagementService) ListAccessGrantsForService(ctx context.Context, req *api.ListAccessGrantsForServiceRequest) (*api.ListAccessGrantsForServiceResponse, error) {
	_, err := s.configDatabase.GetService(ctx, req.ServiceName)
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
	userInfo, err := retrieveUserInfoFromGRPCContext(ctx)
	if err != nil {
		s.logger.Error("could not retrieve user info for audit log from grpc context", zap.Error(err))
		return nil, status.Error(codes.Internal, "could not retrieve user info to create audit log")
	}

	err = s.auditLogger.AccessGrantRevoke(ctx, userInfo.username, userInfo.userAgent, req.OrganizationName, req.ServiceName)
	if err != nil {
		return nil, status.Error(codes.Internal, "could not create audit log")
	}

	accessGrant, err := s.configDatabase.RevokeAccessGrant(ctx, uint(req.AccessGrantID), time.Now())
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

	var revokedAt *timestamp.Timestamp

	if accessGrant.RevokedAt.Valid {
		revokedAt = timestamppb.New(accessGrant.RevokedAt.Time)
	}

	return &api.AccessGrant{
		Id:                   uint64(accessGrant.ID),
		OrganizationName:     accessGrant.IncomingAccessRequest.OrganizationName,
		ServiceName:          accessGrant.IncomingAccessRequest.Service.Name,
		PublicKeyFingerprint: accessGrant.IncomingAccessRequest.PublicKeyFingerprint,
		CreatedAt:            createdAt,
		RevokedAt:            revokedAt,
	}
}
