// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package scheduler

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/diagnostics"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/directory"
	"go.nlx.io/nlx/management-api/pkg/management"
	"go.nlx.io/nlx/management-api/pkg/server"
)

var ErrInvalidTimeStamp = fmt.Errorf("invalid timestamp")

type CreateManagementClientFunc func(context.Context, string, *common_tls.CertificateBundle) (management.Client, error)

type SynchronizeOutgoingAccessRequestJob struct {
	ctx                        context.Context
	orgCert                    *common_tls.CertificateBundle
	directoryClient            directory.Client
	configDatabase             database.ConfigDatabase
	createManagementClientFunc CreateManagementClientFunc
}

func NewSynchronizeOutgoingAccessRequestJob(ctx context.Context, directoryClient directory.Client, configDatabase database.ConfigDatabase, orgCert *common_tls.CertificateBundle, createManagementClientFunc CreateManagementClientFunc) *SynchronizeOutgoingAccessRequestJob {
	return &SynchronizeOutgoingAccessRequestJob{
		ctx:                        ctx,
		orgCert:                    orgCert,
		directoryClient:            directoryClient,
		configDatabase:             configDatabase,
		createManagementClientFunc: createManagementClientFunc,
	}
}

func (job *SynchronizeOutgoingAccessRequestJob) Run(ctx context.Context) error {
	request, err := job.configDatabase.TakePendingOutgoingAccessRequest(ctx)
	if err != nil {
		return err
	}

	if request == nil {
		return nil
	}

	jobCtx, cancel := context.WithTimeout(ctx, jobTimeout)

	defer cancel()

	err = job.synchronize(jobCtx, request)
	if err != nil {
		return err
	}

	return job.configDatabase.UnlockOutgoingAccessRequest(ctx, request)
}

func (job *SynchronizeOutgoingAccessRequestJob) synchronize(ctx context.Context, request *database.OutgoingAccessRequest) error {
	var (
		err          error
		referenceID  uint
		newState     database.OutgoingAccessRequestState
		errorDetails *diagnostics.ErrorDetails
	)

	switch request.State {
	case database.OutgoingAccessRequestCreated:
		newState, referenceID, err = job.sendAccessRequest(ctx, request)

	case database.OutgoingAccessRequestReceived:
		newState, err = job.getAccessRequestState(ctx, request)

	case database.OutgoingAccessRequestApproved:
		err = job.syncAccessProof(ctx, request)
		newState = request.State

	default:
		return fmt.Errorf("invalid state '%s' for pending access request", request.State)
	}

	if err != nil {
		if errors.Is(err, server.ErrServiceDoesNotExist) {
			return job.configDatabase.DeleteOutgoingAccessRequests(ctx, request.Organization.SerialNumber, request.ServiceName)
		}

		// Return err to prevent the state of the outgoing access to be set to failed.
		// If the state of the outgoing access request is set failed it will no longer be picked up by the scheduler.
		if request.State == database.OutgoingAccessRequestApproved || request.State == database.OutgoingAccessRequestReceived {
			return err
		}

		newState = database.OutgoingAccessRequestFailed
		errorDetails = diagnostics.ParseError(err)

		st, ok := status.FromError(err)
		if ok {
			if st.Code() == codes.NotFound {
				errorDetails = errorDetails.WithCode(diagnostics.NoInwaySelectedError)
			}
		}
	}

	return job.configDatabase.UpdateOutgoingAccessRequestState(
		ctx,
		request.ID,
		newState,
		referenceID,
		errorDetails,
	)
}

func (job *SynchronizeOutgoingAccessRequestJob) sendAccessRequest(ctx context.Context, request *database.OutgoingAccessRequest) (database.OutgoingAccessRequestState, uint, error) {
	client, err := job.getOrganizationManagementClient(ctx, request.Organization.SerialNumber)
	if err != nil {
		return database.OutgoingAccessRequestFailed, 0, err
	}

	defer client.Close()

	response, err := client.RequestAccess(ctx, &external.RequestAccessRequest{
		ServiceName: request.ServiceName,
	}, grpc_retry.WithMax(maxRetries))
	if err != nil {
		return database.OutgoingAccessRequestFailed, 0, err
	}

	accessRequestState, err := convertAccessRequestState(response.AccessRequestState)
	if err != nil {
		return database.OutgoingAccessRequestFailed, 0, err
	}

	return accessRequestState, uint(response.ReferenceId), nil
}

func (job *SynchronizeOutgoingAccessRequestJob) getAccessRequestState(ctx context.Context, request *database.OutgoingAccessRequest) (database.OutgoingAccessRequestState, error) {
	client, err := job.getOrganizationManagementClient(ctx, request.Organization.SerialNumber)
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

	return convertAccessRequestState(response.State)
}

func (job *SynchronizeOutgoingAccessRequestJob) syncAccessProof(ctx context.Context, outgoingAccessRequest *database.OutgoingAccessRequest) error {
	remoteProof, err := job.retrieveAccessProof(ctx, outgoingAccessRequest.Organization.SerialNumber, outgoingAccessRequest.ServiceName)
	if err != nil {
		return err
	}

	// skip this AccessRequest as it's not the one related to this AccessProof
	if remoteProof.OutgoingAccessRequest.ID != outgoingAccessRequest.ReferenceID {
		return nil
	}

	localProof, err := job.configDatabase.GetAccessProofForOutgoingAccessRequest(
		ctx,
		outgoingAccessRequest.ID,
	)

	switch err {
	case nil:
	case database.ErrNotFound:
		_, err = job.configDatabase.CreateAccessProof(ctx, outgoingAccessRequest)

		return err
	default:
		return err
	}

	if remoteProof.RevokedAt.Valid &&
		!localProof.RevokedAt.Valid {
		if _, err := job.configDatabase.RevokeAccessProof(
			ctx,
			localProof.ID,
			remoteProof.RevokedAt.Time,
		); err != nil {
			return err
		}
	}

	return nil
}

func (job *SynchronizeOutgoingAccessRequestJob) retrieveAccessProof(ctx context.Context, organizationSerialNumber, serviceName string) (*database.AccessProof, error) {
	client, err := job.getOrganizationManagementClient(ctx, organizationSerialNumber)
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

	return parseAccessProof(response)
}

func parseAccessProof(accessProof *api.AccessProof) (*database.AccessProof, error) {
	var createdAt time.Time

	if accessProof.CreatedAt != nil {
		err := accessProof.CreatedAt.CheckValid()
		if err != nil {
			return nil, ErrInvalidTimeStamp
		}

		createdAt = accessProof.CreatedAt.AsTime()
	}

	revokedAt := accessProof.RevokedAt.AsTime()

	err := accessProof.RevokedAt.CheckValid()
	if err != nil {
		revokedAt = time.Time{}
	}

	dbAccessProof := &database.AccessProof{
		ID:        uint(accessProof.Id),
		CreatedAt: createdAt,
		RevokedAt: sql.NullTime{
			Time:  revokedAt,
			Valid: !revokedAt.IsZero(),
		},
		OutgoingAccessRequest: &database.OutgoingAccessRequest{
			ID:          uint(accessProof.AccessRequestId),
			ServiceName: accessProof.ServiceName,
		},
	}

	if accessProof.Organization != nil {
		dbAccessProof.OutgoingAccessRequest.Organization = database.Organization{
			SerialNumber: accessProof.Organization.SerialNumber,
			Name:         accessProof.Organization.Name,
		}
	}

	return dbAccessProof, nil
}

func (job *SynchronizeOutgoingAccessRequestJob) getOrganizationManagementClient(ctx context.Context, organizationSerialNumber string) (management.Client, error) {
	address, err := job.directoryClient.GetOrganizationInwayProxyAddress(ctx, organizationSerialNumber)
	if err != nil {
		return nil, err
	}

	client, err := job.createManagementClientFunc(ctx, address, job.orgCert)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func convertAccessRequestState(state api.AccessRequestState) (database.OutgoingAccessRequestState, error) {
	switch state {
	case api.AccessRequestState_CREATED:
		return database.OutgoingAccessRequestCreated, nil
	case api.AccessRequestState_APPROVED:
		return database.OutgoingAccessRequestApproved, nil
	case api.AccessRequestState_REJECTED:
		return database.OutgoingAccessRequestRejected, nil
	case api.AccessRequestState_RECEIVED:
		return database.OutgoingAccessRequestReceived, nil
	case api.AccessRequestState_FAILED:
		return database.OutgoingAccessRequestFailed, nil
	default:
		return "", fmt.Errorf("invalid state for outgoing access request: %s", state)
	}
}
