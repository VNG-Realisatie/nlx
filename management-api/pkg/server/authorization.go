// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/pkg/authorization"
	"go.nlx.io/nlx/management-api/pkg/permissions"
)

func (s *ManagementService) authorize(ctx context.Context, permission permissions.Permission) error {
	userInfo, err := retrieveUserFromContext(ctx)
	if err != nil {
		return status.Error(codes.Internal, "could not retrieve user info to authorize user")
	}

	if !authorization.IsAuthorized(permission, userInfo.Permissions) {
		return status.Errorf(codes.PermissionDenied, "user needs the permission %q to execute this request", permission)
	}

	return nil
}
