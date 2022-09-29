// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package server

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/domain"
	"go.nlx.io/nlx/management-api/pkg/permissions"
)

func (s *ManagementService) authorize(ctx context.Context, permission permissions.Permission) error {
	user, _, err := retrieveUserFromContext(ctx)
	if err != nil {
		s.logger.Warn("could not retrieve user info to authorize user", zap.Error(err))
		return status.Error(codes.Internal, "could not retrieve user info to authorize user")
	}

	if !isAuthorized(permission, user) {
		return status.Errorf(codes.PermissionDenied, "user needs the permission %q to execute this request", permission)
	}

	return nil
}

func isAuthorized(permission permissions.Permission, user *domain.User) bool {
	for _, role := range user.Roles {
		for _, p := range role.Permissions {
			if p == permission {
				return true
			}
		}
	}

	return false
}
