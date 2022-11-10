// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package server

import (
	"context"
	"sync"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	common_grpcerrors "go.nlx.io/nlx/common/grpcerrors"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/directory"
	"go.nlx.io/nlx/management-api/pkg/grpcerrors"
	"go.nlx.io/nlx/management-api/pkg/management"
	"go.nlx.io/nlx/management-api/pkg/permissions"
	"go.nlx.io/nlx/management-api/pkg/syncer"
)

func (s *ManagementService) SynchronizeAllOutgoingAccessRequests(ctx context.Context, _ *api.SynchronizeAllOutgoingAccessRequestsRequest) (*api.SynchronizeAllOutgoingAccessRequestsResponse, error) {
	err := s.authorize(ctx, permissions.SyncOutgoingAccessRequests)
	if err != nil {
		return nil, err
	}

	s.logger.Info("rpc request SynchronizeAllLatestOutgoingAccessRequest")

	outgoingAccessRequests, err := s.configDatabase.ListAllLatestOutgoingAccessRequests(ctx)
	if err != nil {
		s.logger.Error("error getting latest access request states", zap.Error(err))
		return nil, status.Errorf(codes.Internal, string(InternalError))
	}

	if len(outgoingAccessRequests) < 1 {
		return &api.SynchronizeAllOutgoingAccessRequestsResponse{}, nil
	}

	accessRequestsByOin := groupAccessRequestsByOin(outgoingAccessRequests)

	oinsWithError := synchronizeAccessRequests(&syncArgs{
		ctx:                        ctx,
		l:                          s.logger,
		dc:                         s.directoryClient,
		db:                         s.configDatabase,
		orgCert:                    s.orgCert,
		createManagementClientFunc: s.createManagementClientFunc,
		accessRequestsByOin:        accessRequestsByOin,
	})

	if len(oinsWithError) > 0 {
		metadata := map[string]string{}

		for oin, syncError := range oinsWithError {
			metadata[oin] = string(syncError)
		}

		return nil, grpcerrors.NewInternal("unreachable organizations", &common_grpcerrors.Metadata{
			Metadata: metadata,
		})
	}

	return &api.SynchronizeAllOutgoingAccessRequestsResponse{}, nil
}

func groupAccessRequestsByOin(accessRequests []*database.OutgoingAccessRequest) map[string][]*database.OutgoingAccessRequest {
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

type syncArgs struct {
	ctx                        context.Context
	l                          *zap.Logger
	dc                         directory.Client
	db                         database.ConfigDatabase
	orgCert                    *common_tls.CertificateBundle
	createManagementClientFunc func(context.Context, string, *common_tls.CertificateBundle) (management.Client, error)
	accessRequestsByOin        map[string][]*database.OutgoingAccessRequest
}

//nolint:funlen // unable to shorten method
func synchronizeAccessRequests(args *syncArgs) map[string]SyncError {
	var oinsWithErrorMutex sync.Mutex

	oinsWithError := map[string]SyncError{}
	waitGroup := sync.WaitGroup{}

	for oin, values := range args.accessRequestsByOin {
		accessRequests := values

		waitGroup.Add(1)

		go func(oin string, m *sync.Mutex) {
			defer waitGroup.Done()

			organizationInwayProxyAddress, err := args.dc.GetOrganizationInwayProxyAddress(args.ctx, oin)
			if err != nil {
				args.l.Error("cannot get organization inway proxy address", zap.Error(err))

				m.Lock()
				oinsWithError[oin] = InternalError
				m.Unlock()

				return
			}

			if organizationInwayProxyAddress == "" {
				m.Lock()
				oinsWithError[oin] = ServiceProviderNoOrganizationInwaySpecified
				m.Unlock()

				return
			}

			client, err := args.createManagementClientFunc(args.ctx, organizationInwayProxyAddress, args.orgCert)
			if err != nil {
				args.l.Error("cannot setup management client", zap.Error(err))

				m.Lock()
				oinsWithError[oin] = ServiceProviderOrganizationInwayUnreachable
				m.Unlock()

				return
			}

			err = syncer.SyncOutgoingAccessRequests(&syncer.SyncArgs{
				Ctx:      args.ctx,
				Logger:   args.l,
				DB:       args.db,
				Client:   client,
				Requests: accessRequests,
			})
			if err != nil {
				args.l.Error("cannot sync outgoing access requests", zap.Error(err))

				m.Lock()
				oinsWithError[oin] = InternalError
				m.Unlock()

				return
			}

			client.Close()

			if err != nil {
				args.l.Error("failed to close client", zap.Error(err))

				m.Lock()
				oinsWithError[oin] = InternalError
				m.Unlock()

				return
			}
		}(oin, &oinsWithErrorMutex)
	}

	waitGroup.Wait()

	return oinsWithError
}
