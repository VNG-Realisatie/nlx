// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package api

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"go.uber.org/zap"
	"golang.org/x/sync/semaphore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/diagnostics"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/directory"
	"go.nlx.io/nlx/management-api/pkg/management"
	"go.nlx.io/nlx/management-api/pkg/util/clock"
)

const (
	maxRetries     = 3
	maxConcurrency = 4
	pollInterval   = 1500 * time.Millisecond

	// jobs are unlocked after 5 minutes, let's wait at least one minute before retrying
	jobTimeout = 4 * time.Minute

	errMessageServiceNoLongerExists = "service no longer exists"
)

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

func (scheduler *accessRequestScheduler) Run(ctx context.Context) {
	wg := &sync.WaitGroup{}
	sem := semaphore.NewWeighted(int64(maxConcurrency))
	ticker := time.NewTicker(pollInterval)

	defer ticker.Stop()

schedulingLoop:
	for {
		select {
		case <-ctx.Done():
			break schedulingLoop
		case <-ticker.C:
			go func() {
				if !sem.TryAcquire(1) {
					return
				}

				wg.Add(1)

				defer sem.Release(1)
				defer wg.Done()
				if err := scheduler.schedulePendingRequest(context.TODO()); err != nil {
					scheduler.logger.Error("failed to schedule pending request", zap.Error(err))
				}
			}()
		}
	}

	wg.Wait()
}

func (scheduler *accessRequestScheduler) getOrganizationManagementClient(ctx context.Context, organizationName string) (management.Client, error) {
	address, err := scheduler.directoryClient.GetOrganizationInwayProxyAddress(ctx, organizationName)
	if err != nil {
		return nil, err
	}

	client, err := scheduler.createManagementClientFunc(ctx, address, scheduler.orgCert)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (scheduler *accessRequestScheduler) sendRequest(ctx context.Context, request *database.OutgoingAccessRequest) (uint, error) {
	client, err := scheduler.getOrganizationManagementClient(ctx, request.OrganizationName)
	if err != nil {
		return 0, err
	}

	defer client.Close()

	response, err := client.RequestAccess(ctx, &external.RequestAccessRequest{
		ServiceName: request.ServiceName,
	}, grpc_retry.WithMax(maxRetries))
	if err != nil {
		return 0, err
	}

	return uint(response.ReferenceId), nil
}

func (scheduler *accessRequestScheduler) schedulePendingRequest(ctx context.Context) error {
	request, err := scheduler.configDatabase.TakePendingOutgoingAccessRequest(ctx)
	if err != nil {
		return err
	}

	if request != nil {
		jobCtx, cancel := context.WithTimeout(ctx, jobTimeout)

		defer cancel()

		switch request.State {
		case database.OutgoingAccessRequestCreated, database.OutgoingAccessRequestReceived:
			if err := scheduler.schedule(jobCtx, request); err != nil {
				return err
			}
		case database.OutgoingAccessRequestApproved:
			if err := scheduler.syncAccessProof(jobCtx, request); err != nil {
				return err
			}
		default:
			return fmt.Errorf("invalid status %s for pending access request", request.State)
		}

		return scheduler.configDatabase.UnlockOutgoingAccessRequest(ctx, request)
	}

	return nil
}

func (scheduler *accessRequestScheduler) getAccessRequestState(ctx context.Context, request *database.OutgoingAccessRequest) (database.OutgoingAccessRequestState, error) {
	client, err := scheduler.getOrganizationManagementClient(ctx, request.OrganizationName)
	if err != nil {
		return "", err
	}

	defer client.Close()

	response, err := client.GetAccessRequestState(ctx, &external.GetAccessRequestStateRequest{
		ServiceName: request.ServiceName,
	})
	if err != nil {
		return "", err
	}

	var state database.OutgoingAccessRequestState

	switch response.State {
	case api.AccessRequestState_CREATED:
		state = database.OutgoingAccessRequestCreated
	case api.AccessRequestState_APPROVED:
		state = database.OutgoingAccessRequestApproved
	case api.AccessRequestState_REJECTED:
		state = database.OutgoingAccessRequestRejected
	case api.AccessRequestState_RECEIVED:
		state = database.OutgoingAccessRequestReceived
	case api.AccessRequestState_FAILED:
		state = database.OutgoingAccessRequestFailed
	default:
		return "", fmt.Errorf("invalid state for outgoing access request: %s", response.State)
	}

	return state, nil
}

func (scheduler *accessRequestScheduler) schedule(ctx context.Context, request *database.OutgoingAccessRequest) error {
	var (
		referenceID uint
		err         error
		newState    database.OutgoingAccessRequestState
	)

	switch request.State {
	case database.OutgoingAccessRequestCreated:
		referenceID, err = scheduler.sendRequest(ctx, request)
		newState = database.OutgoingAccessRequestReceived
	case database.OutgoingAccessRequestReceived:
		state, getStateErr := scheduler.getAccessRequestState(ctx, request)
		if getStateErr != nil {
			return getStateErr
		}

		newState = state
	}

	if err == nil {
		err = scheduler.configDatabase.UpdateOutgoingAccessRequestState(
			ctx,
			request.ID,
			newState,
			referenceID,
			nil,
		)
	} else {
		errorDetails := diagnostics.ParseError(err)

		st, ok := status.FromError(err)
		if ok {
			if st.Code() == codes.NotFound {
				errorDetails = errorDetails.WithCode(diagnostics.NoInwaySelectedError)
			}
		}

		err = scheduler.configDatabase.UpdateOutgoingAccessRequestState(
			ctx,
			request.ID,
			database.OutgoingAccessRequestFailed,
			referenceID,
			errorDetails,
		)
	}

	return err
}

func (scheduler *accessRequestScheduler) parseAccessProof(accessProof *api.AccessProof) (*database.AccessProof, error) {
	var createdAt time.Time

	if accessProof.CreatedAt != nil {
		var err error

		createdAt, err = ptypes.Timestamp(accessProof.CreatedAt)
		if err != nil {
			return nil, err
		}
	}

	revokedAt, err := ptypes.Timestamp(accessProof.RevokedAt)
	if err != nil {
		revokedAt = time.Time{}
	}

	return &database.AccessProof{
		ID:        uint(accessProof.Id),
		CreatedAt: createdAt,
		RevokedAt: sql.NullTime{
			Time:  revokedAt,
			Valid: !revokedAt.IsZero(),
		},
		OutgoingAccessRequest: &database.OutgoingAccessRequest{
			ID:               uint(accessProof.AccessRequestId),
			OrganizationName: accessProof.OrganizationName,
			ServiceName:      accessProof.ServiceName,
		},
	}, nil
}

func (scheduler *accessRequestScheduler) syncAccessProof(ctx context.Context, outgoingAccessRequest *database.OutgoingAccessRequest) error {
	remoteProof, err := scheduler.retrieveAccessProof(ctx, outgoingAccessRequest.OrganizationName, outgoingAccessRequest.ServiceName)
	if err != nil {
		grpcErr, ok := status.FromError(err)
		if !ok {
			return err
		}

		if grpcErr.Message() == errMessageServiceNoLongerExists {
			return scheduler.configDatabase.DeleteOutgoingAccessRequests(ctx, outgoingAccessRequest.OrganizationName, outgoingAccessRequest.ServiceName)
		}

		return err
	}

	// skip this AccessRequest as it's not the one related to this AccessProof
	if remoteProof.OutgoingAccessRequest.ID != outgoingAccessRequest.ReferenceID {
		return nil
	}

	localProof, err := scheduler.configDatabase.GetAccessProofForOutgoingAccessRequest(
		ctx,
		outgoingAccessRequest.ID,
	)

	switch err {
	case nil:
	case database.ErrNotFound:
		_, err = scheduler.configDatabase.CreateAccessProof(ctx, outgoingAccessRequest)

		return err
	default:
		return err
	}

	if remoteProof.RevokedAt.Valid &&
		!localProof.RevokedAt.Valid {
		if _, err := scheduler.configDatabase.RevokeAccessProof(
			ctx,
			localProof.ID,
			remoteProof.RevokedAt.Time,
		); err != nil {
			return err
		}
	}

	return nil
}

func (scheduler *accessRequestScheduler) retrieveAccessProof(ctx context.Context, organizationName, serviceName string) (*database.AccessProof, error) {
	client, err := scheduler.getOrganizationManagementClient(ctx, organizationName)
	if err != nil {
		return nil, err
	}

	defer client.Close()

	response, err := client.GetAccessProof(ctx, &external.GetAccessProofRequest{
		ServiceName: serviceName,
	})
	if err != nil {
		return nil, err
	}

	return scheduler.parseAccessProof(response)
}
