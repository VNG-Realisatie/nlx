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
	"time"

	"github.com/gogo/protobuf/types"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"go.uber.org/zap"

	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/directory-inspection-api/inspectionapi"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/directory"
	"go.nlx.io/nlx/management-api/pkg/management"
	"go.nlx.io/nlx/management-api/pkg/util/clock"
)

const (
	maxRetries   = 3
	pollInterval = 5
)

var ErrMaxRetries = errors.New("unable to retry more than 3 times")

func computeInwayProxyAddress(address string) (string, error) {
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

type accessRequestScheduler struct {
	clock                      clock.Clock
	logger                     *zap.Logger
	directoryClient            directory.Client
	configDatabase             database.ConfigDatabase
	orgCert                    *common_tls.CertificateBundle
	requests                   chan *database.OutgoingAccessRequest
	createManagementClientFunc func(context.Context, string, *common_tls.CertificateBundle) (management.Client, error)
}

func newAccessRequestScheduler(logger *zap.Logger, directoryClient directory.Client, configDatabase database.ConfigDatabase, orgCert *common_tls.CertificateBundle) *accessRequestScheduler {
	return &accessRequestScheduler{
		clock:                      clock.RealClock{},
		logger:                     logger,
		orgCert:                    orgCert,
		directoryClient:            directoryClient,
		configDatabase:             configDatabase,
		requests:                   make(chan *database.OutgoingAccessRequest),
		createManagementClientFunc: management.NewClient,
	}
}

func (scheduler *accessRequestScheduler) listCurrentAccessRequests(ctx context.Context) error {
	requests, err := scheduler.configDatabase.ListAllOutgoingAccessRequests(ctx)
	if err != nil {
		return err
	}

	for _, request := range requests {
		if request.State == database.AccessRequestCreated {
			scheduler.requests <- request
		}
	}

	return nil
}

func (scheduler *accessRequestScheduler) Run(ctx context.Context) {
	go func() {
		if err := scheduler.listCurrentAccessRequests(ctx); err != nil {
			scheduler.logger.Error("failed to list current access requests", zap.Error(err))
		}

		scheduler.configDatabase.WatchOutgoingAccessRequests(ctx, scheduler.requests)
	}()

	wg := &sync.WaitGroup{}
	ticker := time.NewTicker(pollInterval * time.Second)

	defer ticker.Stop()

schedulingLoop:
	for {
		select {
		case <-ctx.Done():
			break schedulingLoop
		case <-ticker.C:
			wg.Add(1)

			go func() {
				defer wg.Done()

				if err := scheduler.schedulePendingRequests(context.TODO()); err != nil {
					scheduler.logger.Error("failed to schedule pending requests", zap.Error(err))
				}
			}()

			wg.Add(1)

			go func() {
				defer wg.Done()

				if err := scheduler.syncAccessProofs(context.TODO()); err != nil {
					scheduler.logger.Error("failed to sync access proofs", zap.Error(err))
				}
			}()
		case request := <-scheduler.requests:
			wg.Add(1)

			go func() {
				defer wg.Done()

				if err := scheduler.schedule(context.TODO(), request); err != nil {
					scheduler.logger.Error("failed to handle request", zap.Error(err))
				}
			}()
		}
	}

	wg.Wait()
}

func (scheduler *accessRequestScheduler) getOrganizationManagementClient(ctx context.Context, organizationName string) (management.Client, error) {
	response, err := scheduler.directoryClient.GetOrganizationInway(ctx, &inspectionapi.GetOrganizationInwayRequest{
		OrganizationName: organizationName,
	})
	if err != nil {
		return nil, err
	}

	address, err := computeInwayProxyAddress(response.Address)
	if err != nil {
		return nil, err
	}

	client, err := scheduler.createManagementClientFunc(ctx, address, scheduler.orgCert)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (scheduler *accessRequestScheduler) sendRequest(ctx context.Context, request *database.OutgoingAccessRequest) error {
	client, err := scheduler.getOrganizationManagementClient(ctx, request.OrganizationName)
	if err != nil {
		return err
	}

	defer client.Close()

	_, err = client.RequestAccess(ctx, &external.RequestAccessRequest{
		ServiceName: request.ServiceName,
	}, grpc_retry.WithMax(maxRetries))

	return err
}

func (scheduler *accessRequestScheduler) schedulePendingRequests(ctx context.Context) error {
	requests, err := scheduler.configDatabase.ListAllOutgoingAccessRequests(ctx)
	if err != nil {
		return err
	}

	for _, request := range requests {
		// only handle requests that are considered to be in-progress
		if request.State != database.AccessRequestReceived {
			continue
		}

		scheduler.requests <- request
	}

	return nil
}

func (scheduler *accessRequestScheduler) getAccessRequestState(ctx context.Context, request *database.OutgoingAccessRequest) (database.AccessRequestState, error) {
	client, err := scheduler.getOrganizationManagementClient(ctx, request.OrganizationName)
	if err != nil {
		return database.AccessRequestUnspecified, err
	}

	defer client.Close()

	response, err := client.GetAccessRequestState(ctx, &external.GetAccessRequestStateRequest{
		ServiceName: request.ServiceName,
	})
	if err != nil {
		return database.AccessRequestUnspecified, err
	}

	state := database.AccessRequestUnspecified

	switch response.State {
	case api.AccessRequestState_CREATED:
		state = database.AccessRequestCreated
	case api.AccessRequestState_APPROVED:
		state = database.AccessRequestApproved
	case api.AccessRequestState_REJECTED:
		state = database.AccessRequestRejected
	case api.AccessRequestState_RECEIVED:
		state = database.AccessRequestReceived
	case api.AccessRequestState_FAILED:
		state = database.AccessRequestFailed
	}

	return state, nil
}

func (scheduler *accessRequestScheduler) schedule(ctx context.Context, request *database.OutgoingAccessRequest) error {
	err := scheduler.configDatabase.LockOutgoingAccessRequest(ctx, request)
	switch err {
	case nil:
		break
	case database.ErrAccessRequestLocked:
		return nil
	default:
		return err
	}

	defer func() {
		if unlockErr := scheduler.configDatabase.UnlockOutgoingAccessRequest(ctx, request); unlockErr != nil {
			scheduler.logger.Error("failed to unlock outgoing access request", zap.Error(unlockErr))
		}
	}()

	var (
		taskErr  error
		newState database.AccessRequestState
	)

	switch request.State {
	case database.AccessRequestCreated:
		taskErr = scheduler.sendRequest(ctx, request)
		newState = database.AccessRequestReceived
	case database.AccessRequestReceived:
		newState, taskErr = scheduler.getAccessRequestState(ctx, request)
	default:
		return fmt.Errorf("invalid access request state: %s", request.State.String())
	}

	if taskErr != nil {
		newState = database.AccessRequestFailed
	}

	if err := scheduler.configDatabase.UpdateOutgoingAccessRequestState(ctx, request, newState); err != nil {
		return err
	}

	return taskErr
}

func (scheduler *accessRequestScheduler) parseAccessProof(accessProof *api.AccessProof) (*database.AccessProof, error) {
	createdAt, err := types.TimestampFromProto(accessProof.CreatedAt)
	if err != nil {
		return nil, err
	}

	revokedAt, err := types.TimestampFromProto(accessProof.RevokedAt)
	if err != nil {
		revokedAt = time.Time{}
	}

	return &database.AccessProof{
		ID:               accessProof.Id,
		CreatedAt:        createdAt,
		RevokedAt:        revokedAt,
		OrganizationName: accessProof.OrganizationName,
		ServiceName:      accessProof.ServiceName,
		AccessRequestID:  accessProof.AccessRequestId,
	}, nil
}

func (scheduler *accessRequestScheduler) syncAccessProof(ctx context.Context, request *database.OutgoingAccessRequest) error {
	client, err := scheduler.getOrganizationManagementClient(ctx, request.OrganizationName)
	if err != nil {
		return err
	}

	defer client.Close()

	response, err := client.GetAccessProof(ctx, &external.GetAccessProofRequest{
		ServiceName: request.ServiceName,
	})
	if err != nil {
		return err
	}

	remoteProof, err := scheduler.parseAccessProof(response)
	if err != nil {
		return err
	}

	localProof, err := scheduler.configDatabase.GetAccessProofForOutgoingAccessRequest(
		ctx,
		request.OrganizationName,
		request.ServiceName,
		request.ID,
	)

	switch err {
	case nil:
	case database.ErrNotFound:
		_, err = scheduler.configDatabase.CreateAccessProof(ctx, &database.AccessProof{
			OrganizationName: request.OrganizationName,
			ServiceName:      remoteProof.ServiceName,
			CreatedAt:        remoteProof.CreatedAt,
			AccessRequestID:  request.ID,
			RevokedAt:        remoteProof.RevokedAt,
		})

		return err
	default:
		return err
	}

	if !remoteProof.RevokedAt.IsZero() &&
		localProof.RevokedAt.IsZero() {
		if _, err := scheduler.configDatabase.RevokeAccessProof(
			ctx,
			localProof.OrganizationName,
			localProof.ServiceName,
			localProof.ID,
			remoteProof.RevokedAt,
		); err != nil {
			return err
		}
	}

	return nil
}

func (scheduler *accessRequestScheduler) syncAccessProofs(ctx context.Context) error {
	requests, err := scheduler.configDatabase.ListAllOutgoingAccessRequests(ctx)
	if err != nil {
		return err
	}

	for _, request := range requests {
		// only sync approved access requests
		if request.State != database.AccessRequestApproved {
			continue
		}

		if err := scheduler.syncAccessProof(ctx, request); err != nil {
			scheduler.logger.Error("failed to sync access proof", zap.Error(err))
		}
	}

	return nil
}
