// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package server

import (
	"context"
	"strings"
	"sync"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	common_grpcerrors "go.nlx.io/nlx/common/grpcerrors"
	common_tls "go.nlx.io/nlx/common/tls"
	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/directory"
	"go.nlx.io/nlx/management-api/pkg/grpcerrors"
	"go.nlx.io/nlx/management-api/pkg/management"
	"go.nlx.io/nlx/management-api/pkg/permissions"
	"go.nlx.io/nlx/management-api/pkg/syncer"
)

func (s *ManagementService) SynchronizeAllOutgoingAccessRequests(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
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
		organizations, err := s.directoryClient.ListOrganizations(ctx, &directoryapi.ListOrganizationsRequest{})
		if err != nil {
			s.logger.Error("failed to retrieve organizations from directory", zap.Error(err))
			return nil, status.Error(codes.Internal, "internal error")
		}

		oinToOrgNameHash := convertOrganizationsToHash(organizations)

		for i, oin := range oinsWithError {
			oinsWithError[i] = oinToOrgNameHash[oin]
		}

		return nil, grpcerrors.NewInternal("unreachable organizations", &common_grpcerrors.Metadata{
			Metadata: map[string]string{
				"organizations": strings.Join(oinsWithError, ", "),
			},
		})
	}

	return &emptypb.Empty{}, nil
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

func synchronizeAccessRequests(args *syncArgs) []string {
	var oinsWithErrorMutex sync.Mutex

	oinsWithError := []string{}
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
				oinsWithError = append(oinsWithError, oin)
				m.Unlock()

				return
			}

			client, err := args.createManagementClientFunc(args.ctx, organizationInwayProxyAddress, args.orgCert)
			if err != nil {
				args.l.Error("cannot setup management client", zap.Error(err))

				m.Lock()
				oinsWithError = append(oinsWithError, oin)
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
				args.l.Error("cannot setup management client", zap.Error(err))

				m.Lock()
				oinsWithError = append(oinsWithError, oin)
				m.Unlock()

				return
			}

			client.Close()

			if err != nil {
				args.l.Error("failed to close client", zap.Error(err))

				m.Lock()
				oinsWithError = append(oinsWithError, oin)
				m.Unlock()

				return
			}
		}(oin, &oinsWithErrorMutex)
	}

	waitGroup.Wait()

	return oinsWithError
}
