// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package api

import (
	"context"
	"errors"
	"sync"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/directory-inspection-api/inspectionapi"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/directory"
	"go.nlx.io/nlx/management-api/pkg/management"
	"go.nlx.io/nlx/management-api/pkg/util/clock"
)

const maxRetries = 3

var ErrMaxRetries = errors.New("unable to retry more than 3 times")

type accessRequestStatusLoop struct {
	clock           clock.Clock
	logger          *zap.Logger
	directoryClient directory.Client
	configDatabase  database.ConfigDatabase
	orgCert         *tls.CertificateBundle
	requests        chan *database.AccessRequest
}

func newAccessRequestStatusLoop(logger *zap.Logger, directoryClient directory.Client, configDatabase database.ConfigDatabase, orgCert *tls.CertificateBundle) *accessRequestStatusLoop {
	return &accessRequestStatusLoop{
		clock:           clock.RealClock{},
		logger:          logger,
		orgCert:         orgCert,
		directoryClient: directoryClient,
		configDatabase:  configDatabase,
		requests:        make(chan *database.AccessRequest),
	}
}

func (loop *accessRequestStatusLoop) streamOutgoingAccessRequests(ctx context.Context) error {
	requests, err := loop.configDatabase.ListAllOutgoingAccessRequests(ctx)
	if err != nil {
		return err
	}

	for _, request := range requests {
		if request.State == database.AccessRequestCreated {
			loop.requests <- request
		}
	}

	return loop.configDatabase.WatchOutgoingAccessRequests(ctx, loop.requests)
}

func (loop *accessRequestStatusLoop) Run(ctx context.Context) {
	go loop.streamOutgoingAccessRequests(ctx)

	wc := &sync.WaitGroup{}

statusLoop:
	for {
		select {
		case <-ctx.Done():
			break statusLoop
		case request := <-loop.requests:
			requestCtx := context.Background()
			err := loop.configDatabase.LockOutgoingAccessRequest(requestCtx, request)

			switch err {
			case nil:
				wc.Add(1)

				go func(c context.Context, r *database.AccessRequest) {
					if err := loop.handleRequest(c, r); err != nil {
						loop.logger.Error("failed to handle request", zap.Error(err))
					}

					wc.Done()
				}(requestCtx, request)
			case database.ErrAccessRequestLocked:
				continue statusLoop
			default:
				loop.logger.Warn("failed to lock access request", zap.Error(err))
			}
		}
	}

	wc.Wait()
}

func (loop *accessRequestStatusLoop) handleRequest(ctx context.Context, request *database.AccessRequest) error {
	defer loop.configDatabase.UnlockOutgoingAccessRequest(ctx, request)

	response, err := loop.directoryClient.GetOrganizationInway(ctx, &inspectionapi.GetOrganizationInwayRequest{
		OrganizationName: request.OrganizationName,
	})
	if err != nil {
		return err
	}

	client, err := management.NewClient(ctx, response.Address, loop.orgCert)
	if err != nil {
		return err
	}

	defer client.Close()

	_, err = client.RequestAccess(ctx, &external.RequestAccessRequest{
		ServiceName: request.ServiceName,
	}, grpc_retry.WithMax(maxRetries))
	st, isStatusErr := status.FromError(err)

	if isStatusErr && st.Code() == codes.AlreadyExists {
		panic("sync not implemented yet")
	} else if err != nil {
		return err
	}

	if err := loop.configDatabase.UpdateAccessRequestState(ctx, request, database.AccessRequestReceived); err != nil {
		return err
	}

	return nil
}
