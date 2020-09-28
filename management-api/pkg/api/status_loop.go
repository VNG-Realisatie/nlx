// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package api

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strconv"
	"sync"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	common_tls "go.nlx.io/nlx/common/tls"
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
	clock                      clock.Clock
	logger                     *zap.Logger
	directoryClient            directory.Client
	configDatabase             database.ConfigDatabase
	orgCert                    *common_tls.CertificateBundle
	requests                   chan *database.OutgoingAccessRequest
	createManagementClientFunc func(context.Context, string, *common_tls.CertificateBundle) (management.Client, error)
}

func newAccessRequestStatusLoop(logger *zap.Logger, directoryClient directory.Client, configDatabase database.ConfigDatabase, orgCert *common_tls.CertificateBundle) *accessRequestStatusLoop {
	return &accessRequestStatusLoop{
		clock:                      clock.RealClock{},
		logger:                     logger,
		orgCert:                    orgCert,
		directoryClient:            directoryClient,
		configDatabase:             configDatabase,
		requests:                   make(chan *database.OutgoingAccessRequest),
		createManagementClientFunc: management.NewClient,
	}
}

func (loop *accessRequestStatusLoop) listCurrentAccessRequests(ctx context.Context) error {
	requests, err := loop.configDatabase.ListAllOutgoingAccessRequests(ctx)
	if err != nil {
		return err
	}

	for _, request := range requests {
		if request.State == database.AccessRequestCreated {
			loop.requests <- request
		}
	}

	return nil
}

func (loop *accessRequestStatusLoop) Run(ctx context.Context) {
	go func() {
		if err := loop.listCurrentAccessRequests(ctx); err != nil {
			loop.logger.Error("failed to list current access requests", zap.Error(err))
		}

		loop.configDatabase.WatchOutgoingAccessRequests(ctx, loop.requests)
	}()

	wg := &sync.WaitGroup{}

statusLoop:
	for {
		select {
		case <-ctx.Done():
			break statusLoop
		case request := <-loop.requests:
			requestCtx := context.Background()
			wg.Add(1)

			go func(c context.Context, r *database.OutgoingAccessRequest) {
				if err := loop.handleRequest(c, r); err != nil {
					loop.logger.Error("failed to handle request", zap.Error(err))
				}

				wg.Done()
			}(requestCtx, request)
		}
	}

	wg.Wait()
}

func (loop *accessRequestStatusLoop) computeInwayProxyAddress(address string) (string, error) {
	host, port, err := net.SplitHostPort(address)
	if err != nil {
		return "", fmt.Errorf("invalid format for inway address: %w", err)
	}

	portNum, err := strconv.Atoi(port)
	if err != nil {
		return "", fmt.Errorf("invalid format for inway address port: %w", err)
	}

	return fmt.Sprintf("%s:%d", host, portNum+1), nil
}

func (loop *accessRequestStatusLoop) sendRequest(ctx context.Context, request *database.OutgoingAccessRequest) error {
	response, err := loop.directoryClient.GetOrganizationInway(ctx, &inspectionapi.GetOrganizationInwayRequest{
		OrganizationName: request.OrganizationName,
	})
	if err != nil {
		return err
	}

	address, err := loop.computeInwayProxyAddress(response.Address)
	if err != nil {
		return err
	}

	client, err := loop.createManagementClientFunc(ctx, address, loop.orgCert)
	if err != nil {
		return err
	}

	defer client.Close()

	_, err = client.RequestAccess(ctx, &external.RequestAccessRequest{
		ServiceName: request.ServiceName,
	}, grpc_retry.WithMax(maxRetries))

	return err
}

//nolint:gocyclo // high complexity because of the state machine
func (loop *accessRequestStatusLoop) handleRequest(ctx context.Context, request *database.OutgoingAccessRequest) error {
	err := loop.configDatabase.LockOutgoingAccessRequest(ctx, request)
	switch err {
	case nil:
		break
	case database.ErrAccessRequestLocked:
		return nil
	default:
		return err
	}

	defer func() {
		if unlockErr := loop.configDatabase.UnlockOutgoingAccessRequest(ctx, request); unlockErr != nil {
			loop.logger.Error("failed to unlock outgoing access request", zap.Error(unlockErr))
		}
	}()

	sendErr := loop.sendRequest(ctx, request)
	st, isStatusErr := status.FromError(sendErr)
	var state database.AccessRequestState

	if sendErr == nil {
		state = database.AccessRequestReceived
	} else if isStatusErr && st.Code() == codes.AlreadyExists {
		state = database.AccessRequestFailed
	} else {
		state = database.AccessRequestFailed
	}

	if err := loop.configDatabase.UpdateOutgoingAccessRequestState(ctx, request, state); err != nil {
		return err
	}

	return sendErr
}
