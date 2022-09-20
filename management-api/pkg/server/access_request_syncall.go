// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package server

import (
	"context"
	"sync"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/permissions"
	"go.nlx.io/nlx/management-api/pkg/syncer"
)

//nolint:funlen // not sure how to shorten this method
func (s *ManagementService) SynchronizeAllOutgoingAccessRequests(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	err := s.authorize(ctx, permissions.SyncOutgoingAccessRequests)
	if err != nil {
		return nil, err
	}

	s.logger.Info("rpc request SynchronizeAllLatestOutgoingAccessRequest")

	outgoingAccessRequests, err := s.configDatabase.ListAllLatestOutgoingAccessRequests(ctx)
	if err != nil {
		s.logger.Error("error getting latest access request states", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	if len(outgoingAccessRequests) < 1 {
		return &emptypb.Empty{}, nil
	}

	accessRequestsByOrganization := groupAccessRequestsByOrganization(outgoingAccessRequests)

	hasError := false
	waitGroup := sync.WaitGroup{}

	for oin, values := range accessRequestsByOrganization {
		accessRequests := values

		waitGroup.Add(1)

		go func(oin string) {
			defer waitGroup.Done()

			organizationInwayProxyAddress, err := s.directoryClient.GetOrganizationInwayProxyAddress(ctx, oin)
			if err != nil {
				s.logger.Error("cannot get organization inway proxy address", zap.Error(err))

				hasError = true

				return
			}

			client, err := s.createManagementClientFunc(ctx, organizationInwayProxyAddress, s.orgCert)
			if err != nil {
				s.logger.Error("cannot setup management client", zap.Error(err))

				hasError = true

				return
			}

			err = syncer.SyncOutgoingAccessRequests(&syncer.SyncArgs{
				Ctx:      ctx,
				Logger:   s.logger,
				Clock:    s.clock,
				DB:       s.configDatabase,
				Client:   client,
				Requests: accessRequests,
			})
			if err != nil {
				s.logger.Error("cannot setup management client", zap.Error(err))

				hasError = true

				return
			}

			client.Close()

			if err != nil {
				s.logger.Error("failed to close client", zap.Error(err))

				hasError = true

				return
			}
		}(oin)
	}

	waitGroup.Wait()

	if hasError {
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return &emptypb.Empty{}, nil
}

func groupAccessRequestsByOrganization(accessRequests []*database.OutgoingAccessRequest) map[string][]*database.OutgoingAccessRequest {
	result := map[string][]*database.OutgoingAccessRequest{}

	for _, ar := range accessRequests {
		_, ok := result[ar.Organization.SerialNumber]
		if !ok {
			result[ar.Organization.SerialNumber] = []*database.OutgoingAccessRequest{}
		}

		result[ar.Organization.SerialNumber] = append(result[ar.Organization.SerialNumber], ar)
	}

	return result
}
