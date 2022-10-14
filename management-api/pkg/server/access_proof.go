// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"context"

	"go.uber.org/zap"

	"go.nlx.io/nlx/management-api/api/external"
)

// Deprecated: use GetAccessGrant instead
func (s *ManagementService) GetAccessProof(ctx context.Context, req *external.GetAccessGrantRequest) (*external.AccessGrant, error) {
	md, err := s.parseProxyMetadata(ctx)
	if err != nil {
		s.logger.Error("failed to parse proxy metadata", zap.Error(err))

		return nil, err
	}

	s.logger.Warn("The organization is using deprecated GetAccessProof RPC. Please use GetAccessGrant instead.", zap.String("organization-name", md.OrganizationName), zap.String("organization-serial-number", md.OrganizationSerialNumber))

	res, err := s.GetAccessGrant(ctx, &external.GetAccessGrantRequest{
		ServiceName:          req.ServiceName,
		PublicKeyFingerprint: req.PublicKeyFingerprint,
	})
	if err != nil {
		return nil, err
	}

	return res.AccessGrant, nil
}
